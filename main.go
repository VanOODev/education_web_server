package main

import (
	"github.com/VanOODev/education_web_server/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatalln(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-sigChan:
			log.Printf("Received signal: %s", sig)
			a.Close()
			os.Exit(0)
		}
	}
	// В main не должно быть никакой логики, кроме запуска приложения и обработки сигналов
}
