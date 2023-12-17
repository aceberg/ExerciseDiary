package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/db"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func exerciseHandler(c *gin.Context) {
	var guiData models.GuiData

	exData.Exs = db.SelectEx(appConfig.DBPath)

	guiData.Config = appConfig
	guiData.ExData = exData
	guiData.GroupMap = createGroupMap()

	id, _ := c.GetQuery("id")
	log.Println("ID =", id)

	// if ok && (id != new) {

	// }

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "exercise.html", guiData)
}

func saveExerciseHandler(c *gin.Context) {
	var oneEx models.Exercise

	oneEx.Group = c.PostForm("group")
	oneEx.Name = c.PostForm("name")
	oneEx.Descr = c.PostForm("descr")

	// id := c.PostForm("id")
	weight := c.PostForm("weight")
	reps := c.PostForm("reps")

	oneEx.Weight, _ = strconv.Atoi(weight)
	oneEx.Reps, _ = strconv.Atoi(reps)

	log.Println("ONEEX =", oneEx)

	db.InsertEx(appConfig.DBPath, oneEx)
	exData.Exs = db.SelectEx(appConfig.DBPath)

	c.Redirect(http.StatusFound, "/exercise/")
}
