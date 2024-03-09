package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

// StartSession for new login
func StartSession(c *gin.Context) {

	sessionToken := uuid.NewString()

	allSessions[sessionToken] = time.Now().Add(60 * time.Second)

	setTokenCookie(c, sessionToken)

	c.Redirect(http.StatusFound, "/")
}

// LogOut - log out
func LogOut(c *gin.Context) {

	sessionToken := getTokenFromCookie(c)

	delete(allSessions, sessionToken)

	setTokenCookie(c, "")

	c.Redirect(http.StatusFound, "/")
}
