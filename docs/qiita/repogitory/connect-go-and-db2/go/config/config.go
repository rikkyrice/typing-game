package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"../db"

	"gopkg.in/yaml.v2"
)

// Config structure that defines the settings
type Config struct {
	DB *db.DBConnConfig `yaml:"db"`
}

var configuration *Config = nil

// Init initialize the settings
func Init(filename string) (*Config, error) {
	if filename == "" {
		return nil, errors.New("設定ファイルが指定されていません。")
	}
	err := load(filename)
	if err != nil {
		return nil, fmt.Errorf("設定ファイルの読み込みに失敗しました。%+v", err)
	}
	return configuration, nil
}

func load(filename string) error {
	data, err := readFile(filename)
	if err != nil {
		return err
	}
	c := &Config{}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	configuration = c
	return nil
}

// GetConfig gets the settings
func GetConfig() *Config {
	return configuration
}

// ReadFile reads file
func readFile(filename string) ([]byte, error) {
	path, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}
