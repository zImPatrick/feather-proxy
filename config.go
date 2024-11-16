package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerToken string
}

var defaultConfig = Config{}

func (c *Config) Load(Path string) error {
	// open config file
	file, err := os.OpenFile(Path, os.O_RDONLY, 0666)
	if os.IsNotExist(err) {
		// save default config to that path
		err := defaultConfig.Save(Path)
		if err != nil {
			return err
		}

		// this should probably not be done like this?
		c = &defaultConfig
		return nil
	} else if err != nil {
		return err
	}
	defer file.Close()

	// read json out of it
	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Save(Path string) error {
	file, err := os.Create(Path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(c)
	if err != nil {
		return err
	}

	return nil
}
