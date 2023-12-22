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

	guiData.Config = appConfig

	guiData.HeatMap = generateHeatMap()

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "stats.html", guiData)
}

func generateHeatMap() (heatMap []models.HeatMapData) {
	var heat models.HeatMapData
	var data models.HeatPlace

	heat.Name = "Su"
	data.X = "2024-12-24"
	data.Y = 0
	heat.Data = append(heat.Data, data)
	data.X = "2024-12-17"
	data.Y = 3
	heat.Data = append(heat.Data, data)
	heatMap = append(heatMap, heat)

	heat.Data = []models.HeatPlace{}

	heat.Name = "Sa"
	data.X = "2024-12-23"
	data.Y = 1
	heat.Data = append(heat.Data, data)
	data.X = "2024-12-16"
	data.Y = 8
	heat.Data = append(heat.Data, data)
	heatMap = append(heatMap, heat)

	heat.Data = []models.HeatPlace{}

	heat.Name = "Fr"
	data.X = "2024-12-22"
	data.Y = 5
	heat.Data = append(heat.Data, data)
	data.X = "2024-12-15"
	data.Y = 2
	heat.Data = append(heat.Data, data)
	heatMap = append(heatMap, heat)

	return heatMap
}
