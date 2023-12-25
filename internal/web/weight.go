package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/db"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func weightHandler(c *gin.Context) {
	var w models.BodyWeight

	w.Date = c.PostForm("date")
	weightStr := c.PostForm("weight")

	w.Weight, _ = strconv.Atoi(weightStr)

	// log.Println("WEIGHT =", w)
	db.InsertW(appConfig.DBPath, w)

	c.Redirect(http.StatusFound, "/")
}
