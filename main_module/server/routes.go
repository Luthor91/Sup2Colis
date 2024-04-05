package server

import (
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
)

// StartWebsite initialise le serveur web avec Fiber
func DefineRoutes() {
	// Initialisation de l'application Fiber
	app := fiber.New()

	app.Static("/", "./public")

	// Gestionnaire pour l'URL racine "/"
	app.Get("/index", func(c *fiber.Ctx) error {
		tmpl := template.Must(template.ParseFiles("../../public/index.html"))
		c.Type("html")
		return tmpl.Execute(c.Response().BodyWriter(), nil)
	})

	/************************
	*	Routes des colis
	*/
	app.Get("/colis/list", getAllColis)

	app.Get("/colis/add", func(c *fiber.Ctx) error {
		tmpl := template.Must(template.ParseFiles("../../public/colis/add_colis.html"))
		c.Type("html")
		return tmpl.Execute(c.Response().BodyWriter(), nil)
	})

	app.Get("/colis/seek", func(c *fiber.Ctx) error {
		tmpl := template.Must(template.ParseFiles("../../public/colis/recherche_colis.html"))
		c.Type("html")
		return tmpl.Execute(c.Response().BodyWriter(), nil)
	})

	app.Get("/colis/affect", func(c *fiber.Ctx) error {
		tmpl := template.Must(template.ParseFiles("../../public/colis/affectation_colis.html"))
		c.Type("html")
		return tmpl.Execute(c.Response().BodyWriter(), nil)
	})
	app.Post("/colis/affect", affectColis)

	app.Post("/colis/add", colisFormHandler)

	app.Get("/colis/:id", getColisById)

	app.Delete("/colis/:id", deleteColisHandler)

	/************************
	*	Routes des Zones
	*/
	app.Get("/zone_colis/list", getAllZone)

	app.Get("/zone_colis/seek", func(c *fiber.Ctx) error {
		tmpl := template.Must(template.ParseFiles("../../public/zone/recherche_zone.html"))
		c.Type("html")
		return tmpl.Execute(c.Response().BodyWriter(), nil)
	})

	app.Get("/zone_colis/add", func(c *fiber.Ctx) error {
		tmpl := template.Must(template.ParseFiles("../../public/zone/add_zone.html"))
		c.Type("html")
		return tmpl.Execute(c.Response().BodyWriter(), nil)
	})
	app.Post("/zone_colis/add", zoneFormHandler)

	app.Get("/zone_colis/:id", getZoneById)

	app.Delete("/zone_colis/:id", deleteZoneHandler)

	log.Fatal(app.Listen(":8080"))
}
