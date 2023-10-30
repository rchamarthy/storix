package storix

import (
	"strings"

	"machinerun.io/disko"
)

// For YAML parsing as disko types dont support YAML
type DiskType disko.DiskType

func ToDiskType(d disko.DiskType) DiskType {
	return DiskType(d)
}

func (d DiskType) String() string {
	return disko.DiskType(d).String()
}

func (d DiskType) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DiskType) UnmarshalText(text []byte) error {
	s := strings.ToUpper(string(text))
	*d = ToDiskType(disko.StringToDiskType(s))
	return nil
}

// For YAML parsing as disko types dont support YAML
type AttachmentType disko.AttachmentType

func ToAttachment(a disko.AttachmentType) AttachmentType {
	return AttachmentType(a)
}

func (a AttachmentType) String() string {
	return disko.AttachmentType(a).String()
}

func (a AttachmentType) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

func (a *AttachmentType) UnmarshalText(text []byte) error {
	s := strings.ToUpper(string(text))
	*a = ToAttachment(disko.StringToAttachmentType(s))
	return nil
}

type DiskProfile struct {
	Pool  string      `json:"disk-pool" yaml:"disk-pool"`
	Match []DiskMatch `json:"match" yaml:"match"`
}

type DiskMatch struct {
	Type       DiskType       `json:"type,omitempty" yaml:"type,omitempty"`
	Attachment AttachmentType `json:"attachment,omitempty" yaml:"attachment,omitempty"`
	MinSize    ByteSize       `json:"min,omitempty" yaml:"min,omitempty"`
	MaxSize    ByteSize       `json:"max,omitempty" yaml:"max,omitempty"`
}
