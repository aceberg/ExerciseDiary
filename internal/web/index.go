package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/AppTemplate/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = AppConfig
	guiData.Example = "This is index page"

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "index.html", guiData)
}
