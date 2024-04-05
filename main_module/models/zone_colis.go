package models

import (
	"errors"

	"gorm.io/gorm"
)

// GetColisByID récupère un colis à partir de la base de données en fonction de son ID
func GetZoneByID(db *gorm.DB, id uint) (Zone, error) {
	var zone Zone
	result := db.First(&zone, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Zone{}, errors.New("zone non trouvé")
		}
		return Zone{}, result.Error
	}
	return zone, nil
}

// UpdateColis met à jour les informations d'un colis dans la base de données en fonction de son ID
func UpdateZone(db *gorm.DB, id uint, updatedZone Zone) error {
	var zone Zone
	result := db.First(&zone, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("colis non trouvé")
		}
		return result.Error
	}
	result = db.Model(&zone).Updates(updatedZone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// AddZone ajoute un nouveau colis à la base de données
func AddZone(db *gorm.DB, zone *Zone) error {
	result := db.Create(zone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteColisByID supprime un colis de la base de données en fonction de son ID
func DeleteZoneByID(db *gorm.DB, id uint) error {
	result := db.Delete(&Zone{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllColis récupère tous les coliss de la base de données
func GetAllZone(db *gorm.DB) ([]Zone, error) {
	var all_zones []Zone
	result := db.Find(&all_zones)
	if result.Error != nil {
		return nil, result.Error
	}
	return all_zones, nil
}
