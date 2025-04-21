package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)


func NewCreateMahasiswaController(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodPost {
			mahasiswa := Mahasiswa{
				Name:    c.FormValue("name"),
				NPWP:    c.FormValue("npwp"),
				Address: c.FormValue("address"),
			}

			_, err := db.NamedExec(`INSERT INTO mahasiswa (name, npwp, address) VALUES (:name, :npwp, :address)`, &mahasiswa)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			//c.Redirect("/mahasiswa", fiber.StatusSeeOther) adalah praktik umum untuk mengarahkan pengguna ke halaman lain setelah operasi seperti pembuatan atau pembaruan data, memastikan pengalaman pengguna yang mulus dan mencegah pengiriman ulang data
			return c.Redirect("/mahasiswa", fiber.StatusSeeOther)
		}
		//c.Render("create", fiber.Map{}) digunakan untuk merender template HTML bernama create.html dan mengirimkannya sebagai respons ke klien.
		return c.Render("create", fiber.Map{})
	}
}

