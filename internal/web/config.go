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

	guiData.Config = AppConfig

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "config.html", guiData)
}

func saveConfigHandler(c *gin.Context) {

	AppConfig.Host = c.PostForm("host")
	AppConfig.Port = c.PostForm("port")
	AppConfig.Theme = c.PostForm("theme")
	AppConfig.Color = c.PostForm("color")

	conf.Write(AppConfig)

	log.Println("INFO: writing new config to", AppConfig.ConfPath)

	c.Redirect(http.StatusFound, "/config")
}

// func clearHandler(w http.ResponseWriter, r *http.Request) {

// 	log.Println("INFO: delting all hosts from DB")

// 	db.Clear(AppConfig.DbPath)

// 	http.Redirect(w, r, r.Header.Get("Referer"), 302)
// }
