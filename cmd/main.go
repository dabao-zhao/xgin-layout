package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	xconfig "github.com/dabao-zhao/xgin/config"
	"github.com/dabao-zhao/xgin/pkg/xfile"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/dabao-zhao/xgin-layout/config"
	"github.com/dabao-zhao/xgin-layout/middleware"
	"github.com/dabao-zhao/xgin-layout/router"
)

var (
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "./config/dev.toml", "config path, eg: -conf config.toml")
}

func main() {
	conf := config.Config{}
	if err := xconfig.Load(flagConf, &conf); err != nil {
		panic(err)
	}

	r := gin.New()

	r.Use(gin.Recovery())

	if conf.Http.AccessLogFile != "" {
		f, err := xfile.CreateAndOpen(conf.Http.AccessLogFile)
		if err != nil {
			panic(err)
		}
		gin.DefaultWriter = io.MultiWriter(f)
	}

	r.Use(middleware.AccessLogFormatter())

	if conf.Http.CorsEnable {
		r.Use(cors.New(cors.Config{
			AllowAllOrigins:  conf.Http.CorsAllowAllOrigins,
			AllowOrigins:     conf.Http.CorsAllowOrigins,
			AllowMethods:     conf.Http.CorsAllowMethods,
			AllowHeaders:     conf.Http.CorsAllowHeaders,
			AllowCredentials: conf.Http.CorsAllowCredentials,
			ExposeHeaders:    conf.Http.CorsExposeHeaders,
			MaxAge:           time.Duration(conf.Http.CorsMaxAge * int64(time.Second)),
		}))
	}

	router.Register(r, &conf)

	addr := fmt.Sprintf("%s:%d", conf.Http.Host, conf.Http.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}

		log.Println("Server exiting")
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
