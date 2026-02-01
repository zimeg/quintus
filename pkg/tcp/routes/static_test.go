package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticCSS(t *testing.T) {
	req := httptest.NewRequest("GET", "/css/output.css", nil)
	w := httptest.NewRecorder()
	StaticCSS(w, req)
	contentType := w.Result().Header.Get("Content-Type")
	assert.Equal(t, "text/css; charset=utf-8", contentType)
	body := w.Body.String()
	assert.Greater(t, len(body), 0)
	assert.Equal(t, staticCSS, body)
}

func TestStaticFaviconAppleIcon(t *testing.T) {
	req := httptest.NewRequest("GET", "/apple-touch-icon.png", nil)
	w := httptest.NewRecorder()
	StaticFaviconAppleIcon(w, req)
	contentType := w.Result().Header.Get("Content-Type")
	assert.Equal(t, "image/png", contentType)
	body := w.Body.String()
	assert.Greater(t, len(body), 0)
	assert.Equal(t, staticFaviconAppleIcon, body)
}

func TestStaticFaviconDefault(t *testing.T) {
	req := httptest.NewRequest("GET", "/favicon.ico", nil)
	w := httptest.NewRecorder()
	StaticFaviconDefault(w, req)
	contentType := w.Result().Header.Get("Content-Type")
	assert.Equal(t, "image/x-icon", contentType)
	body := w.Body.String()
	assert.Greater(t, len(body), 0)
	assert.Equal(t, staticFaviconDefault, body)
}

func TestStaticFaviconSmall(t *testing.T) {
	req := httptest.NewRequest("GET", "/favicon-32x32.png", nil)
	w := httptest.NewRecorder()
	StaticFaviconSmall(w, req)
	contentType := w.Result().Header.Get("Content-Type")
	assert.Equal(t, "image/png", contentType)
	body := w.Body.String()
	assert.Greater(t, len(body), 0)
	assert.Equal(t, staticFaviconSmall, body)
}

func TestStaticFaviconMedium(t *testing.T) {
	req := httptest.NewRequest("GET", "/favicon-192x192.png", nil)
	w := httptest.NewRecorder()
	StaticFaviconMedium(w, req)
	contentType := w.Result().Header.Get("Content-Type")
	assert.Equal(t, "image/png", contentType)
	body := w.Body.String()
	assert.Greater(t, len(body), 0)
	assert.Equal(t, staticFaviconMedium, body)
}

func TestStaticFaviconLarge(t *testing.T) {
	req := httptest.NewRequest("GET", "/favicon-512x512.png", nil)
	w := httptest.NewRecorder()
	StaticFaviconLarge(w, req)
	contentType := w.Result().Header.Get("Content-Type")
	assert.Equal(t, "image/png", contentType)
	body := w.Body.String()
	assert.Greater(t, len(body), 0)
	assert.Equal(t, staticFaviconLarge, body)
}

func TestRobots(t *testing.T) {
	req := httptest.NewRequest("GET", "/robots.txt", nil)
	w := httptest.NewRecorder()
	StaticRobots(w, req)
	contentType := w.Result().Header.Get("Content-Type")
	assert.Equal(t, "text/plain; charset=utf-8", contentType)
	body := w.Body.String()
	expected := `User-agent: *
Allow: /

Sitemap: https://quintus.sh/sitemap.xml
`
	assert.Equal(t, expected, body)
}

func TestStaticSitemap(t *testing.T) {
	req := httptest.NewRequest("GET", "/sitemap.xml", nil)
	w := httptest.NewRecorder()
	StaticSitemap(w, req)
	contentType := w.Result().Header.Get("Content-Type")
	assert.Equal(t, "application/xml; charset=utf-8", contentType)
	body := w.Body.String()
	expected := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://quintus.sh/</loc>
    <priority>1.0</priority>
  </url>
</urlset>
`
	assert.Equal(t, expected, body)
}

func TestStaticWebManifest(t *testing.T) {
	req := httptest.NewRequest("GET", "/manifest.webmanifest", nil)
	w := httptest.NewRecorder()
	StaticWebManifest(w, req)
	contentType := w.Result().Header.Get("Content-Type")
	assert.Equal(t, "application/manifest+json", contentType)
	body := w.Body.String()
	expected := `{
  "$schema": "https://raw.githubusercontent.com/SchemaStore/schemastore/refs/heads/master/src/schemas/json/web-manifest.json",
  "name": "Quintus Calendars",
  "start_url": "https://quintus.sh",
  "short_name": "Quintus",
  "icons": [
    {
      "src": "/favicon-192x192.png",
      "sizes": "192x192",
      "type": "image/png"
    },
    {
      "src": "/favicon-512x512.png",
      "sizes": "512x512",
      "type": "image/png"
    }
  ],
  "theme_color": "#ffffff",
  "background_color": "#ffffff",
  "display": "standalone"
}
`
	assert.Equal(t, expected, body)
}
