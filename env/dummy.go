//go:build dummy
// +build dummy

package env

var (
	Mode = "debug"
	Etcd = etcd{
		Key:      "/config/debug",
		Username: "xxx",
		Password: "xxxxxx",
	}
)

type etcd struct {
	Key      string
	Username string
	Password string
}
