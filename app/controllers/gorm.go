package controllers

import (
	"database/sql"
	"mikes_app/app/models"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

// GormController represents a controller with GORM included.
type GormController struct {
	*revel.Controller
	Tx *gorm.DB
}

// Db represents a gorm.DB which can be accessed in a txn.
var Db *gorm.DB

// InitDB opens and manages the database.
func InitDB() {
	var err error
	Db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&models.User{})
}

// Begin injects the Db (*gorm.DB) into the txn.
func (c *GormController) Begin() revel.Result {
	txn := Db.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	c.Tx = txn
	return nil
}

// Commit and clear the txn.
func (c *GormController) Commit() revel.Result {
	if c.Tx == nil {
		return nil
	}
	c.Tx.Commit()
	if err := c.Tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Tx = nil
	return nil
}

// Rollback the DB session for a bad txn, and clear then clear the txn.
func (c *GormController) Rollback() revel.Result {
	if c.Tx == nil {
		return nil
	}
	c.Tx.Rollback()
	if err := c.Tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Tx = nil
	return nil
}
