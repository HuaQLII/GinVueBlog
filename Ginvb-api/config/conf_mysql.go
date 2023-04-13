package config

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"DB"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	logLevel string `yaml:"log_level"` //日志等级，debug就是全部输出sql，dev,release
}
