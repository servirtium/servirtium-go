package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	servirtiumPackage "github.com/servirtium/servirtium-go"
)

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGTERM,
	)
	servirtium := servirtiumPackage.NewServirtium()
	argsWithoutProg := os.Args[1:]
	switch argsWithoutProg[0] {
	case "record":
		go func() {
			<-sigc
			servirtium.WriteRecord("todobackend_test_suite")
			servirtium.EndRecord()
		}()
		serverRecord := servirtiumPackage.NewServerRecord(servirtium.ManInTheMiddleHandler("https://todo-backend-sinatra.herokuapp.com"), 61417)
		servirtium.ServerRecord = serverRecord
		servirtium.StartRecord()
		break
	case "playback":
		go func() {
			<-sigc
			servirtium.EndPlayback()
		}()
		serverPlayback := servirtiumPackage.NewServerPlayback(servirtium.AnualAvgHandlerPlayback("todobackend_test_suite"), 61417)
		servirtium.ServerPlayback = serverPlayback
		servirtium.StartPlayback()
		break
	default:
		log.Fatal("Oops, should have been record or playback")
	}
}
