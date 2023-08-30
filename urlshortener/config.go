package urlshortener

import (
	"fmt"

	"github.com/Laminator42/go-urlshortener-mongo/util"
)

type AppConfig struct {
	Host string
	Port string
}

func (c AppConfig) HostAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c AppConfig) BaseUrl() string {
	return fmt.Sprintf("http://%s:%s/", c.Host, c.Port)
}

var AppConf = AppConfig{
	Host: util.GetEnv("APP_HOST", "localhost"),
	Port: util.GetEnv("APP_PORT", "8080"),
}
