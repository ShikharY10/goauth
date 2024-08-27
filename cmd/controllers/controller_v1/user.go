package controller_v1

import (
	"encoding/base64"
	"strconv"
	"time"

	"github.com/ShikharY10/goauth/cmd/handlers"
	"github.com/ShikharY10/goauth/cmd/middleware"
	"github.com/ShikharY10/goauth/cmd/models"
	"github.com/ShikharY10/goauth/cmd/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserController struct {
	Database *handlers.DataBase
	Cache    *handlers.CacheHandler
	Jwt      *middleware.JWT
}

// Signup godoc
// @Summary User singup
// @Description Create new user in db
// @Description It returns two token, access token and refresh token, access token is present in response body itself and refresh token is saved in http only cookie
// @Tags User
// @Param user body models.SignupRequest true "Create User"
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {string} bad request
// @Failure 500 {string} internal server error
// @Router /signup [post]
func (uc *UserController) SignUp(c *gin.Context) {

	var request models.SignupRequest
	c.BindJSON(&request)

	err := utils.ExamineSignupRequestBody(request)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
		return
	}

	_id := primitive.NewObjectID()

	var user models.User
	user.Id = _id
	user.Name = request.Name
	user.Username = request.Username

	hashedPassword, err := utils.HashWithSHA256([]byte(request.Password), 2)
	if err != nil {
		c.AbortWithStatus(500)
	}

	user.Password = base64.StdEncoding.EncodeToString(hashedPassword)
	user.Organisation = request.Organisation
	user.Role = "user"

	// generating access token
	accessClaim := map[string]interface{}{
		"id":       _id.Hex(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	accessToken, err := uc.Jwt.GenerateJWT(accessClaim, "access")
	if err != nil {
		c.AbortWithStatus(500)
	}

	// generating refresh token
	refreshClaim := map[string]interface{}{
		"id":  _id.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	refreshToken, err := uc.Jwt.GenerateJWT(refreshClaim, "refresh")
	if err != nil {
		c.AbortWithStatus(500)
	}

	uc.Cache.SetAccessTokenExpiry(_id.Hex(), accessToken, 1*time.Hour)
	uc.Cache.SetRefreshTokenExpiry(_id.Hex(), refreshToken, 24*time.Hour)

	c.SetCookie("refresh", refreshToken, 3600*24, "/", "", false, true)

	err = uc.Database.CreateNewUser(user)
	if err != nil {
		c.AbortWithStatus(500)
	}

	response := user.ToMap()
	response["token"] = accessToken

	c.JSON(200, response)
}

// Login godoc
// @Summary User login
// @Description Login by providing username and password
// @Tags User
// @Param user body models.LoginRequest true "login User"
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {string} bad request
// @Failure 500 {string} internal server error
// @Router /login [post]
func (uc *UserController) Login(c *gin.Context) {
	var request models.LoginRequest
	c.BindJSON(&request)

	err := utils.ExamineLoginRequestBody(request)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	hashedPassword, err := utils.HashWithSHA256([]byte(request.Password), 2)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	filter := bson.M{
		"username": request.Username,
		"password": base64.StdEncoding.EncodeToString(hashedPassword),
	}

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	user, err := uc.Database.GetUserData(filter, opts)
	if err != nil {
		c.AbortWithStatusJSON(400, "username or password is not correct: "+err.Error())
		return
	}

	// generating access token
	accessClaim := map[string]interface{}{
		"id":       user.Id.Hex(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	accessToken, err := uc.Jwt.GenerateJWT(accessClaim, "access")
	if err != nil {
		c.AbortWithStatus(500)
	}

	// generating refresh token
	refreshClaim := map[string]interface{}{
		"id":  user.Id.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	refreshToken, err := uc.Jwt.GenerateJWT(refreshClaim, "refresh")
	if err != nil {
		c.AbortWithStatus(500)
	}

	uc.Cache.SetAccessTokenExpiry(user.Id.Hex(), accessToken, time.Hour*1)
	uc.Cache.SetRefreshTokenExpiry(user.Id.Hex(), refreshToken, time.Hour*24)

	c.SetCookie("refresh", refreshToken, 3600*24, "/", "", false, true)

	response := user.ToMap()
	response["token"] = accessToken

	c.JSON(200, response)
}

// RefreshAccessToKen godoc
// @Summary Refresh the Access Token
// @Description Refresh Access Token using http only cookie "Refresh Token".
// @Tags User
// @Security Cookie
// @Param id path string true "user id"
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {string} bad request
// @Failure 401 {string} Unauthorized Access
// @Failure 500 {string} internal server error
// @Router /refresh/{id} [put]
func (uc *UserController) RefreshAccessToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh")
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(400)
		return
	}

	isTokenValid := uc.Cache.IsTokenValid(id, refreshToken, "refresh")
	if isTokenValid {
		refreshClaims, err := uc.Jwt.VarifyRefreshToken(refreshToken)
		if err != nil {
			c.AbortWithStatusJSON(401, "logged out")
			return
		}
		_id, err := primitive.ObjectIDFromHex(refreshClaims["id"].(string))
		if err != nil {
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		user, err := uc.Database.GetUserData(bson.M{"_id": _id}, nil)
		if err != nil {
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		newAccessTokenClaim := map[string]interface{}{
			"id":       user.Id.Hex(),
			"username": user.Username,
			"role":     user.Role,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		}
		accessToken, err := uc.Jwt.GenerateJWT(newAccessTokenClaim, "access")
		if err != nil {
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		uc.Cache.SetAccessTokenExpiry(user.Id.Hex(), accessToken, time.Hour*1)

		c.JSON(200, map[string]string{
			"accessToken": accessToken,
		})
	} else {
		c.AbortWithStatusJSON(401, "logged out")
		return
	}
}

// Logout godoc
// @Summary User logout
// @Description Logging out by unsetting http based refresh token cookie.
// @Description It also expires both access token and refresh token
// @Tags User
// @Security Bearer
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {string} Unauthorized Access
// @Router /logout [delete]
func (uc *UserController) Logout(c *gin.Context) {
	id := c.Value("id").(string)
	uc.Cache.DeleteTokenExpiry(id)
	c.SetCookie("refresh", "", -1, "/", "", true, true)
	c.JSON(200, "Successfully Logout")
}

// GetOneUser godoc
// @Summary User get one user
// @Description Get one user by providing user id as param
// @Tags User
// @Security Bearer
// @Param username path string true "username of user"
// @Produde json
// @Success 200 {object} models.User
// @Failure 400 {string} Bad Request
// @Failure 401 {string} Unauthorized Access
// @Failure 500 {string} internal server error
// @Router /user/{username} [get]
func (uc *UserController) GetOneUser(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.AbortWithStatusJSON(400, "username not found")
	} else {

		requesterId := c.Value("id").(string)
		requesterOrganisation, err := uc.Database.GetUserOrganisation(requesterId)
		if err != nil {
			c.AbortWithStatusJSON(500, err.Error())
		} else {

			user, err := uc.Database.GetUserData(bson.M{"username": username, "organisation": requesterOrganisation}, nil)
			if err != nil {
				c.AbortWithStatusJSON(400, "user is not present in your organisation")
			} else {
				c.JSON(200, user.ToMap())
			}
		}

	}
}

// GetMultipleUser godoc
// @Summary User get one user
// @Description Get one user by providing user id as param
// @Tags User
// @Security Bearer
// @Param l query string true "number of user"
// @Param p query string true "page number"
// @Produde json
// @Success 200 {object} models.Users
// @Failure 400 {string} Bad Request
// @Failure 401 {string} Unauthorized Access
// @Failure 500 {string} internal server error
// @Router /users [get]
func (uc *UserController) GetMultipleUsers(c *gin.Context) {
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
				admin, err := uc.Database.GetUserData(bson.M{"_id": _id}, nil)
				if err != nil {
					c.AbortWithStatusJSON(500, err.Error())
				} else {
					users, err := uc.Database.GetMultipleUsers(
						bson.M{
							"organisation": admin.Organisation,
							"_id":          bson.D{{Key: "$ne", Value: _id}},
						},
						findOptions,
						false,
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
