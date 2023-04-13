package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       int    `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`    //是否显示行号
	LogInConsole bool   `yaml:"logInConsole"` //是否显示打印路径
}
