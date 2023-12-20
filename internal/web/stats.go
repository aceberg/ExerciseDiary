package web

import (
	// "log"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"

	// "github.com/aceberg/ExerciseDiary/internal/check"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func statsHandler(c *gin.Context) {
	var guiData models.GuiData
	var heat models.HeatMapData

	guiData.Config = appConfig

	heat.Name = "Su"
	heat.Data = []int{1, 2, 3, 0, 5}
	guiData.HeatMap = append(guiData.HeatMap, heat)
	heat.Name = "Sa"
	heat.Data = []int{3, 0, 1, 4, 2}
	guiData.HeatMap = append(guiData.HeatMap, heat)

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "stats.html", guiData)
}
