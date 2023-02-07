package config

type DBConfig struct {
	UserName string `required:"true" envconfig:"DB_USERNAME" `
	Port     int    `required:"true" envconfig:"DB_PORT"`
	Host     string `required:"true" envconfig:"DB_HOST"`
	Password string `required:"true" envconfig:"DB_PASSWORD"`
	Name     string `required:"true" envconfig:"DB_NAME"`
}

type AppConfig struct {
	Debug bool   `required:"true" envconfig:"DEBUG"`
	Port  int    `required:"true" envconfig:"PORT"`
	Host  string `required:"true" envconfig:"HOST"`
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
