package ogone

// Config keep ogone service configuration parameters
type Config struct {
	pspID    string
	userID   string
	password string
	sign     string
	sandbox  bool
}

// NewConfig create new Config configuration for ogone endpoint
func NewConfig(pspID, userID, password, sign string, sandbox bool) *Config {
	return &Config{
		pspID:    pspID,
		userID:   userID,
		password: password,
		sign:     sign,
		sandbox:  sandbox,
	}
}

// GetPspID return pspID configuration
func (c *Config) GetPspID() string {
	return c.pspID
}

// GetUserID return userID configuration
func (c *Config) GetUserID() string {
	return c.userID
}

// GetPassword return password configuration
func (c *Config) GetPassword() string {
	return c.password
}

// GetSign return sign configuration
func (c *Config) GetSign() string {
	return c.sign
}

// IsSandbox true if endpoint must call sandbox instead of live gateway
func (c *Config) IsSandbox() bool {
	return c.sandbox
}
