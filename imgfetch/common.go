package imgfetch

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// Sources
const Core = "core"
const Rancher = "rancher"
const RKE = "rke"
const RKE2 = "rke2"
const K3S = "k3s"

type OS int

const (
	Linux OS = iota
	Windows
)

type ImageMatadata struct {
	Name    string
	Sources []string
	Os      []OS
}

type FetchParams struct {
	RancherVersion string
	Os             OS
	KdmData        string
	ChartsRepo     Repository
	SysChartsRepo  Repository
}

type Repository struct {
	URL    string
	Branch string
}

func genImageName(name string, version string) string {
	return fmt.Sprintf("%s:%s", name, version)
}

func d(i interface{}) {
	spew.Dump(i)
}

func l(i interface{}) {
	log.Info(i)
}

// // CheckIfError should be used to naively panics if an error is not nil.
// func CheckIfError(err error) {
// 	if err == nil {
// 		return
// 	}
// 	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
// 	os.Exit(1)
// }
