package config

type Repo struct{}

type Config struct {
	Repo Repo
}

func NewConfig() *Config {
	return &Config{
		Repo: Repo{},
	}
}
