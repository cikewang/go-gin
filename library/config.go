package library

var Config = struct {
	// 项目名称
	APPName string `default:"app name"`

	// 数据库配置
	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	// Redis 配置
	Redis struct {
		Host     string `default:"127.0.0.1"`
		Password string `required:"true" default:"123456"`
		Port     uint   `default:"6379"`
	}
}{}
