package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var db *gorm.DB
var err error
type User struct {
    ID      int
    Name    string
    Age     int
    AddTime int
}
func (User) TableName() string {
	return "user"
}
func getUsers(c *gin.Context)  {
    users := make([]User, 0)
    //db.Find(&users)
    db.Find(&users)
    c.JSON(http.StatusOK, gin.H{
        "data": users,
    })
}

func getOneUser(c *gin.Context)  {
    id := c.Param("id")
    var user User
    if err := db.Where("id = ?", id).First(&user).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        fmt.Println(err)
    } else {
        c.JSON(http.StatusOK, gin.H{
            "data": user,
        })
    }
}

func createUser(c *gin.Context)  {
    var user User
    c.BindJSON(&user)
    user.Name = "sss"
    user.Age=22

    //user.AddTime = time.Now()
    db.Create(&user)
    c.JSON(http.StatusOK, gin.H{
        "data": user,
    })
}

func updateUser(c *gin.Context)  {
    id := c.Param("id")
    var user User

    if err := db.Where("id = ?", id).First(&user).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        fmt.Println(err)
    } else {
        c.BindJSON(&user)
        db.Save(&user)
        c.JSON(http.StatusOK, gin.H{
            "data": user,
        })
    }
}

func deleteUser(c *gin.Context)  {
    id := c.Param("id")
    var user User
    db.Where("id = ?", id).Delete(&user)
    c.JSON(http.StatusOK, gin.H{
        "data": "this has been deleted!",
    })
}















func main() {
	db, err = gorm.Open("mysql", "root:root@(192.168.1.145:3306)/bee?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
	r.GET("/user/", getUsers)
	r.GET("/user/:id", getOneUser)
	r.POST("/user", createUser)
	r.PUT("/user/:id", updateUser)
	r.DELETE("/user/:id", deleteUser)
    r.Run(":10080") // listen and serve on 0.0.0.0:10080


}


