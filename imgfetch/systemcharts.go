package imgfetch

import "fmt"

type SystemChartsFetcher struct{}

func (sc SystemChartsFetcher) FetchImages(params FetchParams, imagesSet map[string]ImageMatadata) error {
	fmt.Println("inside systemcharts fetch")
	// clone repo
	// load index
	// filter entries
	// Check values files of entries

	return nil
}
