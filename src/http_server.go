package src

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	reuse "github.com/libp2p/go-reuseport"
)

const ServerHTTPServeSocket = "127.0.0.1:8080"

type (
	initSessionRequest struct {
		P *uint64 `json:"p"`
		G *uint64 `json:"g"`
		Y *uint64 `json:"y"`
	}
	xReceive struct {
		X *uint64 `json:"x"`
	}
	sReceive struct {
		S *uint64 `json:"s"`
	}
)

func HTTPServer() {
	router := gin.Default()
	crypter := router.Group("schnorr_auth")
	{
		crypter.POST("init_session", initSession)
		crypter.POST("first_step", firstStep)
		crypter.POST("final_step", finalStep)
		crypter.GET("session_result", getSessionResult)
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
	sessionParammeters := initSessionRequest{
		P: &p,
		G: &g,
		Y: &y,
	}
	if err := gctx.ShouldBindJSON(&sessionParammeters); err != nil {
		log.Printf("Error on unmarshaling request body: %v", err)
		gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Открытый ключ клиента:\n P = %d\n G = %d\n Y = %d", p, g, y)
	successRounds = 0
	gctx.JSON(http.StatusOK, numberAuthenticationRounds)
}

func firstStep(gctx *gin.Context) {
	xReceive := xReceive{
		X: &x,
	}
	if err := gctx.ShouldBindJSON(&xReceive); err != nil {
		log.Printf("Error on unmarshaling request body: %v", err)
		gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Принят параметр X: %d", x)
	generateE()
	log.Printf("Сгенерирован параметр E: %d", e)
	gctx.JSON(http.StatusOK, e)
}

func finalStep(gctx *gin.Context) {
	sReceive := sReceive{
		S: &s,
	}
	if err := gctx.ShouldBindJSON(&sReceive); err != nil {
		log.Printf("Error on unmarshaling request body: %v", err)
		gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Принят праметр S: %d", s)
	roundResult := checkX()
	log.Printf("Результат раунда: %t", roundResult)
	if roundResult {
		successRounds++
	}
	gctx.JSON(http.StatusOK, roundResult)
}

func getSessionResult(gctx *gin.Context) {
	sessionResult := sessionResult()
	log.Printf("Результат сесси: %t", sessionResult)
	gctx.JSON(http.StatusOK, sessionResult)
	successRounds, p, g, y, x, e, s = 0, 0, 0, 0, 0, 0, 0
	sessionResult = false
}
