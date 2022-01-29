package imageexport

import (
	"io/ioutil"

	"github.com/pennyscissors/image-export/imgfetch"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"
)

const DefaultConfigFileName = "config.yaml"

type ExportConfig struct {
	RancherTag   string           `yaml:"rancherTag"`
	OutputMode   OutputOptions    `yaml:"outputMode"`
	Os           imgfetch.OS      `yaml:"os"`
	KdmDataURL   string           `yaml:"kdmDataUrl"`
	Core         ComponentOptions `yaml:"core"`
	Rancher      ComponentOptions `yaml:"rancher"`
	Charts       ComponentOptions `yaml:"charts"`
	SystemCharts ComponentOptions `yaml:"systemCharts"`
	Rke          ComponentOptions `yaml:"rke"`
	Rke2         ComponentOptions `yaml:"rke2"`
	K3s          ComponentOptions `yaml:"k3s"`
}

type OutputOptions struct {
	All     bool
	Split   bool
	Feature bool
}

// all
type ComponentOptions struct {
	Enabled    bool `yaml:"enabled"`
	Repository imgfetch.Repository
	Include    []string `yaml:"include"`
	Exclude    []string `yaml:"exclude"`
}

// func ConfigPackage() {
// 	config, _ := ParseConfigFile()
// 	spew.Dump(config)
// 	cf := fetchlib.ChartsFetcher{}
// 	cf.FetchImages()
// }

func ExportImages() error {
	config, err := ParseConfigFile()
	if err != nil {
		return err
	}

}

func ParseConfigFile() (ExportConfig, error) {
	configYaml, err := ioutil.ReadFile(DefaultConfigFileName)
	if err != nil {
		logrus.Fatalf("Unable to find configuration file: %s", err)
	}
	config := ExportConfig{}
	if err := yaml.Unmarshal(configYaml, &config); err != nil {
		logrus.Fatalf("Unable to unmarshall configuration file: %s", err)
	}
	return config, nil
}

// output:
//   all: true # one file with all images
//   split: true # one file of images for each component (core, rancher, charts, rke, etc)
//   feature: true # one file of images per feature (monitoring, istio).

// os:
//   linux: true
//   windows: true

// core:
//   enabled: true

// rancher:
//   enabled: true

// charts:
//   enabled: true
//   repository:
//     location: "https://github.com/rancher/charts.git"
//     branch: "release-v2.6"

// systemCharts:
//   enabled: true
//   repository:
//     location: "https://github.com/rancher/system-charts.git"
//     branch: "release-v2.6"

// rke:
//   enabled: true
//   kdmData: "http://url/to/data.json"

// rke2:
//   enabled: true
//   kdmData: "http://url/to/data.json"
