package config

type DBConfig struct {
	UserName string `required:"true" envconfig:"USERNAME" `
	Port     int    `required:"true" envconfig:"PORT"`
	Host     string `required:"true" envconfig:"HOST"`
	Password string `required:"true" envconfig:"PASSWORD"`
	Name     string `required:"true" envconfig:"NAME"`
}

type AppConfig struct {
	Debug    bool   `required:"true" envconfig:"DEBUG"`
	Port     int    `required:"true" envconfig:"PORT"`
	Host     string `required:"true" envconfig:"HOST"`
	LogLevel string `envconfig:"HOST" default:"info"`
}

type Config struct {
	DB  DBConfig
	App AppConfig
}

func (c *Config) GetDebug() bool {
	return c.App.Debug
}
func (c *Config) GetHost() string {
	return c.App.Host
}
func (c *Config) GetPort() int {
	return c.App.Port
}

func (c *Config) GetDBUserName() string {
	return c.DB.UserName
}

func (c *Config) GetDBPassword() string {
	return c.DB.Password
}

func (c *Config) GetDBHost() string {
	return c.DB.Host
}

func (c *Config) GetDBPort() int {
	return c.DB.Port
}

func (c *Config) GetDBName() string {
	return c.DB.Name
}
