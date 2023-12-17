package web

import (
	"html/template"
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/check"
	"github.com/aceberg/ExerciseDiary/internal/conf"
	"github.com/aceberg/ExerciseDiary/internal/db"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config.yaml"
	check.Path(confPath)

	appConfig = conf.Get(confPath)

	appConfig.DirPath = dirPath
	appConfig.DBPath = dirPath + "/sqlite.db"
	check.Path(appConfig.DBPath)
	appConfig.ConfPath = confPath
	appConfig.NodePath = nodePath
	appConfig.Icon = icon

	log.Println("INFO: starting web gui with config", appConfig.ConfPath)

	db.Create(appConfig.DBPath)
	// REMOVE LATER
	createExercises()

	address := appConfig.Host + ":" + appConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router.LoadHTMLGlob(TemplPath + "/*.html")
	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ)

	router.Static("/public/", "../../internal/web/public")

	router.GET("/", indexHandler)              // index.go
	router.GET("/config/", configHandler)      // config.go
	router.POST("/config/", saveConfigHandler) // config.go
	router.POST("/set/", setHandler)           // set.go

	err := router.Run(address)
	check.IfError(err)
}

func createExercises() {
	var ex models.Exercise

	ex.Name = "Pull ups"
	ex.Group = "Day 1"
	exData.Exs = append(exData.Exs, ex)

	ex.Name = "Push ups"
	ex.Group = "Day 1"
	exData.Exs = append(exData.Exs, ex)

	ex.Name = "Squats"
	ex.Group = "Day 2"
	ex.Weight = 40
	ex.Reps = 11
	exData.Exs = append(exData.Exs, ex)
}
