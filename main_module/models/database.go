package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectToDatabase initialise la connexion à la base de données
func ConnectToDatabase() *gorm.DB {
	// Connexion à la base de données PostgreSQL
	dsn := "user=postgres password=postgres dbname=Sup2Colis sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erreur de connexion à la base de données: %v", err)
	}

	// Migration: crée la table "ZoneColis" si elle n'existe pas déjà
	err = db.AutoMigrate(&Zone{})
	if err != nil {
		log.Fatalf("Erreur de migration de la base de données: %v", err)
	}

	// Migration: crée la table "colis" si elle n'existe pas déjà
	err = db.AutoMigrate(&Colis{})
	if err != nil {
		log.Fatalf("Erreur de migration de la base de données: %v", err)
	}
	var zone Zone = Zone{Adresse: "Inconnu", Name: "Aucune zone défini", NombreMaxColis: 20}

	AddZone(db, &zone)

	return db
}