package web

import (
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/AppTemplate/internal/check"
	"github.com/aceberg/AppTemplate/internal/conf"
	// "github.com/aceberg/AppTemplate/internal/yaml"
)

// Gui - start web server
func Gui(confPath, yamlPath, nodePath string) {

	AppConfig = conf.Get(confPath)
	AppConfig.ConfPath = confPath
	AppConfig.NodePath = ""
	AppConfig.YamlPath = yamlPath
	AppConfig.Icon = Icon
	log.Println("INFO: starting web gui with config", AppConfig.ConfPath)

	// AllLinks = yaml.Read(AppConfig.YamlPath)
	// log.Println("ALL:", AllLinks)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob(TemplPath + "/*.html")
	router.GET("/", indexHandler) // index.go

	err := router.Run(address)
	check.IfError(err)
}
