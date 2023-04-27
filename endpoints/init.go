package endpoints

import (
	"Moddormy_backend/endpoints/authentication"
	"Moddormy_backend/endpoints/favorite"
	"Moddormy_backend/endpoints/home"
	"Moddormy_backend/endpoints/mange_dorm"
	"Moddormy_backend/endpoints/mange_room"
	"Moddormy_backend/endpoints/profile"
	"Moddormy_backend/endpoints/review"
	"Moddormy_backend/endpoints/upload"

	"github.com/gofiber/fiber/v2"
)

func Register(router fiber.Router) {

	uploadGroup := router.Group("/upload")
	uploadGroup.Post("/dorm", upload.Dorming)
	uploadGroup.Post("/room", upload.Rooming)

	homeGroup := router.Group("/home")
	homeGroup.Get("/test", home.Test)

	favoriteGroup := router.Group("/fav")
	favoriteGroup.Get("/test", favorite.Test)

	authGroup := router.Group("/auth")
	authGroup.Get("/test", authentication.Test)

	mangeDormGroup := router.Group("/manage-dorm")
	mangeDormGroup.Get("/test", mange_dorm.Test)

	mangeRoomGroup := router.Group("/manage-room")
	mangeRoomGroup.Get("/test", mange_room.Test)

	profileGroup := router.Group("/profile")
	profileGroup.Get("/test", profile.Test)

	reviewGroup := router.Group("/review")
	reviewGroup.Get("/test", review.Test)
	reviewGroup.Get("/getDormReview", review.GetDormReview)
	reviewGroup.Post("/addDormReview", review.AddDormReview)

}
