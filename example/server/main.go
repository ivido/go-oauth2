package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/ivido/go-oauth2/server"
)

var (
	verbose      bool
	port         int
	clientId     string
	clientSecret string
	clientRedir  string
)

func init() {
	flag.BoolVar(&verbose, "v", false, "Whether output should be verbose")
	flag.IntVar(&port, "p", 9093, "Port to run the server on")
	flag.StringVar(&clientId, "i", "testclient", "The client ID of the client to configure")
	flag.StringVar(&clientSecret, "s", "testsecret", "The secret of the client to configure")
	flag.StringVar(&clientRedir, "r", "http://localhost:9094", "The redirect URL of the client to configure")
}

func main() {
	flag.Parse()

	if verbose {
		log.Println("Enabled verbose output")
	}

	srv := server.NewDefaultServer()

	http.HandleFunc("/oauth/authorize", func(w http.ResponseWriter, r *http.Request) {
		if verbose {
			dumpRequest(os.Stdout, "authorize", r)
		}

		err := srv.HandleAuthorizationRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		if verbose {
			dumpRequest(os.Stdout, "token", r)
		}

		http.Error(w, "Token requests unsupported", http.StatusNotImplemented)
	})

	log.Printf("Server is running at port %d\n", port)
	log.Printf("Point your OAuth client Auth endpoint to %s:%d%s\n", "http://localhost", port, "/oauth/authorize")
	log.Printf("Point your OAuth client Token endpoint to %s:%d%s\n", "http://localhost", port, "/oauth/token")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func dumpRequest(writer io.Writer, header string, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	writer.Write([]byte("\n" + header + ": \n"))
	writer.Write(data)
	return nil
}
