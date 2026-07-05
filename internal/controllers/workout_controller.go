package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syahrul/fitmetrics-api/internal/config"
	"github.com/syahrul/fitmetrics-api/internal/models"
)

// CreateWorkout adalah fungsi untuk menerima JSON dan menyimpannya ke database
func CreateWorkout(c *gin.Context) {
	// 1. Siapkan wadah kosong berdasarkan struktur tabel kita
	var workout models.WorkoutLog

	// 2. Tangkap data dari luar (JSON) dan masukkan ke wadah kosong tadi
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah!"})
		return
	}

	// 3. Masukkan data yang sudah terisi ke dalam kulkas (database PostgreSQL)
	if err := config.DBGlobal.Create(&workout).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan ke database"})
		return
	}

	// 4. Jika sukses, kembalikan struk (respon) ke pengguna
	c.JSON(http.StatusCreated, gin.H{
		"message": "Mantap! Data latihan berhasil dicatat.",
		"data":    workout,
	})
}

	// GetWorkouts adalah fungsi untuk mengambil semua data latihan dari database
func GetWorkouts(c *gin.Context) {
	// Siapkan wadah untuk menampung banyak data (menggunakan array/slice)
	var workouts []models.WorkoutLog

	// Suruh database mencari semua data dan masukkan ke wadah tadi
	if err := config.DBGlobal.Find(&workouts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari database"})
		return
	}

	// Kembalikan datanya ke pengguna dalam bentuk JSON
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data latihan",
		"data":    workouts,
	})
}

