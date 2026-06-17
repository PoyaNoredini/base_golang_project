package v01

import (
	"BaseProject/api/controllers"
	appErrors "BaseProject/api/error"
	"BaseProject/api/helper"
	"BaseProject/config"
	"BaseProject/models"
	"BaseProject/api/validations"
	"github.com/gin-gonic/gin"
	"BaseProject/api/resource"
	"strconv"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Index(ctx *gin.Context) {
	c.Handle(ctx, func() (interface{}, *appErrors.AppError) {
		var users []models.User
		if err := config.DB.Preload("UserRoles.Role").Find(&users).Error; err != nil {
			return nil, c.HandleDBError(err, "Users not found")
		}
		return resource.NewUserResourceCollection(users), nil
	})
}
	func (c *UserController) Create(ctx *gin.Context) {
		c.Handle(ctx, func() (interface{}, *appErrors.AppError) {

			var req validations.RegisterRequest
			if err := ctx.ShouldBindJSON(&req); err != nil {
				return nil, appErrors.BadRequest(err.Error())
			}
					
			password, err := helper.HashPassword(req.Password)
			if err != nil {
				return nil, appErrors.Internal("Failed to hash password")
			}
	
			user := models.User{
				First_name:   req.FirstName,
				Last_name:    req.LastName,
				Phone_number: req.Phone,
				National_id:  req.NationalID,
				Email:        req.Email,
				Password:     password,
			}
	
			if err := config.DB.Create(&user).Error; err != nil {
				return nil, appErrors.Internal("Failed to create user")
			}
	
			return user, nil
		})
	}


	func (c *UserController) Profile(ctx *gin.Context) {
		c.Handle(ctx, func() (interface{}, *appErrors.AppError) {
			userID := ctx.GetUint("user_id")
			var user models.User
			if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
				return nil, appErrors.Internal("Failed to get user profile")
			}
			return resource.NewUserResource(user), nil
		})
	}	

func (c *UserController) UpdateUser(ctx *gin.Context) {
	c.Handle(ctx, func() (interface{}, *appErrors.AppError) {
		var req validations.UpdateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			return nil, appErrors.BadRequest(err.Error())
			}
			userID := ctx.GetUint("user_id")
			var user models.User
			if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
				return nil, appErrors.Internal("Failed to get user")
			}
			updates := map[string]interface{}{}
			if req.FirstName != "" {
				updates["first_name"] = req.FirstName
			}
			if req.Phone != "" {
				updates["phone_number"] = req.Phone
			}
			if req.NationalID != "" {
				updates["national_id"] = req.NationalID
			}
			if req.Email != "" {
				updates["email"] = req.Email
			}
			if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
				return nil, appErrors.Internal("Failed to update user")
			}
			return user, nil
	})
}

	func (c *UserController) AdminUpdateUser(ctx *gin.Context) {
		c.Handle(ctx, func() (interface{}, *appErrors.AppError) {
			
			var req validations.AdminUpdateRequest
			if err := ctx.ShouldBindJSON(&req); err != nil {
				return nil, appErrors.BadRequest(err.Error())
			}
			userID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
			if err != nil {
				return nil, appErrors.BadRequest("Invalid user ID")
			}
			var user models.User
			if err := config.DB.First(&user, userID).Error; err != nil {
				return nil, c.HandleDBError(err, "User not found")
			}
	
			updates := map[string]interface{}{}
	
			if req.FirstName != "" {
				updates["first_name"] = req.FirstName
			}
			if req.LastName != "" {
				updates["last_name"] = req.LastName
			}
			if req.Phone != "" {
				updates["phone_number"] = req.Phone
			}
			if req.NationalID != "" {
				updates["national_id"] = req.NationalID
			}
			if req.Email != "" {
				updates["email"] = req.Email
			}
			
			if len(updates) == 0 {
				return nil, appErrors.BadRequest("No fields to update")
			}
	
			if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
				return nil, appErrors.Internal("Failed to update user")
			}
	
			return user, nil
		})
	}

	func (c *UserController) DeleteUser(ctx *gin.Context) {
		c.Handle(ctx, func() (interface{}, *appErrors.AppError) {

			userID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
			if err != nil {
				return nil, appErrors.BadRequest("Invalid user ID")
			}
			if err := config.DB.Delete(&models.User{}, uint(userID)).Error; err != nil {
				return nil, appErrors.Internal("Failed to delete user")
			}
			return nil, nil
		})
	}
