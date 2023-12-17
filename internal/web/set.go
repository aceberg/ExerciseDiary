package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/db"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func setHandler(c *gin.Context) {

	var formData []models.Set
	var oneSet models.Set
	var weight, reps int

	_ = c.PostFormMap("sets")

	formMap := c.Request.PostForm
	log.Println("MAP:", formMap)

	len := len(formMap["name"])
	log.Println("LEN:", len)

	for i := 0; i < len; i++ {
		// id, _ = strconv.Atoi(formMap["id"][i])
		// oneSet.ID = id
		oneSet.Date = formMap["date"][0]
		oneSet.Name = formMap["name"][i]
		weight, _ = strconv.Atoi(formMap["weight"][i])
		reps, _ = strconv.Atoi(formMap["reps"][i])
		oneSet.Weight = weight
		oneSet.Reps = reps

		formData = append(formData, oneSet)
		// exData.Sets = append(exData.Sets, oneSet)
		// db.InsertSet(appConfig.DBPath, oneSet)
	}

	db.BulkDeleteSetsByDate(appConfig.DBPath, oneSet.Date)
	db.BulkAddSets(appConfig.DBPath, formData)
	exData.Sets = db.SelectSet(appConfig.DBPath)

	log.Println("FORM DATA:", formData)

	c.Redirect(http.StatusFound, "/")
}
