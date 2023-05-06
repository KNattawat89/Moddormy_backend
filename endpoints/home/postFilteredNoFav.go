package home

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/config"
	"Moddormy_backend/utils/value"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"sort"
	"strings"
)

func PostFilteredNoFav(c *fiber.Ctx) error {

	body := new(payload.FilterNoFav)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}
	//facilities
	filter_parking := true
	filter_wifi := true
	filter_smoke_free := true
	filter_guard := true
	filter_pet := true
	filter_air := true
	filter_fan := true

	var dorms []model.Dorm

	if *body.Distant > 0 {
		if len(body.Facilities) > 0 {
			//have fac + dis
			if len(body.Facilities) == 1 {
				if strings.EqualFold(*body.Facilities[0], "parking") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ?", filter_parking).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ?", filter_wifi).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("smoke_free = ?", filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("security_guard = ?", filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("pet = ?", filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? ", filter_air).Preload("Reviews").Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm that have air-conditioner room",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ? ", filter_fan).Preload("Reviews").Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm that have fan room",
							Err:     result.Error,
						}
					}
				}
			} else if len(body.Facilities) == 2 {
				//21
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND wifi = ?", filter_parking, filter_wifi).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and wifi",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND smoke_free = ?", filter_parking, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and smoke_free",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND security_guard = ?", filter_parking, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND pet = ?", filter_parking, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("parking = ?", filter_parking).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("parking = ?", filter_parking).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND smoke_free = ?", filter_wifi, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and smoke_free",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND security_guard = ?", filter_wifi, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND pet = ?", filter_wifi, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("wifi = ?", filter_wifi).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("wifi = ?", filter_wifi).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("smoke_free = ? AND security_guard = ?", filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free and security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("smoke_free = ? AND pet = ?", filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free and pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("smoke_free = ?", filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("smoke_free = ?", filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("security_guard = ? AND pet = ?", filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard and pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("security_guard = ?", filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("security_guard = ?", filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "pet friendly") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("pet = ?", filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with pet friendly and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "pet friendly") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("pet = ?", filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with pet friendly and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "air-conditioner") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan = ?", filter_air, filter_fan).Preload("Reviews").Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with air-conditioner and fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 3 {
				//35
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND wifi = ? AND smoke_free =?", filter_parking, filter_wifi, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND wifi = ? AND security_guard =?", filter_parking, filter_wifi, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND wifi = ? AND pet =?", filter_parking, filter_wifi, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("parking = ? AND wifi = ?", filter_parking, filter_wifi).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, airc",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("parking = ? AND wifi = ?", filter_parking, filter_wifi).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND smoke_free =? AND security_guard =?", filter_parking, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND smoke_free =? AND pet =?", filter_parking, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking = ? AND smoke_free =?", filter_parking, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking = ? AND smoke_free =?", filter_parking, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND security_guard =? AND pet =?", filter_parking, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking = ? AND security_guard =?", filter_parking, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, airc",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking = ? AND security_guard =?", filter_parking, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking = ? AND pet =?", filter_parking, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, pet friendly, airc",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking = ? AND pet =?", filter_parking, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking = ? ", filter_parking).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND smoke_free =? AND security_guard = ?", filter_wifi, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND smoke_free =? AND pet = ?", filter_wifi, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("wifi = ? AND smoke_free =?", filter_wifi, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, air",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("wifi = ? AND smoke_free =?", filter_wifi, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND security_guard = ? AND pet =?", filter_wifi, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc =?", filter_air).Preload("Reviews").Where("wifi = ? AND security_guard = ?", filter_wifi, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan =?", filter_fan).Preload("Reviews").Where("wifi = ? AND security_guard = ?", filter_wifi, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc =?", filter_air).Preload("Reviews").Where("wifi = ? AND pet = ?", filter_wifi, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan =?", filter_fan).Preload("Reviews").Where("wifi = ? AND pet = ?", filter_wifi, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc =? OR fan =?", filter_air, filter_fan).Preload("Reviews").Where("wifi = ? ", filter_wifi).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("smoke_free = ? AND security_guard = ? AND pet =?", filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("smoke_free = ? AND security_guard = ?", filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("smoke_free = ? AND security_guard = ?", filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("smoke_free = ? AND pet = ?", filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, pet, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("smoke_free = ? AND pet = ?", filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, pet, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan =?", filter_air, filter_fan).Preload("Reviews").Where("smoke_free = ?", filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("security_guard = ? AND pet = ?", filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard, pet, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("security_guard = ? AND pet = ?", filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard, pet, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan = ?", filter_air, filter_fan).Preload("Reviews").Where("security_guard = ?", filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "pet friendly") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan = ?", filter_air, filter_fan).Preload("Reviews").Where("pet = ?", filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 4 {
				//35
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc =?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =?", filter_parking, filter_wifi, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan =?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =?", filter_parking, filter_wifi, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? AND pet=?", filter_parking, filter_wifi, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =?", filter_parking, filter_wifi, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =?", filter_parking, filter_wifi, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND pet =?", filter_parking, filter_wifi, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND pet =?", filter_parking, filter_wifi, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? ", filter_parking, filter_wifi).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? AND pet=?", filter_parking, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =?", filter_parking, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =?", filter_parking, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND smoke_free=? AND pet =?", filter_parking, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND pet =?", filter_parking, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? ", filter_parking, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND security_guard =? AND pet=?", filter_parking, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND security_guard =? AND pet=?", filter_parking, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND security_guard =? ", filter_parking, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], " pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND pet =? ", filter_parking, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking,  pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? AND pet=?", filter_wifi, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? ", filter_wifi, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? ", filter_wifi, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("wifi =? AND smoke_free=? AND pet =? ", filter_wifi, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND pet =? ", filter_wifi, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=?", filter_wifi, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("wifi =? AND security_guard=? AND pet =? ", filter_wifi, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("wifi =? AND security_guard=? AND pet =? ", filter_wifi, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND security_guard=?", filter_wifi, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND pet=?", filter_wifi, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("smoke_free =? AND security_guard=? AND pet =? ", filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("smoke_free =? AND security_guard=? AND pet =? ", filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditione") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("smoke_free =? AND security_guard=?", filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, air-conditione, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditione") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("smoke_free =? AND pet=?", filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, pet friendly, air-conditione, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditione") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("security_guard =? AND pet=?", filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard, pet friendly, air-conditione, fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 5 {
				//21
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, pet friendly,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =?", filter_parking, filter_wifi, filter_smoke_free).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? AND pet=?", filter_parking, filter_wifi, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, pet friendly,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? AND pet=?", filter_parking, filter_wifi, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? ", filter_parking, filter_wifi, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND pet =? ", filter_parking, filter_wifi, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? AND pet=?", filter_parking, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, pet friendly,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? AND pet=?", filter_parking, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, pet friendly,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? ", filter_parking, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, air-conditioner,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND pet =? ", filter_parking, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, pet friendly, air-conditioner,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND security_guard=? AND pet =? ", filter_parking, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, pet friendly, air-conditioner,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? AND pet=?", filter_wifi, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, pet friendly,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? AND pet=?", filter_wifi, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? ", filter_wifi, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND pet =? ", filter_wifi, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND security_guard=? AND pet =? ", filter_wifi, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi,security guard, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("smoke_free =? AND security_guard=? AND pet =? ", filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free,security guard, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 6 {
				//7
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "pet friendly") && strings.EqualFold(*body.Facilities[5], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, pet friendly,air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "pet friendly") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, pet friendly,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, pet friendly, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? AND pet=?", filter_parking, filter_wifi, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, pet friendly, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? AND pet=?", filter_parking, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, pet friendly, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? AND pet=?", filter_wifi, filter_smoke_free, filter_guard, filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, pet friendly, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 7 {
				//1
				if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan =?", filter_air, filter_fan).Preload("Reviews").Where("parking = ?", filter_parking).Where("wifi = ?", filter_wifi).Where("smoke_free = ?", filter_smoke_free).Where("security_guard = ?", filter_guard).Where("pet = ?", filter_pet).Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
					return &response.GenericError{
						Message: "Dorm not found",
						Err:     result.Error,
					}
				}

			}

		} else {
			// no fac, have distant
			if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("distant <= ?", *body.Distant).Find(&dorms); result.Error != nil {
				return &response.GenericError{
					Message: "Dorms not found in this distant",
					Err:     result.Error,
				}
			}
		}

	} else {
		if len(body.Facilities) > 0 {
			// no distant, have fac
			if len(body.Facilities) == 1 {
				if strings.EqualFold(*body.Facilities[0], "parking") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ?", filter_parking).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ?", filter_wifi).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("smoke_free = ?", filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("security_guard = ?", filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("pet = ?", filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? ", filter_air).Preload("Reviews").Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm that have air-conditioner room",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ? ", filter_fan).Preload("Reviews").Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm that have fan room",
							Err:     result.Error,
						}
					}
				}
			} else if len(body.Facilities) == 2 {
				//21
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND wifi = ?", filter_parking, filter_wifi).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and wifi",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND smoke_free = ?", filter_parking, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and smoke_free",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND security_guard = ?", filter_parking, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND pet = ?", filter_parking, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("parking = ?", filter_parking).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("parking = ?", filter_parking).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND smoke_free = ?", filter_wifi, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and smoke_free",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND security_guard = ?", filter_wifi, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND pet = ?", filter_wifi, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("wifi = ?", filter_wifi).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("wifi = ?", filter_wifi).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("smoke_free = ? AND security_guard = ?", filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free and security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("smoke_free = ? AND pet = ?", filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free and pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("smoke_free = ?", filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("smoke_free = ?", filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("security_guard = ? AND pet = ?", filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard and pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("security_guard = ?", filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("security_guard = ?", filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "pet friendly") && strings.EqualFold(*body.Facilities[1], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("pet = ?", filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with pet friendly and air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "pet friendly") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("pet = ?", filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with pet friendly and fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "air-conditioner") && strings.EqualFold(*body.Facilities[1], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan = ?", filter_air, filter_fan).Preload("Reviews").Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with air-conditioner and fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 3 {
				//35
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND wifi = ? AND smoke_free =?", filter_parking, filter_wifi, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND wifi = ? AND security_guard =?", filter_parking, filter_wifi, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND wifi = ? AND pet =?", filter_parking, filter_wifi, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("parking = ? AND wifi = ?", filter_parking, filter_wifi).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, airc",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("parking = ? AND wifi = ?", filter_parking, filter_wifi).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND smoke_free =? AND security_guard =?", filter_parking, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND smoke_free =? AND pet =?", filter_parking, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking = ? AND smoke_free =?", filter_parking, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking = ? AND smoke_free =?", filter_parking, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking = ? AND security_guard =? AND pet =?", filter_parking, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking = ? AND security_guard =?", filter_parking, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, airc",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking = ? AND security_guard =?", filter_parking, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking = ? AND pet =?", filter_parking, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, pet friendly, airc",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking = ? AND pet =?", filter_parking, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking = ? ", filter_parking).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND smoke_free =? AND security_guard = ?", filter_wifi, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND smoke_free =? AND pet = ?", filter_wifi, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("wifi = ? AND smoke_free =?", filter_wifi, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, air",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("wifi = ? AND smoke_free =?", filter_wifi, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi = ? AND security_guard = ? AND pet =?", filter_wifi, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc =?", filter_air).Preload("Reviews").Where("wifi = ? AND security_guard = ?", filter_wifi, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan =?", filter_fan).Preload("Reviews").Where("wifi = ? AND security_guard = ?", filter_wifi, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc =?", filter_air).Preload("Reviews").Where("wifi = ? AND pet = ?", filter_wifi, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan =?", filter_fan).Preload("Reviews").Where("wifi = ? AND pet = ?", filter_wifi, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc =? OR fan =?", filter_air, filter_fan).Preload("Reviews").Where("wifi = ? ", filter_wifi).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("smoke_free = ? AND security_guard = ? AND pet =?", filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("smoke_free = ? AND security_guard = ?", filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("smoke_free = ? AND security_guard = ?", filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("smoke_free = ? AND pet = ?", filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, pet, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("smoke_free = ? AND pet = ?", filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, pet, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan =?", filter_air, filter_fan).Preload("Reviews").Where("smoke_free = ?", filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ?", filter_air).Preload("Reviews").Where("security_guard = ? AND pet = ?", filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard, pet, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan = ?", filter_fan).Preload("Reviews").Where("security_guard = ? AND pet = ?", filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard, pet, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan = ?", filter_air, filter_fan).Preload("Reviews").Where("security_guard = ?", filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "pet friendly") && strings.EqualFold(*body.Facilities[1], "air-conditioner") && strings.EqualFold(*body.Facilities[2], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan = ?", filter_air, filter_fan).Preload("Reviews").Where("pet = ?", filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 4 {
				//35
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc =?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =?", filter_parking, filter_wifi, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan =?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =?", filter_parking, filter_wifi, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? AND pet=?", filter_parking, filter_wifi, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =?", filter_parking, filter_wifi, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =?", filter_parking, filter_wifi, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND pet =?", filter_parking, filter_wifi, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND pet =?", filter_parking, filter_wifi, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? ", filter_parking, filter_wifi).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? AND pet=?", filter_parking, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =?", filter_parking, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =?", filter_parking, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND smoke_free=? AND pet =?", filter_parking, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND pet =?", filter_parking, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? ", filter_parking, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND security_guard =? AND pet=?", filter_parking, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND security_guard =? AND pet=?", filter_parking, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND security_guard =? ", filter_parking, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], " pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND pet =? ", filter_parking, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking,  pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? AND pet=?", filter_wifi, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? ", filter_wifi, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? ", filter_wifi, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("wifi =? AND smoke_free=? AND pet =? ", filter_wifi, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND pet =? ", filter_wifi, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=?", filter_wifi, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("wifi =? AND security_guard=? AND pet =? ", filter_wifi, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("wifi =? AND security_guard=? AND pet =? ", filter_wifi, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND security_guard=?", filter_wifi, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditioner") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND pet=?", filter_wifi, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("smoke_free =? AND security_guard=? AND pet =? ", filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, pet friendly, air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("smoke_free =? AND security_guard=? AND pet =? ", filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "air-conditione") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("smoke_free =? AND security_guard=?", filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, security guard, air-conditione, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditione") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("smoke_free =? AND pet=?", filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free, pet friendly, air-conditione, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "security guard") && strings.EqualFold(*body.Facilities[1], "pet friendly") && strings.EqualFold(*body.Facilities[2], "air-conditione") && strings.EqualFold(*body.Facilities[3], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("security_guard =? AND pet=?", filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with security guard, pet friendly, air-conditione, fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 5 {
				//21
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "pet friendly") {
					if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, pet friendly",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, pet friendly,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =?", filter_parking, filter_wifi, filter_smoke_free).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? AND pet=?", filter_parking, filter_wifi, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, pet friendly,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? AND pet=?", filter_parking, filter_wifi, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? ", filter_parking, filter_wifi, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND pet =? ", filter_parking, filter_wifi, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? AND pet=?", filter_parking, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, pet friendly,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? AND pet=?", filter_parking, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, pet friendly,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? ", filter_parking, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, air-conditioner,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND pet =? ", filter_parking, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, pet friendly, air-conditioner,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND security_guard=? AND pet =? ", filter_parking, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, security guard, pet friendly, air-conditioner,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? AND pet=?", filter_wifi, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, pet friendly,  air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? AND pet=?", filter_wifi, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, pet friendly, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? ", filter_wifi, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND pet =? ", filter_wifi, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND security_guard=? AND pet =? ", filter_wifi, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi,security guard, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "smoke-free") && strings.EqualFold(*body.Facilities[1], "security guard") && strings.EqualFold(*body.Facilities[2], "pet friendly") && strings.EqualFold(*body.Facilities[3], "air-conditioner") && strings.EqualFold(*body.Facilities[4], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("smoke_free =? AND security_guard=? AND pet =? ", filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with smoke-free,security guard, pet friendly, air-conditioner, fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 6 {
				//7
				if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "pet friendly") && strings.EqualFold(*body.Facilities[5], "air-conditioner") {
					if result := mysql.Gorm.Preload("Rooms", "airc=?", filter_air).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, pet friendly,air-conditioner",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "pet friendly") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "fan=?", filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, pet friendly,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "security guard") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND security_guard=?", filter_parking, filter_wifi, filter_smoke_free, filter_guard).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, security guard, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "smoke-free") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND smoke_free =? AND pet=?", filter_parking, filter_wifi, filter_smoke_free, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, smoke-free, pet friendly, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "wifi") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND wifi=? AND security_guard =? AND pet=?", filter_parking, filter_wifi, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, wifi, security guard, pet friendly, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "parking") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("parking =? AND smoke_free=? AND security_guard =? AND pet=?", filter_parking, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with parking, smoke-free, security guard, pet friendly, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				} else if strings.EqualFold(*body.Facilities[0], "wifi") && strings.EqualFold(*body.Facilities[1], "smoke-free") && strings.EqualFold(*body.Facilities[2], "security guard") && strings.EqualFold(*body.Facilities[3], "pet friendly") && strings.EqualFold(*body.Facilities[4], "air-conditioner") && strings.EqualFold(*body.Facilities[5], "fan") {
					if result := mysql.Gorm.Preload("Rooms", "airc=? OR fan=?", filter_air, filter_fan).Preload("Reviews").Where("wifi =? AND smoke_free=? AND security_guard =? AND pet=?", filter_wifi, filter_smoke_free, filter_guard, filter_pet).Find(&dorms); result.Error != nil {
						return &response.GenericError{
							Message: "Not find dorm with wifi, smoke-free, security guard, pet friendly, air-conditioner ,fan",
							Err:     result.Error,
						}
					}
				}

			} else if len(body.Facilities) == 7 {
				//1
				if result := mysql.Gorm.Preload("Rooms", "airc = ? OR fan =?", filter_air, filter_fan).Preload("Reviews").Where("parking = ?", filter_parking).Where("wifi = ?", filter_wifi).Where("smoke_free = ?", filter_smoke_free).Where("security_guard = ?", filter_guard).Where("pet = ?", filter_pet).Find(&dorms); result.Error != nil {
					return &response.GenericError{
						Message: "Dorm not found",
						Err:     result.Error,
					}
				}

			}

		} else {
			// no fac+ distant
			if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Find(&dorms); result.Error != nil {
				return &response.GenericError{
					Message: "Dorms not found",
					Err:     result.Error,
				}
			}
		}

	}

	data, _ := value.Iterate(dorms, func(dorm model.Dorm) (*payload.DormSearch, error) {
		//price
		var prices []float64
		for _, room := range dorm.Rooms {
			prices = append(prices, *room.Price)
		}
		sort.Float64sAreSorted(prices)
		//rate
		var overallRates []float64
		for _, rate := range dorm.Reviews {
			overallRates = append(overallRates, *rate.RatingOverall)
		}
		var sum float64
		var finalRate float64
		if len(overallRates) > 0 {
			for i := 0; i < len(overallRates); i++ {
				sum = sum + overallRates[i]
			}
			finalRate = sum / float64(len(overallRates))
		} else {
			finalRate = 0
		}
		//coverimage
		coverImage, _ := url.JoinPath(config.C.URL, *dorm.CoverImage)

		return &payload.DormSearch{
			DormId:      dorm.Id,
			DormName:    dorm.DormName,
			CoverImage:  &coverImage,
			MinPrice:    &prices[0],
			MaxPrice:    &prices[len(prices)-1],
			OverallRate: &finalRate,
		}, nil
	})

	return c.JSON(response.NewResponse(data))
	// return c.JSON(dorms)
}
