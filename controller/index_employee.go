package controller

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jmoiron/sqlx"
)

type Mahasiswa struct {
    ID      int    `db:"id"`
    Name    string `db:"name"`
    NPWP    string `db:"npwp"`
    Address string `db:"address"`
}

func NewIndexMahasiswa(db *sqlx.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var mahasiswas []Mahasiswa
        err := db.Select(&mahasiswas, "SELECT id, name, npwp, address FROM mahasiswa")
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
        }

        return c.Render("index", fiber.Map{
            "mahasiswas": mahasiswas,
        })
    }
}