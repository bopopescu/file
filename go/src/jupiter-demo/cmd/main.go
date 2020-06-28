package main

import (
    "fmt"
	"log"
	"jupiter-demo/internal/app/engine"
	"jupiter-demo/internal/app/model"
    "jupiter-demo/internal/app/service"
)

func main() {
	eng := engine.NewEngine()
    eng.Defer(func() error {
    	fmt.Println("exit jupiter app ...")
    	return nil
    })
    model.Init()
    service.Init()
    if err := eng.Run(); err != nil {
    	log.Fatal(err)
    }
}

