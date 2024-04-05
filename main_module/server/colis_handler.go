package server

import (
	"html/template"
	"main_module/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// getAllcolis est le gestionnaire HTTP qui affiche les produits
func getAllColis(c *fiber.Ctx) error {
	// Récupérer les produits depuis la base de données
	db := models.ConnectToDatabase()
	all_colis, err := models.GetAllColis(db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération des produits")
	}

	// Préparer les données à passer au template HTML
	data := models.ColisData{
		Liste_colis: all_colis,
	}

	// Charger le template HTML
	mapage := template.Must(template.ParseFiles("../../public/colis/liste_colis.html"))

	// Passer les données au template et générer la réponse HTTP
	err = mapage.Execute(c.Response().BodyWriter(), data)
	if err != nil {
		return err
	}
	c.Type("html")
	return nil
}

// getcolisById est le gestionnaire HTTP qui affiche un produit spécifique en fonction de son ID
func getColisById(c *fiber.Ctx) error {
	// Récupérer l'ID du produit depuis l'URL
	colisID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID de produit invalide")
	}

	// Récupérer le produit depuis la base de données en fonction de son ID
	db := models.ConnectToDatabase()
	colis, err := models.GetColisByID(db, uint(colisID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Colis non trouvé")
	}

	// Passer les données du produit au template HTML
	mapage := template.Must(template.ParseFiles("../../public/colis/colis.html"))

	// Passer les données au template et générer la réponse HTTP
	err = mapage.Execute(c.Response().BodyWriter(), colis)
	if err != nil {
		return err
	}
	c.Type("html")
	return nil
}

// colisFormHandler est le gestionnaire HTTP pour soumettre le formulaire d'ajout de produit
func colisFormHandler(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	identifiant := c.FormValue("identifiant")
	poids_str := c.FormValue("poids")
	adresse := c.FormValue("adresse")
	destination := c.FormValue("destination")

	// Convertir le prix en float64
	poids, err := strconv.ParseFloat(poids_str, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Le poids doit être un nombre valide")
	}

	// Créer un nouveau produit avec les données du formulaire
	newcolis := models.Colis{Identifiant: identifiant, Poids: poids, Adresse: adresse, Destination: destination, FkZone: 1}

	// Ajouter le nouveau produit à la base de données
	db := models.ConnectToDatabase()
	if err := models.AddColis(db, &newcolis); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'ajout du colis")
	}

	// Rediriger l'utilisateur vers la page des produits après l'ajout réussi
	return c.Redirect("/colis/list")
}

func affectColis(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	colisIDStr := c.FormValue("colis_id")
	zoneIDStr := c.FormValue("zone_id")

	// Convertir les ID en int64
	colisID, err := strconv.ParseInt(colisIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("L'ID du colis doit être un nombre valide")
	}
	zoneID, err := strconv.ParseInt(zoneIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("L'ID de la zone doit être un nombre valide")
	}

	// Se connecter à la base de données
	db := models.ConnectToDatabase()

	// Récupérer la zone à partir de la base de données
	var zone models.Zone
	if err := db.Where("id = ?", zoneID).First(&zone).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération de la zone")
	}
	var colis models.Colis
	if err := db.Where("id = ?", colisID).First(&colis).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération du colis")
	}

	// Vérifier si le nombre de colis affectés à la zone dépasse le nombre maximal de colis autorisés
	if zone.NombreMaxColis <= len(zone.Colis) {
		return c.Status(fiber.StatusBadRequest).SendString("Le nombre maximal de colis pour cette zone est atteint")
	}

	// Mettre à jour FkColis de la zone avec l'ID du colis
	colis.FkZone = int(zoneID)
	colis.Adresse = zone.Adresse

	if err := db.Save(&colis).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'affectation du colis à la zone")
	}

	// Rediriger l'utilisateur vers la page listant les colis après l'affectation réussie
	return c.Redirect("/colis/list")
}

// deletecolisHandler est le gestionnaire HTTP pour supprimer un produit par son ID
func deleteColisHandler(c *fiber.Ctx) error {
	// Récupérer l'ID du produit depuis l'URL
	colisID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID de produit invalide")
	}

	// Supprimer le produit depuis la base de données en fonction de son ID
	db := models.ConnectToDatabase()
	err = models.DeleteColisByID(db, uint(colisID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la suppression du produit")
	}

	// Rediriger l'utilisateur vers la liste des produits après la suppression réussie
	return c.Redirect("/colis/list")
}
