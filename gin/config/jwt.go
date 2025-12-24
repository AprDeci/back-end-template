package config

type JWT struct {
	Secret  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Expired int64  `mapstructure:"expired" json:"expired" yaml:"expired"`
}
