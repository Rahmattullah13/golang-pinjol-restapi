package repository

import (
	"golang-pinjol/model"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type NasabahRepository interface {
	CreateNasabah(nasabah *model.Master_Nasabah) (*model.Master_Nasabah, error)
	UpdateNasabah(id uint64, nasabah *model.Master_Nasabah) (*model.Master_Nasabah, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	IsDuplicateNIK(noKtp string) (tx *gorm.DB)
	FindByNIK(nik string) model.Master_Nasabah
	ProfileNasabah(nasabahID string) model.Master_Nasabah
}

type nasabahConnection struct {
	db *gorm.DB
}

func NewNasabahRepository(db *gorm.DB) NasabahRepository {
	return &nasabahConnection{
		db: db,
	}
}

func (db *nasabahConnection) CreateNasabah(nasabah *model.Master_Nasabah) (*model.Master_Nasabah, error) {
	nasabah.Password = HashPassword([]byte(nasabah.Password))
	if err := db.db.Create(nasabah).Error; err != nil {
		panic(err)
	}
	return nasabah, nil
}

func (db *nasabahConnection) UpdateNasabah(id uint64, nasabah *model.Master_Nasabah) (*model.Master_Nasabah, error) {
	var existingNasabah model.Master_Nasabah
	if err := db.db.First(&existingNasabah, id).Error; err != nil {
		return nil, err
	}

	if nasabah.Password != "" {
		existingNasabah.Password = HashPassword([]byte(nasabah.Password))
	} else {
		existingNasabah.Password = nasabah.Password
	}

	existingNasabah.Name = nasabah.Name
	existingNasabah.Email = nasabah.Email
	existingNasabah.NoKtp = nasabah.NoKtp
	existingNasabah.PhoneNumber = nasabah.PhoneNumber
	existingNasabah.Address = nasabah.Address

	if err := db.db.Save(&existingNasabah).Error; err != nil {
		return nil, err
	}

	return &existingNasabah, nil
}

func (db *nasabahConnection) VerifyCredential(email string, password string) interface{} {
	var nasabah model.Master_Nasabah
	res := db.db.Where("email = $1", email).Take(&nasabah)
	if res.Error != nil {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(nasabah.Password), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil
		}
		return nil
	}
	return nasabah
}

func (db *nasabahConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var nasabah model.Master_Nasabah
	return db.db.Where("email = $1", email).Take(&nasabah)
}

func (db *nasabahConnection) IsDuplicateNIK(noKtp string) (tx *gorm.DB) {
	var nasabah model.Master_Nasabah
	return db.db.Where("noKtp = $1", noKtp).Take(&nasabah)
}

func (db *nasabahConnection) FindByNIK(nik string) model.Master_Nasabah {
	var nasabah model.Master_Nasabah
	db.db.Where("no_ktp = $1", nik).Take(&nasabah)
	return nasabah
}

func (db *nasabahConnection) ProfileNasabah(nasabahID string) model.Master_Nasabah {
	var nasabah model.Master_Nasabah
	db.db.Find(&nasabah, nasabahID)
	return nasabah
}

func HashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}
