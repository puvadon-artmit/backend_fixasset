package authControllers

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	AuthDto "github.com/puvadon-artmit/gofiber-template/api/auth/entitys/request"
	AuthServices "github.com/puvadon-artmit/gofiber-template/api/auth/services"
	"github.com/puvadon-artmit/gofiber-template/config"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetUserIdByToken(c *fiber.Ctx) error {
	token := config.GetUser(c)
	return c.JSON(fiber.Map{
		"status": "success",
		"result": token,
	})
}

func GetAll(c *fiber.Ctx) error {
	db := database.DB
	var role model.Role
	type User struct {
		UserID string `json:"user_id"`
	}
	type RoleResult struct {
		RoleName        string `json:"role_name"`
		RoleDisplayName string `json:"role_display_name"`
		RoleDescription string `json:"role_description"`
		Users           []User `json:"users"`
	}
	var result RoleResult
	tx := db.Preload("Users").Find(&role).Scan(&result)
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  tx.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func GetProfile(c *fiber.Ctx) error {
	id := c.Locals("id")
	result, err := AuthServices.GetProfileById(id.(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func GetProfileCookie(c *fiber.Ctx) error {
	// ดึงค่า token จาก cookie ชื่อ "assert"
	cookie := c.Cookies("assert")

	// ตรวจสอบว่ามีค่า token หรือไม่
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	// ตรวจสอบและ verify token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Verify token method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(viper.GetString("token.secret_key")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// ตรวจสอบว่า token ถูกต้องหรือไม่
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userID := claims["user_id"].(string)

		result, err := AuthServices.GetProfileByCookieId(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// ส่งข้อมูลผู้ใช้กลับไปในรูปแบบ JSON
		return c.JSON(fiber.Map{
			"status": "success",
			"result": result,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized access",
	})
}

func GetAllUsersandrole(c *fiber.Ctx) error {
	value, err := AuthServices.GetAllUsers()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func GetProfileById(c *fiber.Ctx) error {
	userID := c.Params("userID")

	profiles, err := AuthServices.GetProfileById(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"profiles": profiles,
	})
}

func GetAllUser(c *fiber.Ctx) error {
	value, err := AuthServices.GetAll()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func SignUp(c *fiber.Ctx) error {
	user := new(model.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	_, err = AuthServices.SignUp(*user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Signup successfully.",
	})
}

func SignInHandler(c *fiber.Ctx) error {
	var dto AuthDto.LoginDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "failed to parse request body",
		})
	}
	if err := validator.New().Struct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	tokenString, message, err := AuthServices.SignIn(dto)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  message,
		})
	}
	if message != "success" {
		return c.Status(200).JSON(fiber.Map{
			"status": "error",
			"error":  message,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"token":  tokenString,
	})
}

// var jwtSecretKey = []byte("theball")

// func authRequired(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")

// 	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecretKey, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	return c.Next()
// }

func CreateNew(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	var roles []model.Role

	createdUser, err := AuthServices.CreateNewUser(user, roles)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   createdUser,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("user_id")

	updatedUser := new(model.User)
	err := c.BodyParser(updatedUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := AuthServices.UpdateUser(id, *updatedUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": updatedStatus,
	})
}

func GetUsersByRoleIDHandler(c *fiber.Ctx) error {
	RoleID := c.Params("role_code")

	users, err := AuthServices.GetUsersByRoleID(RoleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"result": users,
	})
}

func DeleteUserByID(c *fiber.Ctx) error {
	User_id := c.Params("user_id")

	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ต้องการลบ
	user, err := AuthServices.GetByUserId(User_id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = AuthServices.DeleteUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
	})
}

func DeleteUserRolesByID(c *fiber.Ctx) error {
	userID := c.Params("user_user_id")

	err := AuthServices.DeleteUserRolesByUserID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User roles deleted successfully",
	})
}

func Create(c *fiber.Ctx) error {
	userroles := struct {
		UserID string `json:"user_user_id"`
		RoleID string `json:"role_role_id"`
	}{}
	err := c.BodyParser(&userroles)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// ตรวจสอบค่า UUID ที่ถูกต้องสำหรับ user_user_id และ role_role_id
	if _, err := uuid.Parse(userroles.UserID); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid user_user_id",
		})
	}

	if _, err := uuid.Parse(userroles.RoleID); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid role_role_id",
		})
	}

	if err := AuthServices.CreateUserRoles(userroles.UserID, userroles.RoleID); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": userroles,
	})
}

func LoginHandler(c *fiber.Ctx) error {
	var loginData AuthDto.LoginDto
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, err := AuthServices.SignIncookie(loginData)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "assert",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Login": "Successful",
	})
}

func LoginDBHandler(c *fiber.Ctx) error {
	var loginData AuthDto.LoginDto
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, err := AuthServices.SignIncookieDB(loginData)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Set cookie with token
	c.Cookie(&fiber.Cookie{
		Name:     "assert",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "none",
	})

	// Return successful login response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Login": "Successful",
	})
}

// func LoginDBHandler(c *fiber.Ctx) error {
// 	var loginData AuthDto.LoginDto
// 	if err := c.BodyParser(&loginData); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body",
// 		})
// 	}

// 	token, err := AuthServices.SignIncookieDB(loginData)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	// Set cookie with token
// 	c.Cookie(&fiber.Cookie{
// 		Name:     "assert",
// 		Value:    token,
// 		Expires:  time.Now().Add(time.Hour * 24),
// 		HTTPOnly: true,
// 		Secure:   true,
// 		Domain:   "localhost",
// 		Path:     "/",
// 	})

// 	// Return successful login response
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"Login": "Successful",
// 	})
// }
