package config

import (
	"os"
	"path/filepath"
	"time"

	"autograder/pkg/utils"

	"github.com/sirupsen/logrus"
)

const (
	ConfigPathEnvKey = "CONFIG_PATH"
)

var Instance *Config

func init() {
	Instance = InitConfig()
}

func InitConfig() *Config {
	configPath := os.Getenv(ConfigPathEnvKey)
	if configPath == "" {
		configPath, _ = filepath.Abs("./config.yml")
	}
	logrus.Infof("[InitConfig] start to init, config path is %s", configPath)
	config := Config{}
	err := utils.UnmarshalYamlFile(configPath, &config)
	if err != nil {
		logrus.Fatalf("[InitConfig] init config failed, error %+v", err)
		panic(err)
	}
	logrus.Infof("[InitConfig] init config succeeded, config: %s", utils.FormatJsonString(config))
	return &config
}

type Config struct {
	WorkDir      string       `yaml:"WorkDir"`
	TestcasesDir string       `yaml:"TestcasesDir"`
	CanvasToken  string       `yaml:"CanvasToken"`
	SystemDB     *MysqlConfig `yaml:"SystemDB"`
	EBookStoreDB *MysqlConfig `yaml:"EBookStoreDB"`
	Token        *TokenConfig `yaml:"Token"`
}

type MysqlConfig struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Database string `yaml:"Database"`
	Timeout  string `yaml:"Timeout"`
	ShowSql  bool   `yaml:"ShowSql"`
}

type TokenConfig struct {
	Secret      string        `yaml:"Secret"`
	ExpireAfter time.Duration `yaml:"ExpireAfter"`
}
