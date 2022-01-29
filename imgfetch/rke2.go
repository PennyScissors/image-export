package imgfetch

import "fmt"

type RKE2Fetcher struct{}

func (r2 RKE2Fetcher) FetchImages(params FetchParams, imagesSet map[string]ImageMatadata) error {
	fmt.Println("inside rke2 fetch")
	return nil
}
