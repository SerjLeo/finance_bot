package app

type dbConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
}

type botConfig struct {
	Token string `yaml:"token"`
}

type Config struct {
	Database dbConfig  `yaml:"database"`
	Bot      botConfig `yaml:"bot"`
}
