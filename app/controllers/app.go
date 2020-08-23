package controllers

import (
	"github.com/revel/revel"
	"mikes_app/app/models"
)

type App struct {
	GormController
}

func (c App) Index() revel.Result {
	user := &models.User{Name: "Admin"}
	c.Tx.NewRecord(user)
	c.Tx.Create(user)
	return c.Render(user)
}
