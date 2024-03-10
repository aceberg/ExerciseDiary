package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/auth"
	"github.com/aceberg/ExerciseDiary/internal/check"
	"github.com/aceberg/ExerciseDiary/internal/conf"
	"github.com/aceberg/ExerciseDiary/internal/db"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config.yaml"
	check.Path(confPath)

	appConfig, authConf = conf.Get(confPath)

	appConfig.DirPath = dirPath
	appConfig.DBPath = dirPath + "/sqlite.db"
	check.Path(appConfig.DBPath)
	appConfig.ConfPath = confPath
	appConfig.NodePath = nodePath
	appConfig.Icon = icon

	log.Println("INFO: starting web gui with config", appConfig.ConfPath)

	db.Create(appConfig.DBPath)

	address := appConfig.Host + ":" + appConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ) // templates

	router.StaticFS("/fs/", http.FS(pubFS)) // public

	router.GET("/login/", loginHandler)  // login.go
	router.POST("/login/", loginHandler) // login.go

	router.GET("/", auth.Auth(&authConf), indexHandler)             // index.go
	router.GET("/config/", auth.Auth(&authConf), configHandler)     // config.go
	router.GET("/exercise/", auth.Auth(&authConf), exerciseHandler) // exercise.go
	router.GET("/stats/", auth.Auth(&authConf), statsHandler)       // stats.go
	router.GET("/weight/", auth.Auth(&authConf), weightHandler)     // weight.go

	router.POST("/config/", auth.Auth(&authConf), saveConfigHandler)     // config.go
	router.POST("/config/auth", auth.Auth(&authConf), saveConfigAuth)    // config.go
	router.POST("/exercise/", auth.Auth(&authConf), saveExerciseHandler) // exercise.go
	router.POST("/exdel/", auth.Auth(&authConf), deleteExerciseHandler)  // exercise.go
	router.POST("/set/", auth.Auth(&authConf), setHandler)               // set.go
	router.POST("/weight/", auth.Auth(&authConf), addWeightHandler)      // weight.go

	err := router.Run(address)
	check.IfError(err)
}
