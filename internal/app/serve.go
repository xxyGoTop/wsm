package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/xxyGoTop/wsm/internal/app/config"
)

func Serve() error {
	port := config.Http.Port

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        RootRouter,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("listen server on port is %s\n", s.Addr)

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	log.Println("Server stopped")

	os.Exit(o)

	return nil
}
