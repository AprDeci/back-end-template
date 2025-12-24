package config

type Logger struct {
	WriteInFile    bool    `mapstructure:"writeInFile" json:"writeInFile" yaml:"writeInFile"`
	WriteInConsole bool    `mapstructure:"writeInConsole" json:"writeInConsole" yaml:"writeInConsole"`
	LogFile        LogFile `mapstructure:"logFile" json:"logFile" yaml:"logFile"`
}

type LogFile struct {
	FilePath  string `mapstructure:"filePath" json:"filePath" yaml:"filePath"`
	MaxSize   int    `mapstructure:"maxSize" json:"maxSize" yaml:"maxSize"`
	MaxAge    int    `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`
	MaxBackup int    `mapstructure:"maxBackup" json:"maxBackup" yaml:"maxBackup"`
	Compress  bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
