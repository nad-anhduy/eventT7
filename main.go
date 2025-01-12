package main

import (
	"eventT7/common"
	"eventT7/config"
	"eventT7/modules/Event/middleware"
	"eventT7/modules/Event/transport"
	"fmt"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	path := fmt.Sprintf("./log/services-%v.log", common.GetDateLog())

	fileLog, err := common.OpenLogFile(path)
	if err != nil {
		log.Fatalf(`Open file log fail with error: %v`, err)
	}
	log.SetOutput(fileLog)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	log.Println(`========================================`)
	log.Printf(`PID: %d `, os.Getpid())

	envProject, err := config.Load(`confFile`)
	if err != nil {
		log.Fatalf(`Get environment fail with error: %v`, err)
	}

	sessionReadPG, err := common.GetSessionPG(envProject.Env.PGConnRO)
	defer func() { closeDB, _ := sessionReadPG.DB(); closeDB.Close() }()
	if err != nil {
		log.Fatalf(`Connect Postgres fail with error: %v`, err)
	}

	sessionWritePG, err := common.GetSessionPG(envProject.Env.PGConnRW)
	defer func() { closeDB, _ := sessionWritePG.DB(); closeDB.Close() }()
	if err != nil {
		log.Fatalf(`Connect Postgres fail with error: %v`, err)
	}

	err = common.DBInit(sessionWritePG)
	if err != nil {
		log.Fatalf(`Create table fail with error: %v`, err)
	}

	mainStream := gin.Default()
	mainStreamCache := persist.NewMemoryStore(1 * time.Minute)
	mainStream.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://cohon-itc.mservice.com.vn", "https://cohon-itc.mservice.com.vn/"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "X-CSRF-Token", "Host", "Token", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
		AllowAllOrigins:  false,
		MaxAge:           86400,
	}))

	r := mainStream.Group("/soul/api/v1")
	{
		r.OPTIONS("/checkVote", func(c *gin.Context) { c.AbortWithStatus(204) })
		r.GET("/checkVote", cache.CacheByRequestURI(mainStreamCache, 30*time.Second), transport.HandlerCheckAmountGroupId(sessionReadPG))
		r.OPTIONS("/checkTop", func(c *gin.Context) { c.AbortWithStatus(204) })
		r.GET("/checkTop", cache.CacheByRequestURI(mainStreamCache, 30*time.Second), transport.HandlerCheckTop(sessionReadPG))
		r.POST("/deny", transport.HandlerDenyGroupId())
		r.OPTIONS("/increase", func(c *gin.Context) { c.AbortWithStatus(204) })
		r.POST("/increase", middleware.ProxyFilterAll(`127.0.0.1`), transport.HanlderCreateRecord(sessionWritePG))

		r.POST("/decrease", transport.HanlderUpdateRecord(sessionWritePG))
		r.GET("/status", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": true}) })
	}

	err = mainStream.Run(":8001")
	if err != nil {
		log.Fatalf(`Run service fail with error: %v`, err)
	}
}
