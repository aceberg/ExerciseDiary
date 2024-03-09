package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setTokenCookie(c *gin.Context, token string) {

	cookie := http.Cookie{Name: cookieName, Value: token, Path: "/"}
	http.SetCookie(c.Writer, &cookie)
}

func getTokenFromCookie(c *gin.Context) string {

	cookie, err := c.Request.Cookie(cookieName)
	if err != nil {

		return ""
	}

	return cookie.Value
}
