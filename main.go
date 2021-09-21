package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User

func main() {
	r := gin.Default()

	//ここでルーティング
	userRoutes := r.Group("/users", GetUser)
	{
		userRoutes.GET("/", GetUser)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", EditUser)
		userRoutes.DELETE("/:id", DeleteUser)
	}
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())

	}
}

func GetUser(c *gin.Context) {
	c.JSON(200, Users)
}

func CreateUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	reqBody.ID = uuid.New().String()

	Users = append(Users, reqBody)
	c.JSON(200, gin.H{
		"success": true,
	})

}

func EditUser(c *gin.Context) {
	id := c.Param("id")

	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age
			c.JSON(200, gin.H{
				"message": "success",
			})
			//returnでここで終了する
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			//	全ての要素が取得される
			// 消したユーザーを飛ばしている
			// t :=[1,422,33,456,31]
			// t[:2] =[1,422]
			// t[2+1:] =[456,31]
			// これらを組み合わせると
			//[1,422,456,31]

			Users = append(Users[:i], Users[i+1:]...)
			c.JSON(200, gin.H{
				"message": "success",
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}
