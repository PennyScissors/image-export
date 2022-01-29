package imgfetch

import "fmt"

type K3SFetcher struct{}

func (k K3SFetcher) FetchImages(params FetchParams, imagesSet map[string]ImageMatadata) error {
	fmt.Println("inside k3s fetch")

	return nil
}
