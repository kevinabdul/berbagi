package middlewares

import (
	"net/http"
	"fmt"
	"strings"
	"strconv"
	"os"
	"berbagi/config"
	"berbagi/models"
	"berbagi/utils/response"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt"
)

func AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer") {
			return c.JSON(http.StatusBadRequest, struct {
				Status  string
				Message string
			}{Status: "Failed", Message: "Invalid Authorization type"})
		}
		
		tokenArr := strings.Split(authHeader, " ")

		if len(tokenArr) != 2 {
			return c.JSON(http.StatusBadRequest, struct {
				Status  string
				Message string
			}{Status: "Failed", Message: "Invalid JWT Format!"})
		}

		tokenString := strings.Split(tokenArr[1],".")

		if len(tokenString) != 3 {
			return c.JSON(http.StatusBadRequest, struct {
				Status  string
				Message string
			}{Status: "Failed", Message: "Invalid JWT Format!"})
		}

		valid, id, role, _ := checkToken(tokenArr[1])
		
		userId, _ := id.(float64)

		userRole,_ := role.(string)

		fmt.Println(fmt.Sprintf("id is %v. userId is %v", id, userId))

		if !valid {
			return c.JSON(http.StatusBadRequest, struct {
				Status  string
				Message string
			}{Status: "Failed", Message: "JWT is invalid!"})
		}

		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))
		c.Request().Header.Set("role", userRole)

		return next(c)
		}
}

func checkToken(tokenString string) (bool, interface{}, interface{}, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims["userId"], claims["role"], nil
	}

	return false, -1, "", err
}

func AuthorizeUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Request().Header.Get("role")
		method := c.Request().Header.Get("method")
		path := c.Request().Header.Get("path")

		if role == "admin" {
			return next(c)
		}

		rolePermission := models.PermissionAPI{}
		res := config.Db.Table("role_permissions").Select("roles.name as role_name, actions.name as action_name, resources.path as path").
		Joins("left join permissions on role_permissions.permission_id = permissions.id").
		Joins("left join resources on resources.id = permissions.resource_id").
		Joins("left join actions on actions.id = permissions.action_id").Joins("left join roles on roles.id = role_permissions.role_id").
		Where(`roles.name = ? and actions.name = ? and resources.path = ?`, role, method, path).Find(&rolePermission)

		if res.Error != nil {
			return c.JSON(http.StatusBadRequest,response.Create("failed", res.Error.Error(), nil))
		}

		if res.RowsAffected == 0 {
			return c.JSON(http.StatusUnauthorized, response.Create("failed", "You dont have permission to access this resource", nil))
		}
		
		return next(c)
		}
}