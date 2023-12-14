package web

import (
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/check"
	"github.com/aceberg/ExerciseDiary/internal/conf"
	// "github.com/aceberg/ExerciseDiary/internal/yaml"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config.yaml"
	check.Path(confPath)

	appConfig = conf.Get(confPath)

	appConfig.DirPath = dirPath
	appConfig.ConfPath = confPath
	appConfig.NodePath = nodePath
	appConfig.Icon = icon

	log.Println("INFO: starting web gui with config", appConfig.ConfPath)

	address := appConfig.Host + ":" + appConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob(TemplPath + "/*.html")
	router.GET("/", indexHandler)              // index.go
	router.GET("/config/", configHandler)      // config.go
	router.POST("/config/", saveConfigHandler) // config.go

	err := router.Run(address)
	check.IfError(err)
}
