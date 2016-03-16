package service

// Config keep service configuration parameters
type Config struct {
	host       string
	port       int
	path       string
	metricPath string
}

// NewConfig create new Config configuration for service
func NewConfig(host string, port int, path, metricPath string) *Config {
	return &Config{
		host:       host,
		port:       port,
		path:       path,
		metricPath: metricPath,
	}
}

// GetPath return path that service handled
func (c *Config) GetPath() string {
	return c.path
}

// GetMetricPath return path that can be used for request metric information
func (c *Config) GetMetricPath() string {
	return c.metricPath
}

// GetHost return host, that service will be listen
func (c *Config) GetHost() string {
	return c.host
}

// GetPort return port, that service will be listen
func (c *Config) GetPort() int {
	return c.port
}
