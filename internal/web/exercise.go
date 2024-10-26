package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"github.com/aceberg/ExerciseDiary/internal/db"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func exerciseHandler(c *gin.Context) {
	var guiData models.GuiData
	var id int

	exData.Exs = db.SelectEx(appConfig.DBPath)

	guiData.Config = appConfig
	guiData.ExData = exData
	guiData.GroupMap = createGroupMap()

	idStr, ok := c.GetQuery("id")
	// log.Println("ID =", idStr)

	if ok && (idStr != "new") {
		id, _ = strconv.Atoi(idStr)

		for _, oneEx := range exData.Exs {
			if oneEx.ID == id {
				guiData.OneEx = oneEx
				break
			}
		}
	}

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "exercise.html", guiData)
}

func saveExerciseHandler(c *gin.Context) {
	var oneEx models.Exercise

	oneEx.Group = c.PostForm("group")
	oneEx.Place = c.PostForm("place")
	oneEx.Name = c.PostForm("name")
	oneEx.Descr = c.PostForm("descr")
	oneEx.Image = c.PostForm("image")

	id := c.PostForm("id")
	weight := c.PostForm("weight")
	reps := c.PostForm("reps")

	oneEx.ID, _ = strconv.Atoi(id)
	oneEx.Weight, _ = decimal.NewFromString(weight)
	oneEx.Reps, _ = strconv.Atoi(reps)

	// log.Println("ONEEX =", oneEx)

	if oneEx.ID != 0 {
		db.DeleteEx(appConfig.DBPath, oneEx.ID)
	}

	db.InsertEx(appConfig.DBPath, oneEx)
	exData.Exs = db.SelectEx(appConfig.DBPath)

	c.Redirect(http.StatusFound, "/")
}

func deleteExerciseHandler(c *gin.Context) {

	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)

	db.DeleteEx(appConfig.DBPath, id)
	exData.Exs = db.SelectEx(appConfig.DBPath)

	c.Redirect(http.StatusFound, "/")
}
