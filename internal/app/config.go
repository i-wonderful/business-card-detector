package app

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Name    string `yaml:"name" default:"card_detector"`
	Version string `yaml:"version" default:"1.0.0"`

	Port          int    `yaml:"port" default:"8080"`
	TmpFolder     string `yaml:"tmp_folder" default:"./tmp"`
	StorageFolder string `yaml:"storage_folder" default:"./storage"`

	PathProfessionList string `yaml:"path_profession_list"`
	PathCompanyList    string `yaml:"path_company_list"`
	PathNamesList      string `yaml:"path_names_list"`

	Paddleocr struct {
		RunPath string `yaml:"run_path"`
		DetPath string `yaml:"model_det_path"`
		RecPath string `yaml:"model_rec_path"`
	}

	Log struct {
		Level string `yaml:"level" default:"info"`
		Time  bool   `yaml:"time" default:"true"`
	}

	IsDebug bool `yaml:"debug" default:"false"`
}

const defaultConfigPath = "./config/config.yml"

func NewConfigFromYml() (*Config, error) {
	configFilePath := os.Getenv("CONFIG_FILE")
	if configFilePath == "" {
		configFilePath = defaultConfigPath
	}
	var config Config
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Printf("Error reading YAML file: %s\n", err)
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("Error parsing YAML file: %s\n", err)
		return nil, err
	}

	return &config, nil
}
