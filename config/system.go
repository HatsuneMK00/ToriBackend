package config

type System struct {
	Env  string `json:"env" yaml:"env"`   // 环境值
	Port int    `json:"port" yaml:"port"` // 端口值
}
