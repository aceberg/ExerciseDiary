package web

import (
	// "log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"github.com/aceberg/ExerciseDiary/internal/db"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func addWeightHandler(c *gin.Context) {
	var w models.BodyWeight

	w.Date = c.PostForm("date")
	weightStr := c.PostForm("weight")

	w.Weight, _ = decimal.NewFromString(weightStr)

	db.InsertW(appConfig.DBPath, w)

	c.Redirect(http.StatusFound, c.Request.Header["Referer"][0])
}

func weightHandler(c *gin.Context) {
	var guiData models.GuiData

	idStr, ok := c.GetQuery("del")
	if ok {
		id, _ := strconv.Atoi(idStr)
		db.DeleteW(appConfig.DBPath, id)
	}
	exData.Weight = db.SelectW(appConfig.DBPath)

	guiData.Config = appConfig
	guiData.ExData = exData

	// Sort weight by Date
	sort.Slice(guiData.ExData.Weight, func(i, j int) bool {
		return guiData.ExData.Weight[i].Date < guiData.ExData.Weight[j].Date
	})

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "weight.html", guiData)
}
