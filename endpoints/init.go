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
	"Moddormy_backend/loaders/fiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Register(router fiber.Router) {

	uploadGroup := router.Group("/upload")
	uploadGroup.Post("/dorm", upload.Dorming)
	uploadGroup.Post("/room", upload.Rooming)

	homeGroup := router.Group("/home")
	homeGroup.Get("/test", home.Test)
	// case user already sign in bc it has fav status
	homeGroup.Get("/getDormAll", home.GetDormAll)
	homeGroup.Post("/postFilteredDorm", home.PostFilteredDorm)
	// case user does not sign in yet
	homeGroup.Get("/getAllDorm", home.GetAllDorm)
	homeGroup.Post("/postFilteredNoFav", home.PostFilteredNoFav)

	favoriteGroup := router.Group("/fav")
	favoriteGroup.Get("/test", favorite.Test)
	favoriteGroup.Post("/postFav", favorite.PostFav)
	favoriteGroup.Delete("/deleteFav", favorite.DeleteFav)
	favoriteGroup.Get("/getFav", favorite.GetFav)

	authGroup := router.Group("/auth")
	authGroup.Post("/register", authentication.Register)

	mangeDormGroup := router.Group("/manage-dorm")
	mangeDormGroup.Get("/getDormDetail", mange_dorm.GetDormDetail)
	mangeDormGroup.Post("/postDorm", mange_dorm.PostDorm)
	mangeDormGroup.Delete("/deleteDorm", mange_dorm.DeleteDorm)
	mangeDormGroup.Get("/getDormImage", mange_dorm.GetDormImage)
	mangeDormGroup.Put("/editDorm", mange_dorm.UpdateDorm)

	mangeRoomGroup := router.Group("/manage-room")
	mangeRoomGroup.Get("/getRoomDetail", mange_room.GetRoomDetail)
	mangeRoomGroup.Post("/postRoom", mange_room.PostRoom)
	mangeRoomGroup.Delete("/deleteRoom", mange_room.DeleteRoom)
	mangeRoomGroup.Get("/getRoomImage", mange_room.GetRoomImage)
	mangeRoomGroup.Get("/getDormRoom", mange_room.GetDormRoom)
	mangeRoomGroup.Put("/editRoom", mange_room.UpdateRoom)

	profileGroup := router.Group("/profile")
	profileGroup.Get("/test", profile.Test)
	profileGroup.Get("/getProfile", profile.GetProfile)
	profileGroup.Get("/getProfileDorm", profile.GetProfileDorm)
	profileGroup.Put("/editUser", profile.EditUser)

	reviewGroup := router.Group("/review")
	reviewGroup.Get("/test", review.Test)
	reviewGroup.Get("/getDormReview", review.GetDormReview)
	reviewGroup.Post("/addDormReview", middlewares.Jwt(), review.AddDormReview)
	reviewGroup.Delete("/deleteDormReview", review.DeleteDormReview)

}
