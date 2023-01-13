package app

import (
	"github.com/68696c6c/goat"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	// TODO: move this to goat
	Version string
}

func GetConfig() (Config, error) {
	result := Config{
		Version: viper.GetString("build_tag"),
	}
	err := validateConfig(result)
	if err != nil {
		return Config{}, err
	}
	return result, nil
}

func validateConfig(c Config) error {
	var errs []error
	if c.Version == "" {
		errs = append(errs, errors.New("invalid version number"))
	}
	return goat.ErrorsToError(errs)
}
