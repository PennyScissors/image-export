package imgfetch

import "fmt"

type CoreFetcher struct{}

func (co CoreFetcher) FetchImages(params FetchParams, imagesSet map[string]ImageMatadata) error {
	fmt.Println("inside core fetch")
	for name := range co.generateImages(params.RancherVersion) {
		imagesSet[name] = ImageMatadata{
			Name:    name,
			Sources: []string{Rancher},
			Os:      []OS{Linux, Windows},
		}
	}
	return nil
}

func (co CoreFetcher) generateImages(rancherVersion string) map[string]struct{} {
	return map[string]struct{}{
		genImageName("repo/image", "0.0.0"):        {},
		genImageName("repo/image", rancherVersion): {},
	}
}
