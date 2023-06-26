package dto

type RegisterNasabahDTO struct{
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	NoKtp string `json:"no_ktp" form:"no_ktp" binding:"required"`
}

type LoginNasabahDTO struct{
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateNasabahDTO struct{
	Id uint64 `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	NoKtp string `json:"no_ktp" form:"no_ktp" binding:"required"`
}