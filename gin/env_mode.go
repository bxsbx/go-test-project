package gin

import (
	"github.com/gin-gonic/gin"
	"os"
)

const (
	DEV  = "dev"
	TEST = "test"
	PROD = "prod"
)

func GetEnvMode() string {
	env := os.Getenv("DREAMENV")
	switch env {
	case PROD:
		return gin.DebugMode
	case TEST:
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}
