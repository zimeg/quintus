package tcp

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zimeg/quintus/cicero/pkg/tcp/routes"
)

// Listen responds to incoming requests from the web with timing information
func Listen() {
	var port string
	if v, ok := os.LookupEnv("HTTP_PORT"); !ok {
		port = ":80"
	} else {
		port = fmt.Sprintf(":%s", v)
	}
	log.Printf("TCP server handling the HTTP requests on port %s", port)
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/cal/{year...}", routes.Cal)
	http.HandleFunc("/now", routes.Now)
	http.HandleFunc("/utc", routes.UTC)
	log.Fatal(http.ListenAndServe(port, nil))
}
