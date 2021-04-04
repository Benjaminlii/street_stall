package config

// 基础配置类
type BasicConfig struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	Http  Http  `yaml:"http"`
}

// Mysql db相关配置
type Mysql struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DB        string `yaml:"db"`
	UserName  string `yaml:"username"`
	PassWorld string `yaml:"password"`
}

// Redis redis相关配置
type Redis struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DB        string `yaml:"db"`
	PassWorld string `yaml:"password"`
}

// Http Http相关配置
type Http struct {
	SMS SMS `yaml:"sms"`
}

// SMS 短信验证码相关配置
type SMS struct {
	ApiUrl     string `yaml:"api_url"`
	SecretId   string `yaml:"secret_id"`
	SecretKey  string `yaml:"secret_key"`
	TemplateId string `yaml:"template_id"`
}
