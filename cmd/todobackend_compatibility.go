package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"
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
		servirtium.SetCallerRequestHeaderReplacements(map[*regexp.Regexp]string{
			regexp.MustCompile("http://localhost:61417"): "https://todo-backend-sinatra.herokuapp.com",
			regexp.MustCompile("localhost:61417"):        "todo-backend-sinatra.herokuapp.com",
		})
		servirtium.SetCallerResponseHeaderReplacements(map[*regexp.Regexp]string{
			regexp.MustCompile("https://todo-backend-sinatra.herokuapp.com"): "http://localhost:61417",
			regexp.MustCompile("todo-backend-sinatra.herokuapp.com"):         "localhost:61417",
		})
		servirtium.SetCallerResponseBodyReplacement(map[*regexp.Regexp]string{
			regexp.MustCompile("https://todo-backend-sinatra.herokuapp.com"): "http://localhost:61417",
			regexp.MustCompile("todo-backend-sinatra.herokuapp.com"):         "localhost:61417",
			regexp.MustCompile("https"):                                      "http",
		})
		servirtium.StartRecord("https://todo-backend-sinatra.herokuapp.com")
		break
	case "playback":
		go func() {
			<-sigc
			servirtium.EndPlayback()
		}()
		servirtium.SetCallerRequestHeaderReplacements(map[*regexp.Regexp]string{
			regexp.MustCompile("http://localhost:61417"): "https://todo-backend-sinatra.herokuapp.com",
			regexp.MustCompile("localhost:61417"):        "todo-backend-sinatra.herokuapp.com",
		})
		servirtium.SetCallerResponseHeaderReplacements(map[*regexp.Regexp]string{
			regexp.MustCompile("https://todo-backend-sinatra.herokuapp.com"): "http://localhost:61417",
			regexp.MustCompile("todo-backend-sinatra.herokuapp.com"):         "localhost:61417",
		})
		servirtium.SetCallerResponseBodyReplacement(map[*regexp.Regexp]string{
			regexp.MustCompile("https://todo-backend-sinatra.herokuapp.com"): "http://localhost:61417",
			regexp.MustCompile("todo-backend-sinatra.herokuapp.com"):         "localhost:61417",
			regexp.MustCompile("https"):                                      "http",
		})
		servirtium.StartPlayback("todobackend_test_suite")
		break
	default:
		log.Fatal("Oops, should have been record or playback")
	}
}
