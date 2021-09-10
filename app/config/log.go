package config

// LogType 日志配置类型定义
type LogType struct {
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"maxsize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxage"`
	Compress   bool   `yaml:"compress"`
}

// GetFilename 获得文件名称
func (d *LogType) GetFilename() string {
	return d.Filename
}

// GetMaxSize 获得文件最大大小
func (d *LogType) GetMaxSize() int {
	return d.MaxSize
}

// GetMaxBackups 获得保存文件数目
func (d *LogType) GetMaxBackups() int {
	return d.MaxBackups
}

// GetMaxAge 获得最长保存时间
func (d *LogType) GetMaxAge() int {
	return d.MaxAge
}

// GetCompress 获得是否进行日志压缩
func (d *LogType) GetCompress() bool {
	return d.Compress
}
