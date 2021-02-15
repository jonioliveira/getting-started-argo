package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAge struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
	Age   string `json:"age" binding:"required"`
}

var dbV2 = make(map[string]UserAge)

func AddUserRoutesV2(rg *gin.RouterGroup) {
	userRouter := rg.Group("/user")
	userRouter.GET("/:name", getUserHandlerV2)
	userRouter.POST("/", postUserHandlerV2)
}

func getUserHandlerV2(ctx *gin.Context) {
	user := ctx.Params.ByName("name")
	value, ok := db[user]
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
}

func postUserHandlerV2(ctx *gin.Context) {
	var newUser UserAge
	err := ctx.Bind(&newUser)

	if err == nil {
		dbV2[newUser.Name] = newUser
		ctx.JSON(http.StatusOK, "")
	}
}
