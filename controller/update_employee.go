package controller

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jmoiron/sqlx"
)

func NewUpdateMahasiswaController(db *sqlx.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Query("id")

        // Handle POST (update logic)
        if c.Method() == fiber.MethodPost {
            name := c.FormValue("name")
            address := c.FormValue("address")
            npwp := c.FormValue("npwp")

            _, err := db.Exec("UPDATE mahasiswa SET name=$1, npwp=$2, address=$3 WHERE id=$4", name, npwp, address, id)
            if err != nil {
                return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
            }

            return c.Redirect("/mahasiswa", fiber.StatusSeeOther)
        }

        // Handle GET (fetch data for form)
        var mahasiswa Mahasiswa
        err := db.Get(&mahasiswa, "SELECT id, name, npwp, address FROM mahasiswa WHERE id=$1", id)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
        }

        return c.Render("update", fiber.Map{
            "mahasiswa": mahasiswa,
        })
    }
}


