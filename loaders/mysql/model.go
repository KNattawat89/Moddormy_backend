package mysql

import (
	"Moddormy_backend/loaders/mysql/model"

	"gorm.io/gorm"
)

var DormFeatureModel *gorm.DB
var DormImagesModel *gorm.DB
var DormModel *gorm.DB
var FavoriteModel *gorm.DB
var FilesModel *gorm.DB
var RatingModel *gorm.DB
var ReviewsModel *gorm.DB
var RoomImagesModel *gorm.DB
var RoomFeaturesModel *gorm.DB
var RoomsModel *gorm.DB
var UsersModel *gorm.DB

func assignModel() {
	DormFeatureModel = Gorm.Model(new(model.DormFeature))
	DormImagesModel = Gorm.Model(new(model.DormImage))
	DormModel = Gorm.Model(new(model.Dorm))
	FavoriteModel = Gorm.Model(new(model.Favorite))
	FilesModel = Gorm.Model(new(model.File))
	RatingModel = Gorm.Model(new(model.Rating))
	ReviewsModel = Gorm.Model(new(model.Review))
	RoomImagesModel = Gorm.Model(new(model.RoomImage))
	RoomFeaturesModel = Gorm.Model(new(model.RoomFeature))
	RoomsModel = Gorm.Model(new(model.Room))
	UsersModel = Gorm.Model(new(model.User))

}
