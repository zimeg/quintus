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
		port = ":5000"
	} else {
		port = fmt.Sprintf(":%s", v)
	}
	log.Printf("TCP server handling the HTTP requests on port %s", port)
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/apple-touch-icon.png", routes.StaticFaviconAppleIcon)
	http.HandleFunc("/cal/{year...}", routes.Cal)
	http.HandleFunc("/css/output.css", routes.StaticCSS)
	http.HandleFunc("/date/{date}", routes.Date)
	http.HandleFunc("/favicon.ico", routes.StaticFaviconDefault)
	http.HandleFunc("/favicon-32x32.png", routes.StaticFaviconSmall)
	http.HandleFunc("/favicon-192x192.png", routes.StaticFaviconMedium)
	http.HandleFunc("/favicon-512x512.png", routes.StaticFaviconLarge)
	http.HandleFunc("/manifest.webmanifest", routes.StaticWebManifest)
	http.HandleFunc("/now", routes.Now)
	http.HandleFunc("/robots.txt", routes.StaticRobots)
	http.HandleFunc("/sitemap.xml", routes.StaticSitemap)
	http.HandleFunc("/utc", routes.UTC)
	log.Fatal(http.ListenAndServe(port, nil))
}
