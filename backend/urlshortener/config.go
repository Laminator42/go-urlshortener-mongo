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
	switch port := c.Port; port {
	case "80":
		return fmt.Sprintf("http://%s/", c.Host)
	case "443":
		return fmt.Sprintf("https://%s/", c.Host)
	default:
		return fmt.Sprintf("http://%s:%s/", c.Host, c.Port)
	}
}

var AppConf = AppConfig{
	Host: util.GetEnv("BACKEND_HOST", "0.0.0.0"),
	Port: util.GetEnv("BACKEND_PORT", "8080"),
}
