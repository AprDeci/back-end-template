package config

type Logger struct {
	WriteInFile bool `mapstructure:"writeInFile" json:"writeInFile" yaml:"writeInFile"`
}
