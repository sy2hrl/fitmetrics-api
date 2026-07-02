package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Membaca konfigurasi dari file .env (keluar 2 tingkat folder untuk mencari .env)
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Peringatan: Gagal membaca file .env, menggunakan pengaturan bawaan.")
	}

	// Mengambil port dari file .env, jika kosong default ke 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Inisiasi framework Gin
	router := gin.Default()

	// Membuat lintasan/endpoint pertamamu
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "sukses",
			"message": "FitMetrics API berhasil berjalan di komputer Syahrul!",
		})
	})

	// Menyalakan server
	log.Printf("Server siap menerima request di port %s...", port)
	router.Run(":" + port)
}
