package storix

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var ProfileKeyError = errors.New("profile key storix not found")

type Profile struct {
	Disks        []DiskProfile   `json:"disks" yaml:"disks"`
	VolumeGroups []VGProfile     `json:"volume-groups" yaml:"volume-groups"`
	Volumes      []VolumeProfile `json:"volumes" yaml:"volumes"`
}

func Load(file string) (*Profile, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Parse the data as YAML
	pmap := map[string]*Profile{}
	if err := yaml.Unmarshal(data, &pmap); err != nil {
		return nil, err
	}

	if profile, ok := pmap["storix"]; ok {
		return profile, nil
	}

	return nil, ProfileKeyError
}
