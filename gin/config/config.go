package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Logger Logger `mapstructure:"logger" json:"logger" yaml:"logger"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
