package models

import (
  "gopkg.in/mgo.v2/bson"
  db "../helper/db"
)

type User struct{
  Id bson.ObjectId `json:"id,omitempty"bson:"_id,omitempty"`
  FirstName string `json:"firstName,omitempty"bson:"firstName,omitempty"`
  LastName  string `json:"lastName,omitempty"bson:"lastName,omitempty"`
  UserName  string `json:"userName,omitempty"bson:"UserName,omitempty"`
  Password  string `json:"passName,omitempty"bson:"passName,omitempty"`
  Imge string `json:"imge,omitempty"bson:"imge,omitempty"`
  Deteil string `json:"detil,omitempty"bson:"detil,omitempty"`
}

func (u *User)SaveToDB() error  {
  err := db.UsersCollection.Insert(&u)
  if err != nil{
    return err
  }
  return nil
}

func (u *User)ReadFromDB() ([]User,error)  {
  result := []User{}
  err := db.UsersCollection.Find(nil).All(&result)
  if err != nil {
    return nil,err
  }
  return result,nil
}

func (u *User)ReadFromDBByID()(*User, error){
  err := db.UsersCollection.Find(bson.M{"_id": u.Id}).One(&u)
  if err != nil{
    return nil, err
  }
  return u, nil
  }
  func (u *User)DeleteFromDBByID()(*User, error){
    err := db.UsersCollection.RemoveId( u.Id)
    if err != nil{
      return nil, err
    }
    return u, nil
    }
