package main

import (
	"log"
	"net/http"
	"time"


	v1	"clientSystem/router"

	"golang.org/x/sync/errgroup"

)

var (
	g errgroup.Group
)

func main() {

	// route 1
	v1.InitRouter()



	srv1 := &http.Server{
		Addr:		":8080",
		Handler:	v1.R,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := srv1.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}


}
