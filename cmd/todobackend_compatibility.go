package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/servirtium/servirtium-go"
)

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGTERM,
	)
	servirtium := servirtium.NewServirtium()
	argsWithoutProg := os.Args[1:]
	switch argsWithoutProg[0] {
	case "record":
		go func() {
			<-sigc
			servirtium.WriteRecord("todobackend_test_suite")
			servirtium.EndRecord()
		}()
		servirtium.StartRecord("https://todo-backend-sinatra.herokuapp.com")
		break
	case "playback":
		go func() {
			<-sigc
			servirtium.EndPlayback()
		}()
		servirtium.StartPlayback("todobackend_test_suite")
		break
	default:
		log.Fatal("Oops, should have been record or playback")
	}
}
