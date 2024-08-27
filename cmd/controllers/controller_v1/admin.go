// This folder contains controllers that will be used by admin privileged users.
package controller_v1

import (
	"encoding/base64"
	"strconv"

	"github.com/ShikharY10/goauth/cmd/handlers"
	"github.com/ShikharY10/goauth/cmd/models"
	"github.com/ShikharY10/goauth/cmd/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AdminController struct {
	Database *handlers.DataBase
}

// AdminGetOneUserData godoc
// @Summary Admin get one user
// @Description Used by admins to retrive complete data of a user using user id.
// @Tags Admin
// @Security Bearer
// @Param id path string true "user id"
// @Produde json
// @Success 200 {object} models.User
// @Failure 401 {string} Unauthorized Access
// @Failure 500 {string} internal server error
// @Router /admin/user/{id} [get]
func (ac *AdminController) GetOneUserData(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(400, "user id not found")
	} else {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.AbortWithStatus(500)
		} else {
			user, err := ac.Database.GetUserData(bson.M{"_id": _id}, nil)
			if err != nil {
				c.AbortWithStatus(500)
			} else {
				c.JSON(200, user)
			}
		}

	}
}

// AdminGetAllUserData godoc
// @Summary Admin Get Multiple user data
// @Description Admins multiple user data based on page and limit specified in query.
// @Tags Admin
// @Security Bearer
// @Param l query string true "number of user"
// @Param p query string true "page number"
// @Produce json
// @Success 200 {object} models.Users
// @Failure 400 {string} Bad Request
// @Failure 401 {string} Unauthorized Access
// @Failure 500 {string} internal server error
// @Router /admin/users [get]
func (ac *AdminController) GetAllUserData(c *gin.Context) {
	page := c.Query("p")
	limit := c.Query("l")

	if page == "" || limit == "" {
		c.AbortWithStatusJSON(400, "page or limit not found")
	} else {
		l, errl := strconv.Atoi(limit)
		p, errp := strconv.Atoi(page)

		if errl != nil || errp != nil {
			c.AbortWithStatusJSON(400, "page and limit not in correct format")
		} else {
			limit64 := int64(l)
			skip64 := int64(p*l - l)
			findOptions := options.FindOptions{Limit: &limit64, Skip: &skip64}

			_id, err := primitive.ObjectIDFromHex(c.Value("id").(string))
			if err != nil {
				c.AbortWithStatusJSON(500, err.Error())
			} else {
				admin, err := ac.Database.GetUserData(bson.M{"_id": _id}, nil)
				if err != nil {
					c.AbortWithStatusJSON(500, err.Error())
				} else {
					users, err := ac.Database.GetMultipleUsers(
						bson.M{
							"organisation": admin.Organisation,
							"_id":          bson.D{{Key: "$ne", Value: _id}},
						},
						findOptions,
						true,
					)
					if err != nil {
						c.JSON(404, []string{})
					} else {
						c.JSON(200, users)
					}
				}
			}
		}
	}
}

// AdminCreateNewUser godoc
// @Summary Admin creates new user
// @Description Admin creates new user. It does not return any type token.
// @Description In order get the token user needs to login.
// @Tags Admin
// @Security Bearer
// @Param user body models.SignupRequest true "Create User"
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {string} bad request
// @Failure 401 {string} Unauthorized Access
// @Failure 500 {string} internal server error
// @Router /admin/user [post]
func (ac *AdminController) CreateNewUser(c *gin.Context) {
	var request models.SignupRequest
	c.BindJSON(&request)

	err := utils.ExamineSignupRequestBody(request)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
		return
	}

	adminOrganisation, err := ac.Database.GetUserOrganisation(c.Value("id").(string))
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	if adminOrganisation != request.Organisation {
		c.AbortWithStatusJSON(400, "you are trying to add user in other organisation")
		return
	}

	_id := primitive.NewObjectID()

	var user models.User
	user.Id = _id
	user.Name = request.Name
	user.Username = request.Username

	hashedPassword, err := utils.HashWithSHA256([]byte(request.Password), 2)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	user.Password = base64.StdEncoding.EncodeToString(hashedPassword)
	user.Organisation = request.Organisation
	user.Role = "user"

	err = ac.Database.CreateNewUser(user)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	user.Password = ""
	c.JSON(200, user)
}

// AdminDeleteOneUser godoc
// @Summary Admin delete one user
// @Description Admin delete one user based on user id and same organisation
// @Tags Admin
// @Security Bearer
// @Param id path string true "User id"
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {string} bad request
// @Failure 401 {string} Unauthorized Access
// @Failure 500 {string} internal server error
// @Router /admin/user/{id} [delete]
func (ac *AdminController) DeleteOneUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(400, "user id not found")
	} else {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.AbortWithStatusJSON(400, "invalid user id")
		} else {
			_idAdmin, err := primitive.ObjectIDFromHex(c.Value("id").(string))
			if err != nil {
				c.AbortWithStatusJSON(500, err.Error())
			} else {
				admin, err := ac.Database.GetUserData(bson.M{"_id": _idAdmin}, nil)
				if err != nil {
					c.AbortWithStatusJSON(500, "1"+err.Error())
				} else {
					user, err := ac.Database.DeleteUser(bson.M{"_id": _id, "organisation": admin.Organisation})
					if err != nil {
						c.AbortWithStatusJSON(500, "2"+err.Error())
					} else {
						c.JSON(200, user)
					}
				}
			}
		}

	}
}

// AdminCreateNewAdmin godoc
// @Summary Admin create new admin
// @Description Admin create new admin inside there organisation
// @Tags Admin
// @Security Bearer
// @Param id path string true "user id"
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {string} bad request
// @Failure 401 {string} Unauthorized Access
// @Failure 500 {string} internal server error
// @Router /admin/new/{id} [put]
func (ac *AdminController) CreateNewAdmin(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(400, "user id not found")
	} else {
		userOrganisation, errU := ac.Database.GetUserOrganisation(id)
		if errU != nil {
			c.AbortWithStatusJSON(400, "invalid user id")
		} else {
			adminOrganisation, errA := ac.Database.GetUserOrganisation(c.Value("id").(string))
			if errA != nil {
				c.AbortWithStatusJSON(500, errA.Error())
			} else {
				if userOrganisation == adminOrganisation {
					err := ac.Database.ChangeRole(id, "admin")
					if err != nil {
						c.AbortWithStatusJSON(500, err.Error())
					} else {
						c.JSON(200, "Successfully updated role")
					}
				} else {
					c.AbortWithStatusJSON(400, "you are trying to change role of user who are not the part of you organisation")
				}
			}
		}
	}
}
