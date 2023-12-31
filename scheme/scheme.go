package scheme

import (
	"gopkg.in/yaml.v3"
)

type Scheme struct {
	Items []*SchemeItem
}

func New(r []byte) (*Scheme, error) {
	s := &Scheme{}

	err := yaml.Unmarshal(r, &s.Items)
	if err != nil {
		return nil, err
	}

	for _, i := range s.Items {
		if i.Type == "" {
			i.Type = "string"
		}
	}

	return s, nil
}

type SchemeItem struct {
	Name     string
	Optional bool
	// Datetime format as a Go format string
	SrcFormat string `yaml:"src_format"`
	DstFormat string `yaml:"dst_format"`

	// Whether an Array/Object will be displayed inline or on multiple lines.
	Multiline bool

	SchemeLiteral `yaml:",inline"`

	Enum    []EnumVariant
	Prefix  *Decorator
	Postfix *Decorator
}

type SchemeLiteral struct {
	Literal      string
	DisplayProps `yaml:",inline"`
}

type Decorator struct {
	SchemeLiteral `yaml:",inline"`
	Inherit       bool
}

type EnumVariant struct {
	Value string

	DisplayProps `yaml:",inline"`
}
