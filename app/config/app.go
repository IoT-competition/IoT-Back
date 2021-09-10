package config

// AppType app 配置信息
type AppType struct {
	Port    string `yaml:"port"`
	Mode    string `yaml:"mode"`
	Timeout int    `yaml:"timeout"`
	Secret  string `yaml:"secret"`
}

// GetPort 获取运行端口
func (d *AppType) GetPort() string {
	return d.Port
}

// GetMode 获取运行模式
func (d *AppType) GetMode() string {
	return d.Mode
}

// GetTimeout 获取请求超时时间
func (d *AppType) GetTimeout() int {
	return d.Timeout
}

// GetSecret 获取 Session 加密秘钥
func (d *AppType) GetSecret() string {
	return d.Secret
}
