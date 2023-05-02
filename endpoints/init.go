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
	homeGroup.Get("/getDormAll", home.GetAllDorm)

	favoriteGroup := router.Group("/fav")
	favoriteGroup.Get("/test", favorite.Test)
	favoriteGroup.Post("/postFav", favorite.PostFav)
	favoriteGroup.Delete("/deleteFav", favorite.DeleteFav)

	authGroup := router.Group("/auth")
	authGroup.Get("/test", authentication.Test)

	mangeDormGroup := router.Group("/manage-dorm")
	mangeDormGroup.Get("/getDormDetail", mange_dorm.GetDormDetail)
	mangeDormGroup.Post("/postDorm", mange_dorm.PostDorm)
	mangeDormGroup.Delete("/deleteDorm", mange_dorm.DeleteDorm)
	mangeDormGroup.Get("/getDormImage", mange_dorm.GetDormImage)
	mangeDormGroup.Get("/getAllDorm", mange_dorm.GetAllDorm)

	mangeRoomGroup := router.Group("/manage-room")
	mangeRoomGroup.Get("/getRoomDetail", mange_room.GetRoomDetail)
	mangeRoomGroup.Post("/postRoom", mange_room.PostRoom)
	mangeRoomGroup.Delete("/deleteRoom", mange_room.DeleteRoom)
	mangeRoomGroup.Get("/getRoomImage", mange_room.GetRoomImage)
	mangeRoomGroup.Get("/getDormRoom", mange_room.GetDormRoom)

	profileGroup := router.Group("/profile")
	profileGroup.Get("/test", profile.Test)
	profileGroup.Get("/getProfile", profile.GetProfile)

	reviewGroup := router.Group("/review")
	reviewGroup.Get("/test", review.Test)
	reviewGroup.Get("/getDormReview", review.GetDormReview)
	reviewGroup.Post("/addDormReview", review.AddDormReview)
	reviewGroup.Delete("/deleteDormReview", review.DeleteDormReview)

}
