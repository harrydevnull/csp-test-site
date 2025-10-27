package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Get environment variables or set defaults
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	// Get the current directory for serving static files
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Define handlers
	http.HandleFunc("/", homeHandler(wd))

	// Login endpoints (login1-login19)
	http.HandleFunc("/login1", loginHandler(wd, "script-src https: 'self' ; object-src 'none'"))
	http.HandleFunc("/login2", loginHandler(wd, "default-src https: 'self' ; object-src 'none'"))
	http.HandleFunc("/login3", loginHandler(wd, "script-src https: 'unsafe-inline' ; object-src 'none'"))
	http.HandleFunc("/login4", loginHandler(wd, "default-src https: 'unsafe-inline' ; object-src 'none'"))
	http.HandleFunc("/login5", loginHandler(wd, "script-src https: 'unsafe-inline' 'strict-dynamic' ; object-src 'none'"))
	http.HandleFunc("/login6", loginHandler(wd, "default-src https: 'unsafe-inline' 'strict-dynamic' ; object-src 'none'"))
	http.HandleFunc("/login7", loginHandler(wd, "default-src https: 'self'; script-src https: 'self'; object-src 'none'"))
	http.HandleFunc("/login8", loginHandler(wd, "default-src https: 'unsafe-inline' 'self'; script-src https: 'self'; object-src 'none'"))
	http.HandleFunc("/login9", loginHandler(wd, "default-src https: 'self'; script-src https: 'unsafe-inline'; object-src 'none'"))
	http.HandleFunc("/login10", loginHandler(wd, "default-src https: 'unsafe-inline'; script-src https: 'unsafe-inline'; object-src 'none'"))
	http.HandleFunc("/login11", loginHandler(wd, "default-src https: 'unsafe-inline' 'strict-dynamic' ; script-src https: 'self'; object-src 'none'"))
	http.HandleFunc("/login12", loginHandler(wd, "default-src https: 'self'; script-src https: 'unsafe-inline'  'strict-dynamic' ; object-src 'none'"))
	http.HandleFunc("/login13", loginHandler(wd, "script-src https: 'unsafe-inline' 'unsafe-eval' ; object-src 'none'"))
	http.HandleFunc("/login14", loginHandler(wd, "script-src https: 'unsafe-inline' 'unsafe-eval'  'strict-dynamic' ; object-src 'none'"))
	http.HandleFunc("/login15", loginHandler(wd, "style-src 'self'; object-src 'none'"))
	http.HandleFunc("/login16", loginHandler(wd, "script-src https: 'unsafe-inline' 'self' ; object-src 'none'"))
	http.HandleFunc("/login17", loginHandler(wd, "default-src https: 'unsafe-inline' 'self' ; object-src 'none'"))
	http.HandleFunc("/login18", loginHandler(wd, "default-src https: 'unsafe-inline' 'strict-dynamic'; script-src https: 'unsafe-inline'; object-src 'none'"))
	http.HandleFunc("/login19", loginHandler(wd, "default-src https: 'unsafe-inline' ; script-src https: 'unsafe-inline' 'strict-dynamic' ; object-src 'none'"))

	// Login endpoints with nonce (login1x-login19x)
	http.HandleFunc("/login1x", loginHandler(wd, "script-src  https: 'nonce-dcd7d07a8645fca5bfc7ed3eflogin1x' 'self' ; object-src 'none'"))
	http.HandleFunc("/login2x", loginHandler(wd, "default-src  https: 'nonce-dcd7d07a8645fca5bfc7ed3eflogin2x' 'self' ; object-src 'none'"))
	http.HandleFunc("/login3x", loginHandler(wd, "script-src https: 'nonce-dcd7d07a8645fca5bfc7ed3eflogin3x' 'unsafe-inline' ; object-src 'none'"))
	http.HandleFunc("/login4x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3eflogin4x' 'unsafe-inline' ; object-src 'none'"))
	http.HandleFunc("/login5x", loginHandler(wd, "script-src https: 'nonce-dcd7d07a8645fca5bfc7ed3eflogin5x' 'unsafe-inline' 'strict-dynamic' ; object-src 'none'"))
	http.HandleFunc("/login6x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3eflogin6x' 'unsafe-inline' 'strict-dynamic' ; object-src 'none'"))
	http.HandleFunc("/login7x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3eflogin7x' 'self'; script-src https: 'self'; object-src 'none'"))
	http.HandleFunc("/login8x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3eflogin8x' 'unsafe-inline' 'self'; script-src https: 'self'; object-src 'none'"))
	http.HandleFunc("/login9x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3eflogin9x' 'self'; script-src https: 'unsafe-inline'; object-src 'none'"))
	http.HandleFunc("/login10x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3elogin10x' 'unsafe-inline'; script-src https: 'unsafe-inline'; object-src 'none'"))
	http.HandleFunc("/login11x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3elogin11x' 'unsafe-inline' 'strict-dynamic' ; script-src https: 'self'; object-src 'none'"))
	http.HandleFunc("/login12x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3elogin12x' 'self'; script-src https: 'unsafe-inline'  'strict-dynamic' ; object-src 'none'"))
	http.HandleFunc("/login13x", loginHandler(wd, "script-src https: 'nonce-dcd7d07a8645fca5bfc7ed3elogin13x' 'unsafe-inline' 'unsafe-eval' ; object-src 'none'"))
	http.HandleFunc("/login14x", loginHandler(wd, "script-src https: 'nonce-dcd7d07a8645fca5bfc7ed3elogin14x' 'unsafe-inline' 'unsafe-eval'  'strict-dynamic' ; object-src 'none'"))
	http.HandleFunc("/login15x", loginHandler(wd, "style-src 'nonce-dcd7d07a8645fca5bfc7ed3elogin15x' 'self'; object-src 'none'"))
	http.HandleFunc("/login16x", loginHandler(wd, "script-src https: 'nonce-dcd7d07a8645fca5bfc7ed3elogin16x' 'unsafe-inline' 'self' ; object-src 'none'"))
	http.HandleFunc("/login17x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3elogin17x' 'unsafe-inline' 'self' ; object-src 'none'"))
	http.HandleFunc("/login18x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3elogin18x' 'unsafe-inline' 'strict-dynamic'; script-src https: 'unsafe-inline'; object-src 'none'"))
	http.HandleFunc("/login19x", loginHandler(wd, "default-src https: 'nonce-dcd7d07a8645fca5bfc7ed3elogin19x' 'unsafe-inline' ; script-src https: 'unsafe-inline' 'strict-dynamic' ; object-src 'none'"))

	// CSP test endpoints
	http.HandleFunc("/csp1", loginHandler(wd, "script-src https: 'unsafe-eval'; object-src 'none'"))
	http.HandleFunc("/csp2", loginHandler(wd, "script-src 'self'; object-src 'none'"))
	http.HandleFunc("/csp8", cspHandlerWithCache(wd, "script-src 'self'; object-src 'none'"))
	http.HandleFunc("/csp10", cspHandlerWithCacheOnly(wd))
	http.HandleFunc("/csp11", cspHandlerWithCache(wd, "style-src 'self'; object-src 'none'"))
	http.HandleFunc("/csp12", cspHandlerWithCache(wd, "img-src 'self'; script-src 'self'; object-src 'none'"))
	http.HandleFunc("/csp13", cspHandlerWithCache(wd, "img-src 'self'; script-src 'strict-dynamic' 'self'; object-src 'none'"))
	http.HandleFunc("/csp14", cspHandlerWithCache(wd, "img-src 'self'; script-src 'unsafe-inline' 'self'; object-src 'none'"))
	http.HandleFunc("/csp15", cspHandlerWithCache(wd, "img-src 'self'; script-src 'strict-dynamic' 'unsafe-inline' 'self'; object-src 'none'"))
	http.HandleFunc("/csp16", cspHandlerWithCache(wd, "img-src 'self'; default-src 'self'; object-src 'none'"))
	http.HandleFunc("/csp17", cspHandlerWithCache(wd, "img-src 'self'; default-src 'strict-dynamic' 'self'; object-src 'none'"))
	http.HandleFunc("/csp18", cspHandlerWithCache(wd, "img-src 'self'; default-src 'unsafe-inline' 'self'; object-src 'none'"))
	http.HandleFunc("/csp19", cspHandlerWithCache(wd, "img-src 'self'; default-src 'strict-dynamic' 'unsafe-inline' 'self'; object-src 'none'"))
	http.HandleFunc("/csp20", cspHandlerWithCache(wd, "default-src 'self'; script-src 'self'; object-src 'none'"))
	http.HandleFunc("/csp21", cspHandlerWithCache(wd, "default-src 'self'; script-src 'strict-dynamic' 'self'; object-src 'none'"))
	http.HandleFunc("/csp22", cspHandlerWithCache(wd, "default-src 'self'; script-src 'unsafe-inline' 'self'; object-src 'none'"))
	http.HandleFunc("/csp23", cspHandlerWithCache(wd, "default-src 'self'; script-src 'strict-dynamic' 'unsafe-inline' 'self'; object-src 'none'"))
	http.HandleFunc("/csp24", cspHandlerWithCache(wd, "img-src 'self'; ScRiPt-SrC 'self'; object-src 'none'"))
	http.HandleFunc("/csp25", cspHandlerWithCache(wd, "img-src 'self'; script-src 'UnSaFe-InLiNe' 'self'; object-src 'none'"))
	http.HandleFunc("/csp26", cspHandlerWithCache(wd, "img-src 'self'; DeFaUlT-SrC 'self'; object-src 'none'"))
	http.HandleFunc("/csp27", cspHandlerWithCache(wd, "img-src 'self'; default-src 'UnSaFe-InLiNe' 'self'; object-src 'none'"))
	http.HandleFunc("/csp28", cspHandlerWithCustomHeader(wd, "CONtENT-SECUrITY-POLiCY", "script-src 'self'; object-src 'none'"))
	http.HandleFunc("/csp40", loginHandler(wd, "script-src https: 'unsafe-eval' ; object-src 'none'"))
	http.HandleFunc("/csp41", loginHandler(wd, "script-src https: 'strict-dynamic' 'unsafe-eval' ; object-src 'none'"))
	http.HandleFunc("/csp42", cspHandlerWithDifferentFile(wd, "script-src 'unsafe-inline' ; object-src 'none'", "index1.html"))
	http.HandleFunc("/csp43", cspHandlerWithDifferentFile(wd, "script-src 'nonce-5843b9e95a432059d901fb701597edf0' ; object-src 'none'", "index1.html"))
	http.HandleFunc("/csp44", cspHandlerWithDifferentFile(wd, "script-src 'nonce-5843b9e95a432059d901fb701597edf0' 'unsafe-inline' ; object-src 'none'", "index1.html"))
	http.HandleFunc("/csp45", cspHandlerWithCache(wd, "img-src 'self'; default-src 'nonce-5843b9e95a432059d901fb701597edf0'; object-src 'none'"))
	http.HandleFunc("/csp46", cspHandlerWithCache(wd, "img-src 'self'; default-src 'nonce-5843b9e95a432059d901fb701597edf0' 'unsafe-inline'; object-src 'none'"))

	// Start server
	addr := host + ":" + port
	fmt.Printf("csp-test-site app listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// homeHandler serves the index.html with basic CSP
func homeHandler(workingDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "script-src 'self'; object-src 'none'")
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, filepath.Join(workingDir, "index.html"))
	}
}

// loginHandler creates a handler with specific CSP policy
func loginHandler(workingDir, cspPolicy string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", cspPolicy)
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, filepath.Join(workingDir, "index.html"))
	}
}

// cspHandlerWithCache creates a handler with CSP policy and no-store cache control
func cspHandlerWithCache(workingDir, cspPolicy string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", cspPolicy)
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, filepath.Join(workingDir, "index.html"))
	}
}

// cspHandlerWithCacheOnly creates a handler with only cache control (no CSP)
func cspHandlerWithCacheOnly(workingDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, filepath.Join(workingDir, "index.html"))
	}
}

// cspHandlerWithCustomHeader creates a handler with custom header name
func cspHandlerWithCustomHeader(workingDir, headerName, cspPolicy string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerName, cspPolicy)
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, filepath.Join(workingDir, "index.html"))
	}
}

// cspHandlerWithDifferentFile creates a handler that serves a different HTML file
func cspHandlerWithDifferentFile(workingDir, cspPolicy, filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", cspPolicy)
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, filepath.Join(workingDir, filename))
	}
}
