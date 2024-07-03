package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	authControllers "github.com/puvadon-artmit/gofiber-template/api/auth/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAuthRoutes(router fiber.Router) {
	app := router.Group("auth")
	app.Get("/get-all", middleware.AuthorizationRequired(), authControllers.GetAll)
	app.Post("/sign-up", middleware.AuthorizationRequired(), authControllers.SignUp)
	app.Post("/create-user", middleware.AuthorizationRequired(), authControllers.CreateNew)

	app.Post("/add-role", middleware.AuthorizationRequired(), authControllers.Create)

	app.Post("/sign-in", authControllers.SignInHandler)

	app.Post("/sign-in-cookie", authControllers.LoginDBHandler)

	app.Get("/get-profile", middleware.AuthorizationRequired(), authControllers.GetProfile)
	app.Get("/get-profile-cookie", authControllers.GetProfileCookie)
	app.Get("/get-by-profile/:userID", middleware.AuthorizationRequired(), authControllers.GetProfileById)
	// app.Get("/get-alluser", authControllers.GetAllUser)
	app.Get("/get-alluser", middleware.AuthorizationRequired(), authControllers.GetAllUsersandrole)
	app.Patch("/update-user/:user_id", middleware.AuthorizationRequired(), authControllers.UpdateUser)
	app.Delete("/delete-user/:user_id", middleware.AuthorizationRequired(), authControllers.DeleteUserByID)
	app.Delete("/delete-user-and-role/:user_user_id", authControllers.DeleteUserRolesByID)
	app.Get("/GetUserIdByToken", middleware.AuthorizationRequired(), authControllers.GetUserIdByToken)
	app.Get("/get-rold-code/:role_code", middleware.AuthorizationRequired(), authControllers.GetUsersByRoleIDHandler)
}
