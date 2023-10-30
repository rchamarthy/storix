package storix_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rchamarthy/storix"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"machinerun.io/disko"
)

func save(file string, p *storix.Profile) {
	data, err := yaml.Marshal(map[string]*storix.Profile{"storix": p})
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(file, data, 0644); err != nil {
		panic(err)
	}
}

func mkProfile(dir string) string {
	f := filepath.Join(dir, "storix.yaml")
	s := &storix.Profile{
		Disks: []storix.DiskProfile{
			{Pool: "fast-disks", Match: []storix.DiskMatch{
				{Type: storix.ToDiskType(disko.SSD), Attachment: storix.ToAttachment(disko.PCIE), MinSize: 1 * storix.GB, MaxSize: 1 * storix.TB},
				{Type: storix.ToDiskType(disko.SSD), Attachment: storix.ToAttachment(disko.RAID), MinSize: 1 * storix.GB, MaxSize: 1 * storix.TB},
			}},
			{Pool: "slow-disks", Match: []storix.DiskMatch{
				{Type: storix.ToDiskType(disko.HDD), Attachment: storix.ToAttachment(disko.ATA), MinSize: 1 * storix.GB, MaxSize: 1 * storix.TB},
				{Type: storix.ToDiskType(disko.HDD), Attachment: storix.ToAttachment(disko.RAID), MinSize: 1 * storix.GB, MaxSize: 1 * storix.TB},
			}},
		},
		VolumeGroups: []storix.VGProfile{
			{Label: "fast", Name: "fast0", DiskPool: "fast-disks"},
			{Label: "slow", Name: "slow0", DiskPool: "slow-disks"},
		},
		Volumes: []storix.VolumeProfile{
			{Name: "logs", Group: "fast", Size: 1 * storix.TB, FS: "ext4", Mount: "/var/logs"},
			{Name: "data", Group: "slow", Size: 1 * storix.TB, FS: "ext4", Mount: "/var/data"},
		},
	}

	save(f, s)
	return f
}

func badYAML(dir string) string {
	f := filepath.Join(dir, "bad-storix.yaml")
	if err := os.WriteFile(f, []byte("blah\nblee"), 0644); err != nil {
		panic(err)
	}

	return f
}

func badKey(dir string) string {
	f := filepath.Join(dir, "bad-key-storix.yaml")
	if err := os.WriteFile(f, []byte(`
something:
  disks:
    - label: fast-disks
      match:
      - type: ssd
`), 0644); err != nil {
		panic(err)
	}

	return f
}

func TestProfileLoad(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	tdir, err := os.MkdirTemp("", "storix-")
	assert.NoError(err)
	defer os.RemoveAll(tdir)

	profile, err := storix.Load("blah")
	assert.Error(err)
	assert.Nil(profile)

	p := mkProfile(tdir)
	profile, err = storix.Load(p)
	assert.NoError(err)
	assert.NotNil(profile)

	profile, err = storix.Load(badYAML(tdir))
	assert.Error(err)
	assert.Nil(profile)

	profile, err = storix.Load(badKey(tdir))
	assert.Nil(profile)
	assert.Equal(storix.ProfileKeyError, err)
}
