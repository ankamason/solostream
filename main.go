package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	port         string
	jwtSecret    string
	platform     string
	filepathRoot string
	assetsRoot   string
	s3Bucket     string
	s3Region     string
	s3CfDistro   string
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	cfg := apiConfig{
		port:         port,
		jwtSecret:    os.Getenv("JWT_SECRET"),
		platform:     os.Getenv("PLATFORM"),
		filepathRoot: os.Getenv("FILEPATH_ROOT"),
		assetsRoot:   os.Getenv("ASSETS_ROOT"),
		s3Bucket:     os.Getenv("S3_BUCKET"),
		s3Region:     os.Getenv("S3_REGION"),
		s3CfDistro:   os.Getenv("S3_CF_DISTRO"),
	}

	mux := http.NewServeMux()

	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir(cfg.filepathRoot)))
	mux.Handle("/app/", appHandler)

	assetsHandler := http.StripPrefix("/assets", http.FileServer(http.Dir(cfg.assetsRoot)))
	mux.Handle("/assets/", assetsHandler)

	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":\"ok\"}"))
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("SoloStream serving on http://localhost:%s/app/\n", port)
	log.Fatal(srv.ListenAndServe())
}
