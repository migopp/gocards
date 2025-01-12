package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/migopp/gocards/db"
	"github.com/migopp/gocards/env"
)

const GocardsAuthCookie = "GocardsAuth"

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
	c.SetCookie(GocardsAuthCookie, signedToken, 3600*24*7, "", "", false, true)
	return nil
}

func getSessionAuth(c *gin.Context) (string, error) {
	return c.Cookie(GocardsAuthCookie)
}

func getSessionUser(c *gin.Context) (db.User, error) {
	var u db.User

	// Get `GocardsAuthCookie`
	auth, err := getSessionAuth(c)
	if err != nil {
		return u, err
	}

	// Parse JWT
	tok, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(env.GCV.JWTSecret), nil
	})
	if err != nil {
		return u, err
	}

	// Get `sub` claim
	claims, _ := tok.Claims.(jwt.MapClaims)
	id, sok := claims["sub"].(uint)
	if !sok {
		return u, fmt.Errorf("Unable to fetch `sub` claim")
	}

	// Use given `id` to find the user and hand it back
	return db.GCDB.FetchUserWithID(id)
}
