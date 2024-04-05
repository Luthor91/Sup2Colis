package models

import (
	"gorm.io/gorm"
)


type Colis struct {
	gorm.Model
	Id          int     `gorm:"column:id"`
	Identifiant string  `gorm:"column:identifiant"`
	Poids       float64 `gorm:"column:poids"`
	Adresse     string  `gorm:"column:adresse"`
	Destination string  `gorm:"column:destination"`
	FkZone      int     `gorm:"column:fk_zone"`  // Clé étrangère vers la zone
	Zone        Zone    `gorm:"foreignKey:FkZone"` // Relation belongsTo
}

type ColisData struct {
	Liste_colis []Colis
}


type Zone struct {
	gorm.Model
	Id             int     `gorm:"column:id"`
	Name           string  `gorm:"column:name;default:Name_Zone_Colis"`
	Adresse        string  `gorm:"column:adresse"`
	NombreMaxColis int     `gorm:"column:nombre_max_colis"`
	Colis          []Colis `gorm:"foreignKey:FkZone"` // Relation hasMany vers les colis
}

type ZoneData struct {
	Liste_zone []Zone
}