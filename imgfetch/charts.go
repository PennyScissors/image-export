package imgfetch

import (
	"fmt"
	"io"
	"os"

	billy "github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	log "github.com/sirupsen/logrus"

	// yaml "gopkg.in/yaml.v2"
	// "helm.sh/helm/v3/pkg/repo"

	"helm.sh/helm/v3/pkg/repo"
)

type ChartsFetcher struct{}

func (c ChartsFetcher) FetchImages(params FetchParams, imagesSet map[string]ImageMatadata) error {
	fmt.Println("inside charts fetch")
	params = FetchParams{
		ChartsRepo: Repository{
			URL:    "https://github.com/rancher/charts.git",
			Branch: "release-v2.6",
		},
	}
	fs, err := loadGitRepoToMemory(params.ChartsRepo.URL, params.ChartsRepo.Branch)
	if err != nil {
		return err
	}

	index, err := c.loadIndexFile(fs)
	if err != nil {
		return err
	}

	for _, versions := range index.Entries {
		l(versions[0].Metadata.Name)
	}
	// load index
	// filter entries
	// Check values files of entries

	return nil
}

func (c ChartsFetcher) loadIndexFile(fs billy.Filesystem) (*repo.IndexFile, error) {
	indexFile, err := fs.Open("index.yaml")
	if err != nil {
		log.Info(err)
	}
	defer indexFile.Close()

	data, err := io.ReadAll(indexFile)
	if err != nil {
		l(err)
	}

	tempIndexFile, err := os.CreateTemp("", "index")
	if err != nil {
		l(err)
	}
	defer os.Remove(tempIndexFile.Name())
	defer tempIndexFile.Close()

	_, err = tempIndexFile.Write(data)
	if err != nil {
		l(err)
	}

	index, err := repo.LoadIndexFile(tempIndexFile.Name())
	if err != nil {
		l(err)
	}

	return index, nil
}

func loadGitRepoToMemory(URL, branch string) (billy.Filesystem, error) {
	// _, err := git.PlainClone(url, false, &git.CloneOptions{
	// 	URL:      url,
	// 	Progress: os.Stdout,
	// })
	// ref := fmt.Sprintf("ref/heads/%s", branch)
	fs := memfs.New()
	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           URL,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		SingleBranch:  true,
		Depth:         1,
	})
	if err != nil {
		return nil, err
	}
	return fs, nil
}
