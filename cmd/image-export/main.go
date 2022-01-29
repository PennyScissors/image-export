package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/pennyscissors/image-export/imgfetch"
)

type ImageName string

func main() {
	fmt.Println("inside cmd/image-export/main.go")
	// cf := fetchlib.ChartsFetcher{}
	// err := cf.FetchImages(fetchlib.FetchParams{}, map[string]fetchlib.ImageMatadata{})
	// if err != nil {
	// 	logrus.Info(err)
	// }

	imagesSet := make(map[string]imgfetch.ImageMatadata)
	co := imgfetch.CoreFetcher{}
	_ = co.FetchImages(imgfetch.FetchParams{}, imagesSet)
	spew.Dump(imagesSet)
}
