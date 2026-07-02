package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBGlobal adalah variabel global agar database bisa dipakai di file lain
var DBGlobal *gorm.DB

func ConnectDatabase() {
	// Ambil konfigurasi dari .env
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Susun koneksi PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, name, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menyambung ke database PostgreSQL: ", err)
	}

	DBGlobal = db
	log.Println("Koneksi ke PostgreSQL berhasil!")
}