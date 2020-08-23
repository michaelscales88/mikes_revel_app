package controllers

import (
	"mikes_app/app/models"

	"github.com/revel/revel"
)

// App represents a controller with a gorm orm txn.
type App struct {
	GormController
}

// Index represents the / page.
func (c App) Index() revel.Result {
	var user = models.User{Name: "Admin", Email: "admin@password.com"}
	user.SetNewPassword("demo")
	user.Active = true
	c.Tx.Create(user)
	return c.Render(user)
}
