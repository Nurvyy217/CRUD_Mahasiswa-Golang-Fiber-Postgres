package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	mahasiswaRoute := router.Group("mahasiswa")
	{
		
		mahasiswaRoute.Get("", NewIndexMahasiswa(db))
		mahasiswaRoute.Get("/create", NewCreateMahasiswaController(db))
		mahasiswaRoute.Post("/create", NewCreateMahasiswaController(db)) // <- ini juga perlu biar create jalan
		mahasiswaRoute.Get("/update", NewUpdateMahasiswaController(db))
		mahasiswaRoute.Post("/update", NewUpdateMahasiswaController(db)) // <- fix edit di sini
		mahasiswaRoute.Get("/delete", NewDeleteMahasiswaController(db))
		
	}
}
