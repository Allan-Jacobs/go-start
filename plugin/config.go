package plugin

import (
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"os"
	"path"
)

type Config struct {
	ModuleUrlBase string `json:"module_url_base"`
}

func SetConfig(config Config) error {

	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	file, err := os.Create(path.Join(dir, "go-start", "config.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func createDefaultConfig() error {
	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	config_path := path.Join(dir, "go-start", "config.json")

	if _, err := os.Stat(config_path); err == nil {
		return nil // exists

	} else if !errors.Is(err, fs.ErrNotExist) {
		return err
	} // file doesnt exist, create it

	err = os.MkdirAll(path.Join(dir, "go-start"), 0700|fs.ModeDir)
	if err != nil {
		return err
	}

	file, err := os.Create(config_path)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := json.MarshalIndent(Config{}, "", "\t")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	return err // nil if successful
}

func GetConfig() (Config, error) {
	var config Config

	dir, err := os.UserConfigDir()
	if err != nil {
		return config, err
	}

	config_path := path.Join(dir, "go-start", "config.json")

	createDefaultConfig() // creates default config if it doesnt exist

	file, err := os.Open(config_path)

	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
