package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type application struct {
	logFile string
}

func (a *application) generate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	id, _ := uuid.NewV7()
	go a.log(id.String())
	w.Write([]byte(id.String()))
}

func (a *application) log(id string) error {
	file, err := os.OpenFile(a.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %v\n", err)
		return err
	}
	defer file.Close()
	data := fmt.Sprintf("%s\n", id)
	if _, err := file.WriteString(data); err != nil {
		log.Fatalf("failed to write to file: %v\n", err)
		return err
	}
	return nil
}