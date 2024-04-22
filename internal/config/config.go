package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	config *Config
)

type option struct {
	configFile string
}

// Init ...
func Init(opts ...Option) error {
	opt := &option{
		configFile: getDefaultConfigFile(),
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(out, &config)
	if err != nil {
		return err
	}

	return nil
}

// Option ...
type Option func(*option)

// WithConfigFile ...
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

func getDefaultConfigFile() string {

	configPath := "./files/project-testing.development.yaml"

	return configPath
}

// Get ...
func Get() *Config {
	return config
}
