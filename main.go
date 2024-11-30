package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	// First, try to find log file name as an argument
	var logFile string
	flag.StringVar(&logFile, "log", "", "Path to the log file (shorthand: -l)")
	flag.StringVar(&logFile, "l", "", "Path to the log file")
	flag.Parse()

	// if not provided check env variable
	if logFile == "" {
		logFile = os.Getenv("IDGEN_LOG")
	}

	// if still empty then use default
	if logFile == "" {
		logFile = "generated.log"
	}

	app := &application{
		logFile: logFile,
	}
	r := chi.NewRouter()
	r.Mount("/v1", v1Router(*app))
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
