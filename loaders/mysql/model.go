package mysql

import (
	"Moddormy_backend/loaders/mysql/model"
	"gorm.io/gorm"
)

var DormImagesModel *gorm.DB
var DormModel *gorm.DB
var FavoriteModel *gorm.DB
var ReviewsModel *gorm.DB
var RoomImagesModel *gorm.DB
var RoomsModel *gorm.DB
var UsersModel *gorm.DB

func assignModel() {
	DormImagesModel = Gorm.Model(new(model.DormImage))
	DormModel = Gorm.Model(new(model.Dorm))
	FavoriteModel = Gorm.Model(new(model.Favorite))
	ReviewsModel = Gorm.Model(new(model.Review))
	RoomImagesModel = Gorm.Model(new(model.RoomImage))
	RoomsModel = Gorm.Model(new(model.Room))
	UsersModel = Gorm.Model(new(model.User))

}
