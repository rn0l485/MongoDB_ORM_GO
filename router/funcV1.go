package router

import (
	//"fmt"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"clientSystem/db"
	"clientSystem/db/models"
)


func help(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"err":"NoneHelp",
	})
}


func add(c *gin.Context){
	var payload models.AddUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	if ok := verify(payload.Ticket); !ok {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Ticket verification error.",
			},
		)
		return
	}

	mgoC, err := mgo.GetClient()
	defer mgo.ReturnClient(mgoC)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return		
	}

	userCollection := mgoC.Database("BASE").Collection("user")
	_, err = userCollection.InsertOne(context.TODO(), payload.Payload.NewUser)
	if err != nil {
	    c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	} else {
		c.JSON( 
			http.StatusOK, 
			gin.H{
				"error": "None",
			},
		)
	}
}
func search(c *gin.Context){
	var payload models.SearchUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	if ok := verify(payload.Ticket); !ok {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Ticket verification error.",
			},
		)
		return
	}
	
}
func modify(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"err":"testing",
	})
}
func delete(c *gin.Context){
	var payload models.DeleteUser
	var userTarget models.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	if ok := verify(payload.Ticket); !ok {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Ticket verification error.",
			},
		)
		return
	}

	mgoC, err := mgo.GetClient()
	defer mgo.ReturnClient(mgoC)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return		
	}
	userCollection := mgoC.Database("BASE").Collection("user")

	err = userCollection.FindOne(
		context.TODO(), 
		bson.M{
			"Account": payload.Payload.Account,
		},
	).Decode(&userTarget)

	if userTarget.Password != payload.Payload.Password{
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Password error.",
			},
		)
		return		
	}

	_, err = userCollection.DeleteOne(
		context.TODO(), 
		bson.M{
			"Account": payload.Payload.Account,
		},
	)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return			
	} else {
		c.JSON( http.StatusOK, gin.H{
				"err": "None",
		})
	}
}
func verify(ticket string) bool{
	return true
}







