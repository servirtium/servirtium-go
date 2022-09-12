package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
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
	realUrl := argsWithoutProg[1]
	realDomain := strings.Replace(strings.Replace(argsWithoutProg[1], "https://", "", 1), "http://", "", 1)
	fmt.Fprintf(os.Stdout, "real URL= "+realUrl+"\n")
	fmt.Fprintf(os.Stdout, "real domain= "+realDomain+"\n")
	switch argsWithoutProg[0] {
	case "record":
		go func() {
			<-sigc
			servirtium.WriteRecord("todobackend_test_suite")
			servirtium.EndRecord()
		}()
		servirtium.SetCallerRequestHeaderReplacements(map[*regexp.Regexp]string{
			regexp.MustCompile("http://localhost:61417"): realUrl,
			regexp.MustCompile("localhost:61417"):        realDomain,
		})
		servirtium.SetCallerResponseHeaderReplacements(map[*regexp.Regexp]string{
			regexp.MustCompile(realUrl):    "http://localhost:61417",
			regexp.MustCompile(realDomain): "localhost:61417",
		})
		servirtium.SetCallerResponseBodyReplacement(map[*regexp.Regexp]string{
			regexp.MustCompile(realUrl):    "http://localhost:61417",
			regexp.MustCompile(realDomain): "localhost:61417",
			regexp.MustCompile("https"):    "http",
		})
		servirtium.StartRecord(realDomain, 61417)
		break
	case "playback":
		go func() {
			<-sigc
			servirtium.EndPlayback()
		}()
		servirtium.SetCallerRequestHeaderReplacements(map[*regexp.Regexp]string{
			regexp.MustCompile("http://localhost:61417"): realUrl,
			regexp.MustCompile("localhost:61417"):        realDomain,
		})
		servirtium.SetCallerResponseHeaderReplacements(map[*regexp.Regexp]string{
			regexp.MustCompile(realUrl):    "http://localhost:61417",
			regexp.MustCompile(realDomain): "localhost:61417",
		})
		servirtium.SetCallerResponseBodyReplacement(map[*regexp.Regexp]string{
			regexp.MustCompile(realUrl):    "http://localhost:61417",
			regexp.MustCompile(realDomain): "localhost:61417",
			regexp.MustCompile("https"):    "http",
		})
		servirtium.StartPlayback("todobackend_test_suite", 61417)
		break
	default:
		log.Fatal("Oops, should have been record or playback")
	}
}
