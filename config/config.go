package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"runtime"
	"strings"
)

type Config struct {
	Server struct {
		Port string `envconfig:"APP_PORT" required:"true"`
		Mode string `envconfig:"APP_MODE" required:"true"`
		Name string `envconfig:"APP_NAME" required:"true"`
	}

	Redis struct {
		Addr     string `envconfig:"REDIS_ADDR" required:"true"`
		Password string `envconfig:"REDIS_PASSWORD" required:"true"`
		DB       int    `envconfig:"REDIS_DB" required:"true"`
	}
	Grpc struct {
		Url string `envconfig:"GRPC_URL" required:"true"`
	}

	Nsq struct {
		PublisherAddr string `envconfig:"NSQD_PUBLISHER_ADDR" required:"true"`
	}

	Cache struct {
		TimeRefresh string `envconfig:"TIME_TO_REFRESH_CACHE" required:"true"`
	}
}

func LoadTest() *Config {
	return load(".env.test")
}

func LoadDefault() *Config {
	return load(".env")
}

func load(env string) *Config {
	var config Config
	readEnv(&config, env)
	err := envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}
	return &config
}

func readEnv(c *Config, env string) {
	err := godotenv.Overload(getSourcesPath() + "/../" + env)
	if err != nil {
		log.Println(err)
	}
}

func getSourcesPath() string {
	_, filename, _, _ := runtime.Caller(0)
	if strings.Contains(filename, "config.go") {
		filename = strings.Replace(filename, "/config.go", "", 1)
	}
	return filename
}
