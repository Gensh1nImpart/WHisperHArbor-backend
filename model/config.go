package model

type Config struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Passwd   string `yaml:"passwd"`
		Database string `yaml:"database"`
		Charset  string `yaml:"charset"`
	} `yaml:"mysql"`
	Base struct {
		Port string `yaml:"port"`
	} `yaml:"base"`
}
