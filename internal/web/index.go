package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/db"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData

	// exData.Exs = db.SelectEx(appConfig.DBPath)
	exData.Sets = db.SelectSet(appConfig.DBPath)

	guiData.Config = appConfig
	guiData.ExData = exData
	guiData.GroupMap = createGroupMap()
	guiData.Today = time.Now().Format("2006-01-02")

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

	return grMap
}
