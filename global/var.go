package global

import (
	"github.com/tech-showcase/auth-service/config"
)

var Configuration = config.Config{}

func init() {
	Configuration = config.Read()
}
