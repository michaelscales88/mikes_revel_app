package models

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "github.com/coopernurse/gorp"
   "log"
)

// Global database references
var db *sql.DB
var dbmap *gorp.DbMap

// Database settings
var db_name = "mydb"
var db_user = "root"
var db_pw = "mypw"

// Create database connection
func Init_DB() {
   var err error

   db, err = sql.Open("mysql", db_user + ":" + db_pw + "@tcp(127.0.0.1:3306)/" + db_name)
   dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

   if err != nil {
      log.Println("Failed to connect to database: ")
      log.Panic(err)
   } else {
      err = db.Ping()

      if err != nil {
         log.Println("Failed to ping database: ")
         log.Panic(err)
      } else {
         log.Println("Database connected.")
      }
   }

   _ = dbmap.AddTableWithName(User{}, "user").SetKeys(false, "Id")
   dbmap.CreateTablesIfNotExists()
}
