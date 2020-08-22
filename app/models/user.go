package models

import (
   "log"
)

type User struct {
   Id int
   First_name string
   Last_name string
}

func FindUser(number int) (bool, *User) {
   obj, err := dbmap.Get(User{}, number)
   usr := obj.(*User)

   if err != nil {
      log.Print("ERROR FindUser: ")
      log.Println(err)
   }

   return (err == nil), usr
}
