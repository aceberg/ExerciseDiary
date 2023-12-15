package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig
	guiData.ExData = exData
	guiData.GroupMap = createGroupMap()

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "index.html", guiData)
}

func createGroupMap() map[string]string {

	grMap := make(map[string]string)
	i := 0

	for _, ex := range exData.Exs {

		_, ok := grMap[ex.Group]
		if !ok {
			grMap[ex.Group] = "grID" + fmt.Sprintf("%d", i)
			i = i + 1
		}
	}
	// fmt.Println("GRMAP", grMap)

	return grMap
}
