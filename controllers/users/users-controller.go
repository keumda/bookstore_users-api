package users

import (
	"github.com/gin-gonic/gin"
	"github.com/keumda/domain/users"
	"github.com/keumda/services"
	"github.com/keumda/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid param")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusFound, result)
}

func CreateUser(c *gin.Context) {
	var user users.User
	/*  Replaced by ShouldBindJson
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil{
		fmt.Println(err.Error())
		return
	}*/
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, nil)
}
