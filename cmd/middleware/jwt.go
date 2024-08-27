// This files contains all the function and code used to authorize a authenticated user.
package middleware

import (
	"errors"
	"net/http"

	config "github.com/HousewareHQ/backend-engineering-octernship/cmd/configs"
	"github.com/HousewareHQ/backend-engineering-octernship/cmd/handlers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	jWT_ACCESS_TOKEN_SECRET_KEY  []byte
	jWT_REFRESH_TOKEN_SECRET_KEY []byte
	DataBase                     *handlers.DataBase
	Cache                        *handlers.CacheHandler
}

// Initializes JWT struct
func CreateJWT(env *config.ENV, database *handlers.DataBase, cache *handlers.CacheHandler) *JWT {
	return &JWT{
		jWT_ACCESS_TOKEN_SECRET_KEY:  []byte(env.JWT_ACCESS_TOKEN_SECRET_KEY),
		jWT_REFRESH_TOKEN_SECRET_KEY: []byte(env.JWT_REFRESH_TOKEN_SECRET_KEY),
		DataBase:                     database,
		Cache:                        cache,
	}
}

// Creates a JWT token using SHA256 hashing algorithm.
func (j *JWT) GenerateJWT(claim map[string]interface{}, Tokentype string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for k, v := range claim {
		claims[k] = v
	}

	var key []byte
	if Tokentype == "access" {
		key = j.jWT_ACCESS_TOKEN_SECRET_KEY
	} else {
		key = j.jWT_REFRESH_TOKEN_SECRET_KEY
	}

	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// varifies JWT access token and the claims the where set while creating the token
func (j *JWT) VarifyAccessToken(token string) (claim jwt.MapClaims, err error) {
	newToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("something went wrong")
		}
		return j.jWT_ACCESS_TOKEN_SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := newToken.Claims.(jwt.MapClaims); ok && newToken.Valid {
		return claims, nil
	} else {
		return nil, errors.New("bad token")
	}
}

// varifies JWT refresh token and the claims the where set while creating the token
func (j *JWT) VarifyRefreshToken(token string) (claim jwt.MapClaims, err error) {
	newToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("something went wrong")
		}
		return j.jWT_REFRESH_TOKEN_SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := newToken.Claims.(jwt.MapClaims); ok && newToken.Valid {
		return claims, nil
	} else {
		return nil, errors.New("bad token")
	}
}

// Middleware for authorizing user using access token
func (j *JWT) APIV1_Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		if bearer == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, "token not found")
			return
		} else {
			token := bearer[len("Bearer "):]
			if token == "" {
				c.AbortWithStatusJSON(http.StatusForbidden, "token not found")
				return
			} else {
				claim, err := j.VarifyAccessToken(token)
				if err != nil {
					if err.Error() == "Token is expired" {
						c.AbortWithStatusJSON(401, err.Error())
					} else {
						c.AbortWithStatus(400)
					}
				} else {
					isTokenValid := j.Cache.IsTokenValid(claim["id"].(string), token, "access")
					if isTokenValid {
						data := map[string]interface{}{
							"id":       claim["id"].(string),
							"username": claim["username"].(string),
							"role":     claim["role"].(string),
						}
						c.Keys = data
						c.Next()
					} else {
						c.AbortWithStatus(401)
					}

				}
			}
		}
	}
}
