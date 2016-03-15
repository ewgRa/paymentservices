package service

type Config struct {
	host       string
	port       int
	path       string
	metricPath string
}

func NewConfig(host string, port int, path, metricPath string) *Config {
	return &Config {
		host: host,
		port: port,
		path: path,
		metricPath: metricPath,
	}
}

func (c *Config) GetPath() string {
	return c.path
}

func (c *Config) GetMetricPath() string {
	return c.metricPath
}

func (c *Config) GetHost() string {
	return c.host
}

func (c *Config) GetPort() int {
	return c.port
}