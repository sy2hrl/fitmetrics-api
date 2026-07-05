package models

import (
	"time"
)

// WorkoutLog merepresentasikan tabel catatan latihan di database
type WorkoutLog struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ActivityType string    `gorm:"type:varchar(50)" json:"activity_type"` // Contoh: "Running", "Bodyweight"
	Duration     int       `json:"duration"`                              // Durasi latihan dalam menit
	Repetitions  int       `json:"repetitions"`                           // Untuk pencatatan repetisi maksimal beban tubuh
	Distance     float64   `json:"distance"`                              // Untuk pencatatan jarak treadmill dalam km
	CreatedAt    time.Time `json:"created_at"`
}
