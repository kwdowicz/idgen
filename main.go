package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	// First, try to find log and port file name as an argument
	var logFile string
	var port string 
	flag.StringVar(&logFile, "l", "", "Path to the log file")
	flag.StringVar(&port, "p", "", "Port")

	flag.Parse()

	// if not provided check env variable
	if logFile == "" {
		logFile = os.Getenv("IDGEN_LOG")
	}

	if port == "" {
		port = os.Getenv("IDGEN_PORT")
	}

	// if still empty then use default
	if logFile == "" {
		logFile = "generated.log"
	}

	if port == "" {
		port = "7777"
	}

	app := &application{
		logFile: logFile,
	}
	r := chi.NewRouter()
	r.Mount("/v1", v1Router(*app))
	log.Printf("Server is listening on port %s...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
