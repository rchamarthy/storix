package storix

type VGProfile struct {
	Label    string `json:"label" yaml:"label"`
	Name     string `json:"name" yaml:"name"`
	DiskPool string `json:"disk-pool" yaml:"disk-pool"`
}

type VolumeProfile struct {
	Name  string   `json:"name" yaml:"name"`
	Group string   `json:"group" yaml:"group"`
	Size  ByteSize `json:"size" yaml:"size"`
	FS    string   `json:"fs" yaml:"fs"`
	Mount string   `json:"mount" yaml:"mount"`
}
