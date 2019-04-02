package config

import (
	"github.com/jingweno/jqplay/jq"
	"github.com/joeshaw/envdecode"
)

type Config struct {
	Host        string `env:"HOST,default=0.0.0.0"`
	Port        string `env:"PORT,default=3000"`
	GinMode     string `env:"GIN_MODE,default=debug"`
	DatabaseURL string `env:"DATABASE_URL,required"`
	AssetHost   string `env:"ASSET_HOST"`
	JQVer       string
}

func (c *Config) IsProd() bool {
	return c.GinMode == "release"
}

func Load() (*Config, error) {
	conf := &Config{}
	err := envdecode.Decode(conf)
	if err != nil {
		return nil, err
	}

	conf.JQVer = jq.Version

	return conf, nil
}
