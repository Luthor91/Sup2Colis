package server

import (
	"html/template"
	"main_module/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// getAllcolis est le gestionnaire HTTP qui affiche les produits
func getAllZone(c *fiber.Ctx) error {
	// Récupérer les produits depuis la base de données
	db := models.ConnectToDatabase()
	all_zones, err := models.GetAllZone(db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération des produits")
	}

	// Préparer les données à passer au template HTML
	data := models.ZoneData{
		Liste_zone: all_zones,
	}

	// Charger le template HTML
	mapage := template.Must(template.ParseFiles("../../public/zone/liste_zone.html"))

	// Passer les données au template et générer la réponse HTTP
	err = mapage.Execute(c.Response().BodyWriter(), data)
	if err != nil {
		return err
	}
	c.Type("html")
	return nil
}

// getcolisById est le gestionnaire HTTP qui affiche un produit spécifique en fonction de son ID
func getZoneById(c *fiber.Ctx) error {
	// Récupérer l'ID du produit depuis l'URL
	zoneID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID de produit invalide")
	}

	// Récupérer le produit depuis la base de données en fonction de son ID
	db := models.ConnectToDatabase()
	zone, err := models.GetZoneByID(db, uint(zoneID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Zone non trouvé")
	}

	// Passer les données du produit au template HTML
	mapage := template.Must(template.ParseFiles("../../public/zone/zone_colis.html"))

	// Passer les données au template et générer la réponse HTTP
	err = mapage.Execute(c.Response().BodyWriter(), zone)
	if err != nil {
		return err
	}
	c.Type("html")
	return nil
}

// colisFormHandler est le gestionnaire HTTP pour soumettre le formulaire d'ajout de produit
func zoneFormHandler(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	adresse := c.FormValue("adresse")
	name := c.FormValue("name")
	nombreMaxColis_str := c.FormValue("nombreMaxColis")

	// Convertir le prix en float64
	nombreMaxColis, err := strconv.ParseInt(nombreMaxColis_str, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Le poids doit être un nombre valide")
	}

	// Créer un nouveau produit avec les données du formulaire
	newzone := models.Zone{Adresse: adresse, Name: name, NombreMaxColis: int(nombreMaxColis)}

	// Ajouter le nouveau produit à la base de données
	db := models.ConnectToDatabase()
	if err := models.AddZone(db, &newzone); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'ajout du colis")
	}

	// Rediriger l'utilisateur vers la page des produits après l'ajout réussi
	return c.Redirect("/zone_colis/list")
}

// deletecolisHandler est le gestionnaire HTTP pour supprimer un produit par son ID
func deleteZoneHandler(c *fiber.Ctx) error {
	// Récupérer l'ID du produit depuis l'URL
	zoneID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID de produit invalide")
	}

	// Supprimer le produit depuis la base de données en fonction de son ID
	db := models.ConnectToDatabase()
	err = models.DeleteZoneByID(db, uint(zoneID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la suppression du produit")
	}

	// Rediriger l'utilisateur vers la liste des produits après la suppression réussie
	return c.Redirect("/zone_colis/list")
}
