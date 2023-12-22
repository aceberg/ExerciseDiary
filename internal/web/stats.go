package web

import (
	// "log"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/db"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func statsHandler(c *gin.Context) {
	var guiData models.GuiData

	exData.Sets = db.SelectSet(appConfig.DBPath)

	guiData.Config = appConfig
	guiData.HeatMap = generateHeatMap()

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "stats.html", guiData)
}
