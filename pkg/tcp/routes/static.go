package routes

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed static/css/output.css
var staticCSS string

// StaticCSS returns the compiled tailwind stylesheet
func StaticCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	fmt.Fprintf(w, "%s", staticCSS)
}

//go:embed static/favicon/apple-touch-icon.png
var staticFaviconAppleIcon string

// StaticFaviconAppleIcon contains a retina preview
func StaticFaviconAppleIcon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	fmt.Fprintf(w, "%s", staticFaviconAppleIcon)
}

//go:embed static/favicon/favicon.ico
var staticFaviconDefault string

// StaticFaviconDefault has the default favicon
func StaticFaviconDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")
	fmt.Fprintf(w, "%s", staticFaviconDefault)
}

//go:embed static/favicon/favicon-32x32.png
var staticFaviconSmall string

// StaticFaviconSmall has the small favicon
func StaticFaviconSmall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	fmt.Fprintf(w, "%s", staticFaviconSmall)
}

//go:embed static/favicon/favicon-192x192.png
var staticFaviconMedium string

// StaticFaviconMedium has the medium favicon
func StaticFaviconMedium(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	fmt.Fprintf(w, "%s", staticFaviconMedium)
}

//go:embed static/favicon/favicon-512x512.png
var staticFaviconLarge string

// StaticFaviconLarge has the large favicon
func StaticFaviconLarge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	fmt.Fprintf(w, "%s", staticFaviconLarge)
}

//go:embed static/robots.txt
var staticRobots string

// StaticRobots lets machines read these pages
func StaticRobots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "%s", staticRobots)
}

//go:embed static/sitemap.xml
var staticSitemap string

// StaticSitemap returns the path to pages known
func StaticSitemap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	fmt.Fprintf(w, "%s", staticSitemap)
}

//go:embed static/site.webmanifest
var staticWebManifest string

// StaticWebManifest contains the web manifest
func StaticWebManifest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/manifest+json")
	fmt.Fprintf(w, "%s", staticWebManifest)
}
