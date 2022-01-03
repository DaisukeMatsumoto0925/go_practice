package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
)

type Config struct {
	Env    string
	Db     Db
	Server ServerConfig
}

type ServerConfig struct {
	Port string `env:"PORT"`
}

type Db struct {
	Port     string `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
}

var config *Config
var once sync.Once

const projectName = ""

func Get() *Config {
	return config
}

func getProjectRootPath() string {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, projectName) && !strings.HasSuffix(wd, "app") {
		wd = filepath.Dir(wd)
	}

	return wd
}

func init() {
	once.Do(func() {
		goenv := os.Getenv("GO_ENV")
		config = &Config{Env: goenv}

		projectRootPath := getProjectRootPath()
		filepath := fmt.Sprintf("%s/config/%s.toml", projectRootPath, goenv)

		if _, err := toml.DecodeFile(filepath, config); err != nil {
			log.Fatal("failed to load config toml: ", err)
		}

		if err := env.Parse(&config.Db); err != nil {
			fmt.Printf("failed to load Db error: %s", err)
		}
		if err := env.Parse(&config.Server); err != nil {
			fmt.Printf("failed to load Server error: %s", err)
		}
	})
}
