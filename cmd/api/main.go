package main

import (
	"log"

	"github.com/Nurvyy217/crud-mahasiswa-go/controller"
	"github.com/Nurvyy217/crud-mahasiswa-go/external/database"
	"github.com/Nurvyy217/crud-mahasiswa-go/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	if db != nil {
		log.Println("db connected")
	}

	engine := html.New("./views", ".html")
	router := fiber.New(fiber.Config{
		AppName: config.Cfg.App.Name,
		Views: engine,
	})

	// route utama
	// router.Get("/mahasiswa", controller.NewIndexMahasiswa(db))
	// router.Get("/mahasiswa/create", controller.NewCreateMahasiswaController(db))
	// router.Post("/mahasiswa/create", controller.NewCreateMahasiswaController(db))
	// router.Get("/mahasiswa/update", controller.NewUpdateMahasiswaController(db))
	// router.Post("/mahasiswa/update", controller.NewUpdateMahasiswaController(db))
	// router.Get("/mahasiswa/delete", controller.NewDeleteMahasiswaController(db))

	controller.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
