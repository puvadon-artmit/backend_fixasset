package AuthServices

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	AuthDto "github.com/puvadon-artmit/gofiber-template/api/auth/entitys/request"
	authResponse "github.com/puvadon-artmit/gofiber-template/api/auth/entitys/response"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/puvadon-artmit/gofiber-template/utils"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

// func GetProfileById(id string) (Result *authResponse.ResultGetProfile, Error error) {
// 	db := database.DB
// 	result := authResponse.ResultGetProfile{}
// 	tx := db.Table("users").Select("user_id, username, firstname, lastname").Where("user_id=?", id).Scan(&result)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &result, nil
// }

func GetProfileById(userID string) (result []*model.User, err error) {
	db := database.DB
	var records []*model.User
	tx := db.Where("user_id = ?", userID).Preload("Roles.PermissionGroup.Permission").Preload("Roles.PermissionGroup.PermissionComponent").Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func GetProfileByCookieId(userID string) (*model.User, error) {
	db := database.DB
	var user model.User
	tx := db.Where("user_id = ?", userID).Preload("Roles.PermissionGroup.Permission").Preload("Roles.PermissionGroup.PermissionComponent").First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func SignIn(dto AuthDto.LoginDto) (token *string, message string, err error) {
	db := database.DB
	var user []model.User
	db.Find(&user, "username = ?", dto.Username)
	if len(user) == 0 {
		return nil, "username not found", errors.New("username not found")
	}
	if !utils.Compare(dto.Password, user[0].Password) {
		return nil, "password does not match", errors.New("password does not match")
	}
	if user[0].Status == "ปิดการใช้งาน" {
		return nil, "account is disabled", nil
	}
	tokenString, err := utils.GenerateJwt(user[0])
	if err != nil {
		return nil, "failed to generate token, please try again", errors.New("failed to generate token, please try again")
	}
	return &tokenString, "success", nil
}

func SignIncookie(dto AuthDto.LoginDto) (token string, err error) {
	db := database.DB
	var jwtSecretKey = []byte(viper.GetString("token.secret_key"))
	var user model.User
	if err := db.Where("username = ?", dto.Username).First(&user).Error; err != nil {
		return "", errors.New("username not found")
	}
	if !utils.Compare(dto.Password, user.Password) {
		return "", errors.New("password does not match")
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "fixed-asset",
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString(jwtSecretKey)
	if err != nil {
		return "", errors.New("failed to sign JWT")
	}

	return signedToken, nil
}

// func SignUp(user model.User) (result *model.User, Error error) {
// 	db := database.DB
// 	var userDB []model.User
// 	db.Find(&userDB, "username = ?", user.Username)
// 	if len(userDB) != 0 {
// 		return nil, errors.New("User Already has!")
// 	}
// 	var password = user.Password
// 	user.Password = utils.Encode(password)
// 	user.UserID = uuid.New().String()
// 	err := db.Create(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }

func SignUp(user model.User) (result *model.User, Error error) {
	db := database.DB
	var userDB []model.User
	db.Find(&userDB, "username = ?", user.Username)
	if len(userDB) != 0 {
		return nil, errors.New("user already exists")
	}
	var password = user.Password
	user.Password = utils.Encode(password)
	user.UserID = uuid.New().String()
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// func SignIn(dto AuthDto.LoginDto) (token *string, Error error) {
// 	db := database.DB
// 	var user []model.User
// 	db.Find(&user, "username = ?", dto.Username)
// 	if len(user) == 0 {
// 		return nil, errors.New("Username not found")
// 	}
// 	if !utils.Compare(dto.Password, user[0].Password) {
// 		return nil, errors.New("Password is not compare!")
// 	}
// 	tokenString, err := utils.GenerateJwt(user[0])
// 	if err != nil {
// 		return nil, errors.New("Generate token fail, try again.")
// 	}
// 	return &tokenString, nil
// }

func GetAll() (value *[]authResponse.ResultGetProfile, err error) {
	db := database.DB
	var result []authResponse.ResultGetProfile
	tx := db.Table("users").Select("user_id, username, firstname, lastname, role_active, employee_code, employee_code").Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

// func CreateNewUser(user model.User, roles []model.Role) (value *model.User, Error error) {
// 	db := database.DB
// 	user.UserID = uuid.New().String()

// 	user.Password = utils.Encode(user.Password)

// 	tx := db.Create(&user)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	for i := range roles {
// 		userroles := model.UserRoles{
// 			UserID: user.UserID,
// 			RoleID: roles[i].RoleID,
// 		}
// 		tx := db.Create(&userroles)
// 		if tx.Error != nil {
// 			return nil, tx.Error
// 		}
// 	}

// 	return &user, nil
// }

func CreateNewUser(user model.User, roles []model.Role) (value *model.User, Error error) {
	db := database.DB
	user.UserID = uuid.New().String()

	// เริ่ม transaction
	tx := db.Begin()

	var existingUser model.User
	if err := tx.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		// หากพบว่ามี username ที่ซ้ำกัน
		tx.Rollback() // ยกเลิก transaction
		return nil, errors.New("username is already taken")
	}

	// Hash the password
	user.Password = utils.Encode(user.Password)

	if err := tx.Create(&user).Error; err != nil {
		// หากเกิดข้อผิดพลาดในการสร้างผู้ใช้
		tx.Rollback() // ยกเลิก transaction
		return nil, err
	}

	for i := range roles {
		userroles := model.UserRoles{
			UserID: user.UserID,
			RoleID: roles[i].RoleID,
		}
		if err := tx.Create(&userroles).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// สำเร็จทั้งหมด ยืนยัน transaction
	if err := tx.Commit().Error; err != nil {
		// หากเกิดข้อผิดพลาดในการยืนยัน transaction
		tx.Rollback() // ยกเลิก transaction
		return nil, err
	}

	// ส่งข้อมูลผู้ใช้ที่สร้างไป
	return &user, nil
}

func GetAllUsers() (result []*model.User, err error) {
	db := database.DB
	var records []*model.User
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func UpdateUser(id string, updatedUser model.User) (value *model.User, Error error) {
// 	db := database.DB
// 	existingUser := model.User{UserID: id}
// 	tx := db.Preload(clause.Associations).Find(&existingUser)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	tx = db.Model(&existingUser).Updates(updatedUser)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &existingUser, nil
// }

func UpdateUser(id string, updatedUser model.User) (value *model.User, Error error) {
	db := database.DB
	existingUser := model.User{UserID: id}
	tx := db.Preload(clause.Associations).Find(&existingUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Hash the updated password if it exists
	if updatedUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updatedUser.Password = string(hashedPassword)
	}

	tx = db.Model(&existingUser).Updates(updatedUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingUser, nil
}

func UpdateUserRole(id string, updatedUser model.User, updatedRoles []model.Role) (*model.User, error) {
	db := database.DB
	existingUser := model.User{UserID: id}

	// Preload existing associations
	tx := db.Preload(clause.Associations).Find(&existingUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Hash the updated password if it exists
	if updatedUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updatedUser.Password = string(hashedPassword)
	}

	// Update user fields
	tx = db.Model(&existingUser).Updates(updatedUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Update roles
	if len(updatedRoles) > 0 {
		// Remove existing roles
		if err := db.Model(&existingUser).Association("Roles").Clear(); err != nil {
			return nil, err
		}

		// Add new roles
		for _, role := range updatedRoles {
			userRole := model.UserRoles{
				UserID: id,
				RoleID: role.RoleID,
			}
			if err := db.Create(&userRole).Error; err != nil {
				return nil, err
			}
		}
	}

	// Reload updated user with associations
	tx = db.Preload(clause.Associations).Find(&existingUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingUser, nil
}

func GetByCategoryIDDB(CategoryID string) ([]*model.User, error) {
	db := database.DB
	var category []*model.User
	tx := db.Preload(clause.Associations).Where("category_id = ?", CategoryID).Find(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return category, nil
}

func GetUsersByRoleID(RoleID string) (result []*model.User, err error) {
	db := database.DB
	var records []*model.User
	tx := db.
		Table("users").
		Select("users.*").
		Joins("INNER JOIN user_roles ON users.user_id = user_roles.user_user_id").
		Joins("INNER JOIN roles ON roles.role_id = user_roles.role_role_id").
		Where("roles.role_code = ?", RoleID).
		Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func DeleteUser(User *model.User) error {
	db := database.DB
	tx := db.Delete(User)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByUserId(id string) (value *model.User, Error error) {
	db := database.DB
	User := model.User{UserID: id}
	tx := db.Preload(clause.Associations).Find(&User)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &User, nil
}

func DeleteUserRolesByUserID(userID string) error {
	db := database.DB

	// ลบ UserRoles ที่เกี่ยวข้องกับ userID
	if err := db.Where("user_user_id = ?", userID).Unscoped().Delete(&model.UserRoles{}).Error; err != nil {
		return err
	}

	return nil
}

func CreateUserRoles(userID, roleID string) error {
	db := database.DB
	query := "INSERT INTO user_roles (user_user_id, role_role_id) VALUES (?, ?)"
	if err := db.Exec(query, userID, roleID).Error; err != nil {
		return err
	}
	return nil
}

func SignIncookieDB(dto AuthDto.LoginDto) (token string, err error) {
	db := database.DB
	var jwtSecretKey = []byte(viper.GetString("token.secret_key"))
	var user model.User
	if err := db.Where("username = ?", dto.Username).First(&user).Error; err != nil {
		return "", errors.New("username not found")
	}
	if !utils.Compare(dto.Password, user.Password) {
		return "", errors.New("password does not match")
	}

	// Create a custom claims object with user data
	claims := jwt.MapClaims{
		"user_id":       user.UserID,
		"username":      user.Username,
		"firstname":     user.Firstname,
		"lastname":      user.Lastname,
		"role_active":   user.RoleActive,
		"employee_code": user.EmployeeCode,
		"status":        user.Status,
		"iat":           time.Now().Unix(),                     // Issued at time
		"exp":           time.Now().Add(time.Hour * 24).Unix(), // Expiration time
	}

	// Create JWT token with claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	signedToken, err := jwtToken.SignedString(jwtSecretKey)
	if err != nil {
		return "", errors.New("failed to sign JWT")
	}

	return signedToken, nil
}
