package models

import (
	"errors"

	"gorm.io/gorm"
)

// GetColisByID récupère un colis à partir de la base de données en fonction de son ID
func GetColisByID(db *gorm.DB, id uint) (Colis, error) {
	var colis Colis
	result := db.First(&colis, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Colis{}, errors.New("colis non trouvé")
		}
		return Colis{}, result.Error
	}
	return colis, nil
}

// UpdateColis met à jour les informations d'un colis dans la base de données en fonction de son ID
func UpdateColis(db *gorm.DB, id uint, updatedColis Colis) error {
	var colis Colis
	result := db.First(&colis, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("colis non trouvé")
		}
		return result.Error
	}
	result = db.Model(&colis).Updates(updatedColis)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// AddColis ajoute un nouveau colis à la base de données
func AddColis(db *gorm.DB, colis *Colis) error {
	result := db.Create(colis)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteColisByID supprime un colis de la base de données en fonction de son ID
func DeleteColisByID(db *gorm.DB, id uint) error {
	result := db.Delete(&Colis{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllColis récupère tous les coliss de la base de données
func GetAllColis(db *gorm.DB) ([]Colis, error) {
	var all_colis []Colis
	result := db.Find(&all_colis)
	if result.Error != nil {
		return nil, result.Error
	}
	return all_colis, nil
}