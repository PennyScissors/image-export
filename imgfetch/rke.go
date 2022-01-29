package imgfetch

import "fmt"

type RKEFetcher struct{}

func (r RKEFetcher) FetchImages(params FetchParams, imagesSet map[string]ImageMatadata) error {
	fmt.Println("inside rke fetch")
	return nil
}
