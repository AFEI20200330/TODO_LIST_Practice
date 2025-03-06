package config

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Charset  string
}

type Config struct {
	DB *DBConfig
}

func NewConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "Fay0810_",
			Name:     "todo_list_db",
			Charset:  "utf8",
		},
	}
}
