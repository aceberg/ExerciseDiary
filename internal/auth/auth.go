package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth - main auth func
func Auth(conf *Conf) gin.HandlerFunc {
	return func(c *gin.Context) {

		if !conf.Auth || conf.User == "" || conf.Password == "" {
			c.Next()
			return
		}

		sessionToken := getTokenFromCookie(c)

		userSession, exists := allSessions[sessionToken]
		if !exists {
			c.Redirect(http.StatusFound, "/login/")
			return
		}
		if userSession.Before(time.Now()) {
			delete(allSessions, sessionToken)
			c.Redirect(http.StatusFound, "/login/")
			return
		}

		userSession = time.Now().Add(conf.Expire)
		allSessions[sessionToken] = userSession

		c.Next()
	}
}
