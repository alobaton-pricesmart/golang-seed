package config

import (
	"io/ioutil"
	"os"

	"github.com/altipla-consulting/services"
	"github.com/juju/errors"
	"gopkg.in/yaml.v2"
)

var Settings SettingsRoot

type SettingsRoot struct {
	Database Database `yaml:"database"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
}

func ParseSettings() error {
	path := "/etc/auth/config.yml"
	if services.IsLocal() {
		path = "/etc/auth/config.dev.yml"
	}

	f, err := os.Open(path)
	if err != nil {
		return errors.Trace(err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Trace(err)
	}

	if err := yaml.Unmarshal(content, &Settings); err != nil {
		return errors.Trace(err)
	}

	return nil
}
