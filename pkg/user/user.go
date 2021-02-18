package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jonioliveira/getting-started-argo/pkg/logger"
	"github.com/schollz/progressbar/v3"

	err "github.com/jonioliveira/getting-started-argo/pkg/error"

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
	// fake database access
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		_ = bar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}

	if ok {
		ctx.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		return
	}

	e := err.NewErrDetails(err.ErrItemNotFound, fmt.Sprintf("Could not find user with name: %s", user))
	logger.Error(e.Error())
	ctx.JSON(http.StatusNotFound, gin.H{"message": e.Error()})
}

func postUserHandler(ctx *gin.Context) {
	var newUser User
	err := ctx.Bind(&newUser)

	if err == nil {
		db[newUser.Name] = newUser.Value
		ctx.JSON(http.StatusOK, "")
	}
}
