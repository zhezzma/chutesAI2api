// @title HIX-AI-2API
// @version 1.0.0
// @description HIX-AI-2API
// @BasePath
package main

import (
	"chutesai2api/check"
	"chutesai2api/common"
	"chutesai2api/common/config"
	logger "chutesai2api/common/loggger"
	"chutesai2api/middleware"
	"chutesai2api/model"
	"chutesai2api/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

//var buildFS embed.FS

func main() {
	logger.SetupLogger()
	logger.SysLog(fmt.Sprintf("chutesai2api %s starting...", common.Version))
	check.CheckEnvVariable()

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	var err error

	model.InitTokenEncoders()

	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middleware.RequestId())
	middleware.SetUpLogger(server)

	router.SetRouter(server)
	//router.SetRouter(server, buildFS)

	var port = os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(*common.Port)
	}

	if config.DebugEnabled {
		logger.SysLog("running in DEBUG mode.")
	}

	logger.SysLog("chutesai2api start success. enjoy it! ^_^\n")

	err = server.Run(":" + port)

	if err != nil {
		logger.FatalLog("failed to start HTTP server: " + err.Error())
	}

}
