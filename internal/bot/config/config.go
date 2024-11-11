package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type BotConfig struct {
	Bot struct {
		Token string `yaml:"token"`
	} `yaml:"bot"`
}

var cfg BotConfig
var once sync.Once

func GetConfig() *BotConfig {
	once.Do(func() {
		err := cleanenv.ReadConfig("./configs/config.yaml", &cfg)
		if err != nil {
			log.Fatal(err)
		}
	})

	return &cfg
}
