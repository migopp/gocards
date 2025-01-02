package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/migopp/gocards/db"
	"github.com/migopp/gocards/env"
)

func issueJWT(u db.User, c *gin.Context) error {
	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// Sign it
	signedToken, err := token.SignedString([]byte(env.GCV.JWTSecret))
	if err != nil {
		log.Printf("Auth generation failed: %s\n", err)
		return err
	}

	// Attach to cookie and ship
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", signedToken, 3600*24*7, "", "", false, true)
	return nil
}
