package controllers

import (
  "database/sql"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "github.com/revel/revel"
  "mikes_app/app/models"
)

type GormController struct {
  *revel.Controller
  Tx *gorm.DB
}

var Db *gorm.DB

func InitDB() {
	var err error
  	Db, err = gorm.Open("sqlite3", "test.db")
  	if err != nil {
    	panic(err)
  	}
  	Db.AutoMigrate(&models.User{})
}

func (c *GormController) Begin() revel.Result {
  	txn := Db.Begin()
  	if txn.Error != nil {
    	panic(txn.Error)
  	}
  	c.Tx = txn
  	return nil
}

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