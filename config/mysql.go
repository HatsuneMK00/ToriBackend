package config

type Mysql struct {
	Path         string `json:"path" yaml:"path"`                                                   // 服务器地址
	Port         string `json:"port" yaml:"port"`                                                   // 端口
	Config       string `json:"config" yaml:"config"`                                               // 高级配置
	Dbname       string `json:"db-name" yaml:"db-name" mapstructure:"db-name"`                      // 数据库名
	Username     string `json:"username" yaml:"username"`                                           // 数据库用户名
	Password     string `json:"password" yaml:"password"`                                           // 数据库密码
	MaxIdleConns int    `json:"max-idle-conns" yaml:"max-idle-conns" mapstructure:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `json:"max-open-conns" yaml:"max-open-conns" mapstructure:"max-open-conns"` // 打开到数据库的最大连接数
}
