package main

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)
type User struct{
  FirstName string `json:"firstName,omitempty"`
  LastName  string `json:"lastName,omitempty"`
  UserName  string `json:"userName,omitempty"`
  Password  string `json:"passName,omitempty"`
}

func index(c echo.Context) error{
    return c.String(http.StatusOK, "Hello World")
}
func getUsers(c echo.Context) error{
  beer := User{
    FirstName:"thanikarn",
    LastName:"yana",
    UserName:"beer",

  }
    return c.JSON(http.StatusOK, beer)
}
    func  getUsersByID(c echo.Context) error  {
      id := c.Param("id")
     return c.JSON(http.StatusOK,id)
    }
func saveUser(c echo.Context) error{
  user := new(User)
  err := c.Bind(user)
  if  err != nil{
    return c.JSON(http.StatusBadRequest,nil)
  }
  return c.JSON(http.StatusOK,user)
}
func main() {
  e := echo.New()
  e.Use(middleware.Logger())
  e.GET("/", index)
  e.GET("/users",getUsers)
  e.GET("/users/:id",getUsersByID)
  e.POST("/users",saveUser)
  e.Logger.Fatal(e.Start(":1323"))
}
