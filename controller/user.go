package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (ctrl *userController) GetAllUsers(db *gorm.DB) echo.HandlerFunc {
	return func (c echo.Context) error {
	
		users, err := ctrl.userService.FindAll()
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(200, echo.Map{
			"data": users,
		})
	}
}

func (ctrl *userController) CreateUser(db *gorm.DB) echo.HandlerFunc {
	return func (c echo.Context) error {
		var user model.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		user, err := ctrl.userService.CreateUser(user)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}

// func CreateUser(db *gorm.DB) echo.HandlerFunc {
// 	user := model.User{}
// 	return func(c echo.Context) error {
// 		if err := c.Bind(&user); err != nil {
// 			return c.JSON(400, echo.Map{
// 				"error": err.Error(),
// 			})
// 		}
// 		err := db.Create(&user).Error
// 		if err != nil {
// 			return c.JSON(500, echo.Map{
// 				"error": err.Error(),
// 			})
// 		}
// 		return c.JSON(200, echo.Map{
// 			"data": user,
// 		})
// 	}
// }
