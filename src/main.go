package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort    = "3000"
	defaultEnvName = "development"
)

type fixedResponse string

func (s fixedResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, s) }

func main() {
	http.Handle("/_ready", fixedResponse("~> scapegoat is ready..."))
	http.Handle("/_healthy", fixedResponse("~> scapegoat is healthy..."))
	http.Handle("/", fixedResponse(fmt.Sprintf("It's live!! env: %s\n", getEnv("ENV_NAME", defaultEnvName))))

	addr := fmt.Sprintf(":%s", getEnv("PORT", defaultPort))

	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
