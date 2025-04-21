package controller

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jmoiron/sqlx"
)

func NewDeleteMahasiswaController(db *sqlx.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Query("id")

        _, err := db.Exec("DELETE FROM mahasiswa WHERE id = $1", id)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
        }

        return c.Redirect("/mahasiswa", fiber.StatusSeeOther)
    }
}

