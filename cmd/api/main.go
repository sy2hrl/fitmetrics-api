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

	// 4. Daftar Jalur API (Routes)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Server berjalan dengan baik!",
		})
	})

	r.POST("/workouts", controllers.CreateWorkout)
	r.GET("/workouts", controllers.GetWorkouts)
	r.PUT("/workouts/:id", controllers.UpdateWorkout)
	r.DELETE("/workouts/:id", controllers.DeleteWorkout) // <--- INI RUTE TERAKHIR KITA
	// 5. Ambil port dari .env, atau gunakan 8080 sebagai default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 6. Jalankan server
	log.Printf("Server siap menerima request di port %s...", port)
	r.Run(":" + port)
}
