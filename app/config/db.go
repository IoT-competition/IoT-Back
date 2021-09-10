package config

// DBType 数据库配置定义
type DBType struct {
	URL string `yaml:"url"`
}

// GetURL 获取数据库 URL
func (d *DBType) GetURL() string {
	return d.URL
}
