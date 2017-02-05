package main

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  db "./helper/db"
  "./models"
)

func index(c echo.Context) error{
    return c.XML(http.StatusOK, "Hello World")
}
func getUsers(c echo.Context) error{
  user := new(models.User)
  result, _ := user.ReadFromDB()

    return c.JSON(http.StatusOK, result)
}
    func  getUsersByID(c echo.Context) error  {
      user := new(models.User)
      id := c.Param("id")
      user.Id = bson.ObjectIdHex(id)
      result, _ := user.ReadFromDBByID()
     return c.JSON(http.StatusOK,result)
    }

    func  deleteUserByID(c echo.Context) error  {
      user := new(models.User)
      id := c.Param("id")
      user.Id = bson.ObjectIdHex(id)
      user.DeleteFromDBByID()
     return c.NoContent(http.StatusOK)
    }


    func saveUser(c echo.Context) error{
  user := new(models.User)
  err := c.Bind(user)
  if  err != nil{
    return c.JSON(http.StatusBadRequest,nil)
  }
  user.SaveToDB()
  return c.NoContent(http.StatusCreated)
}

    func init()  {
    mongoSession, err := mgo.Dial("localhost:27017")
    if err != nil {
      panic(err)
      }
  mongoSession.SetMode(mgo.Monotonic,true)
  db.MongoSession = mongoSession
  db.UsersCollection = db.MongoSession.DB("maejo").C("users")
}

func main() {

  defer db.MongoSession.Close()
e := echo.New()

e.Use(middleware.Logger())
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"*"},
  AllowMethods: []string{
    echo.GET,
    echo.POST,
  },
  }))



  e.GET("/", index)
  e.GET("/users",getUsers)
  e.GET("/users/:id",getUsersByID)
  e.POST("/users",saveUser)
  e.DELETE("/users/:id",deleteUserByID)
  e.Logger.Fatal(e.Start(":1323"))
  
}
