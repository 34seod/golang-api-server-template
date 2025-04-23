package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"golang-api-server-template/internal/dto"
	"golang-api-server-template/internal/model"
	"golang-api-server-template/internal/service"
	"golang-api-server-template/tools"
)

// @Summary		Find user by ID
// @Description	Get a single user by ID
// @Tags			users
// @Param			id	path	int	true	"User ID"
// @Produce		json
// @Success		200	{object}	model.User
// @Router			/v1/users/{id} [get]
func UserFindByID(c *gin.Context) {
	status := http.StatusOK
	var errMsg any = nil

	var userDto dto.UserIdFromUri
	if err := c.ShouldBindUri(&userDto); err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}

	user, err := service.UserFindByID(&userDto)
	if err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}

	response(c, status, user, errMsg)
}

// @Summary		Get all users
// @Description	List all users
// @Tags			users
// @Produce		json
// @Success		200	{array}	model.User
// @Router			/v1/users [get]
func UserFindAll(c *gin.Context) {
	status := http.StatusOK
	var errMsg any = nil

	users, err := service.UserFindAll()
	if err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}

	response(c, status, users, errMsg)
}

// @Summary		Create user
// @Description	Create a new user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user	body	dto.UserBodyFromCreateRequest	true	"User data"
// @Success		204
// @Router			/v1/users [post]
func UserCreate(c *gin.Context) {
	status := http.StatusNoContent
	var errMsg any = nil

	var userDto dto.UserBodyFromCreateRequest
	if err := c.ShouldBind(&userDto); err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}
	fmt.Println("userDto.Name")
	fmt.Println(userDto.Name)

	user := model.User{}
	copier.Copy(&user, &userDto)
	if err := user.Validate(); err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err, validateErrors(err))
	}

	if err := service.UserCreate(&user); err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}

	response(c, status, nil, errMsg)
}

// @Summary		Update user
// @Description	Update a user's info
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path	int								true	"User ID"
// @Param			user	body	dto.UserBodyFromUpdateRequest	true	"User data"
// @Success		204
// @Router			/v1/users/{id} [put]
func UserUpdate(c *gin.Context) {
	status := http.StatusNoContent
	var errMsg any = nil

	var userDto dto.UserBodyFromUpdateRequest
	if err := c.ShouldBindUri(&userDto); err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}
	if err := c.ShouldBind(&userDto); err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}
	user := model.User{}
	copier.Copy(&user, &userDto)
	if err := user.Validate(); err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err, validateErrors(err))
	}

	err := service.UserUpdate(&userDto)
	if err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}

	response(c, status, nil, errMsg)
}

// @Summary		Delete user
// @Description	Delete a user by ID
// @Tags			users
// @Produce		json
// @Param			id	path	int	true	"User ID"
// @Success		204
// @Router			/v1/users/{id} [delete]
func UserDelete(c *gin.Context) {
	status := http.StatusNoContent
	var errMsg any = nil

	var userDto dto.UserIdFromUri
	if err := c.ShouldBindUri(&userDto); err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}

	err := service.UserDelete(&userDto)
	if err != nil {
		tools.PrintTrace()
		status, errMsg = errorCtrl(http.StatusBadRequest, err)
	}

	response(c, status, nil, errMsg)
}
