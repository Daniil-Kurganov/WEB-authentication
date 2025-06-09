package src

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	reuse "github.com/libp2p/go-reuseport"
)

const (
	numberAuthenticationRounds = 3
	ServerHTTPServeSocket      = "127.0.0.1:8080"
)

type initSessionRequest struct {
	P *uint64 `json:"p"`
	G *uint64 `json:"g"`
	Y *uint64 `json:"y"`
}

func HTTPServer() {
	router := gin.Default()
	crypter := router.Group("schnorr_auth")
	{
		crypter.POST("init_session", initSession)
	}
	var listener net.Listener
	var err error
	if listener, err = reuse.Listen("tcp", ServerHTTPServeSocket); err != nil {
		log.Fatalf("Error on creating listener: %s", err)
	}
	if err = router.RunListener(listener); err != nil {
		log.Fatalf("Error on starting HTTP-server: %s", err)
	}
}

func initSession(gctx *gin.Context) {
	inS := initSessionRequest{
		P: &p,
		G: &g,
		Y: &y,
	}
	if err := gctx.ShouldBindJSON(&inS); err != nil {
		log.Printf("Error on unmarshaling request body: %v", err)
		gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Open key of client:\n P = %d\n G = %d\n Y = %d", p, g, y)
	gctx.JSON(http.StatusOK, numberAuthenticationRounds)
}
