package app

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Endpoint     string `envconfig:"ENDPOINT" required:"true"`
	Region       string `envconfig:"REGION" required:"true"`
	AccessKey    string `envconfig:"ACCESS_KEY" required:"true"`
	SecretKey    string `envconfig:"SECRET_KEY" required:"true"`
	BucketName   string `envconfig:"BUCKET_NAME" required:"true"`
	GettingCount int    `envconfig:"GETTING_COUNT" default:"10"`
}

func initConfig() (*Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return &Config{}, err
	}

	return &config, nil
}
