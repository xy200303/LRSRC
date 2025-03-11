package main

import (
	"xiaoyun/backend/config"
	"xiaoyun/backend/routes"
)

func main() {
	switch config.Config.Server.StartMode {
	case "api":
		routes.Api()
	case "all":
		routes.All()
	}
}
