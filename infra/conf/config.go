package conf

type Config struct {
	Server   Server   `json:"server"   yaml:"server"`
	Database Database `json:"database" yaml:"database"`
	Logger   Logger   `json:"logger"   yaml:"logger"`
}
