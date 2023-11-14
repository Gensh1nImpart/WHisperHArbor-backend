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
	Redis struct {
		Addr          string `yaml:"Addr"`
		Passwd        string `yaml:"Passwd"`
		DB            string `yaml:"DB"`
		CacheDuration string `yaml:"CacheDuration"`
	} `yaml:"redis"`
}

type Email struct {
	Server  string `yaml:"server"`
	Port    string `yaml:"port"`
	Account string `yaml:"account"`
	Passwd  string `yaml:"passwd"`
}
