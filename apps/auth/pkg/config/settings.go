package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/juju/errors"
	"gopkg.in/yaml.v2"

	"golang-seed/pkg/server"
)

var Settings SettingsRoot

type SettingsRoot struct {
	Name     string   `yaml:"name"`
	Port     int      `yaml:"port"`
	Database Database `yaml:"database"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Name     string `yaml:"name"`
}

func ParseSettings() error {
	path := "apps/auth/config/config.yml"
	if server.IsLocal() {
		path = "apps/auth/config/config.dev.yml"
	}

	p, err := filepath.Abs(path)
	if err != nil {
		return errors.Trace(err)
	}

	f, err := os.Open(p)
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
