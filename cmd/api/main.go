package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/syahrul/fitmetrics-api/internal/config"
	"github.com/syahrul/fitmetrics-api/internal/controllers" // Import controller baru kita
)

func main() {
	// 1. Muat file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan variabel sistem.")
	}

	// 2. Hubungkan ke Database PostgreSQL
	config.ConnectDatabase()

	// 3. Inisialisasi router Gin
	r := gin.Default()

	// INI JALUR BARU KITA: Menerima data JSON dan meneruskannya ke fungsi CreateWorkout
	r.POST("/workouts", controllers.CreateWorkout)

	// 4. Daftar Jalur API (Routes)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Server berjalan dengan baik!",
		})
	})

	// Rute CRUD Latihan
	r.GET("/workouts", controllers.GetWorkouts) // <--- INI RUTE BARU KITA

	// 5. Ambil port dari .env, atau gunakan 8080 sebagai default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 6. Jalankan server
	log.Printf("Server siap menerima request di port %s...", port)
	r.Run(":" + port)
}
