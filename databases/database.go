package databases

import (
	"bioskop-management-gin/configs"
	"log"
	"os"
)

const migrationVersion = "001_create_bioskop"

func RunMigration() {
	db := configs.DB

	// pastikan tabel schema_migrations ada
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(50) PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("failed to create schema_migrations:", err)
	}

	// cek apakah migration sudah pernah dijalankan
	var exists bool
	err = db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM schema_migrations WHERE version = $1
		)
	`, migrationVersion).Scan(&exists)

	if err != nil {
		log.Fatal("failed to check migration version:", err)
	}

	if exists {
		log.Println("Migration already applied:", migrationVersion)
		return
	}

	// baca file migration
	sqlBytes, err := os.ReadFile("migrations/001_create_bioskop.sql")
	if err != nil {
		log.Fatal("failed to read migration file:", err)
	}

	// jalankan migration
	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal("failed to run migration:", err)
	}

	// simpan version
	_, err = db.Exec(`
		INSERT INTO schema_migrations (version)
		VALUES ($1)
	`, migrationVersion)

	if err != nil {
		log.Fatal("failed to save migration version:", err)
	}

	log.Println("Migration applied:", migrationVersion)
}
