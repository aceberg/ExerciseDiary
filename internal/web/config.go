package web

import (
	"log"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"

	// "github.com/aceberg/ExerciseDiary/internal/check"
	"github.com/aceberg/ExerciseDiary/internal/conf"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func configHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "config.html", guiData)
}

func saveConfigHandler(c *gin.Context) {

	appConfig.Host = c.PostForm("host")
	appConfig.Port = c.PostForm("port")
	appConfig.Theme = c.PostForm("theme")
	appConfig.Color = c.PostForm("color")

	conf.Write(appConfig)

	log.Println("INFO: writing new config to", appConfig.ConfPath)

	c.Redirect(http.StatusFound, "/config")
}

// func clearHandler(w http.ResponseWriter, r *http.Request) {

// 	log.Println("INFO: delting all hosts from DB")

// 	db.Clear(appConfig.DbPath)

// 	http.Redirect(w, r, r.Header.Get("Referer"), 302)
// }
