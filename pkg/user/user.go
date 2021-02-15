package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

type User struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func AddUserRoutesV1(rg *gin.RouterGroup) {
	userRouter := rg.Group("/user")
	userRouter.GET("/:name", getUserHandler)
	userRouter.POST("/", postUserHandler)
}

func getUserHandler(ctx *gin.Context) {
	user := ctx.Params.ByName("name")
	value, ok := db[user]
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
}

func postUserHandler(ctx *gin.Context) {
	var newUser User
	err := ctx.Bind(&newUser)

	if err == nil {
		db[newUser.Name] = newUser.Value
		ctx.JSON(http.StatusOK, "")
	}
}
