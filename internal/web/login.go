package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ExerciseDiary/internal/auth"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

func loginHandler(c *gin.Context) {
	var guiData models.GuiData

	username := c.PostForm("username")
	password := c.PostForm("password")
	logout, ok := c.GetQuery("logout")

	if ok && logout == "yes" {

		log.Println("INFO: user logged out")
		auth.LogOut(c)

	} else if username == authConf.User && auth.MatchPasswords(authConf.Password, password) {

		log.Println("INFO: user '"+username+"' logged in. Session expire time", authConf.Expire)

		auth.StartSession(c)

	} else {

		guiData.Config = appConfig

		c.HTML(http.StatusOK, "header.html", guiData)
		c.HTML(http.StatusOK, "login.html", guiData)
	}
}
