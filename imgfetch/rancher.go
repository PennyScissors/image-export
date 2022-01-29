package imgfetch

import "fmt"

type RancherFetcher struct{}

func (rn RancherFetcher) FetchImages(params FetchParams, imagesSet map[string]ImageMatadata) error {
	fmt.Println("inside core fetch")
	for name := range rn.generateImages(params.RancherVersion) {
		imagesSet[name] = ImageMatadata{
			Name:    name,
			Sources: []string{Core},
			Os:      []OS{Linux, Windows},
		}
	}
	return nil
}

func (rn RancherFetcher) generateImages(rancherVersion string) map[string]struct{} {
	return map[string]struct{}{
		newImage("repo/image", "0.0.0"):        {},
		newImage("repo/image", rancherVersion): {},
	}
}
