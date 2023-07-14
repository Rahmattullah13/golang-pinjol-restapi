package repository

import (
	"golang-pinjol/model"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *model.Master_Customer) (*model.Master_Customer, error)
	UpdateCustomer(id uint64, customer *model.Master_Customer) (*model.Master_Customer, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	IsDuplicateNIK(noKtp string) (tx *gorm.DB)
	FindByNIK(nik string) model.Master_Customer
	ProfileCustomer(customerID string) model.Master_Customer
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (db *customerRepository) CreateCustomer(customer *model.Master_Customer) (*model.Master_Customer, error) {
	customer.Password = HashPassword([]byte(customer.Password))
	if err := db.db.Create(customer).Error; err != nil {
		panic(err)
	}
	return customer, nil
}

func (db *customerRepository) UpdateCustomer(id uint64, customer *model.Master_Customer) (*model.Master_Customer, error) {
	var existingCustomer model.Master_Customer
	if err := db.db.First(&existingCustomer, id).Error; err != nil {
		return nil, err
	}

	if customer.Password != "" {
		existingCustomer.Password = HashPassword([]byte(customer.Password))
	} else {
		existingCustomer.Password = customer.Password
	}

	existingCustomer.Name = customer.Name
	existingCustomer.Email = customer.Email
	existingCustomer.NoKtp = customer.NoKtp
	existingCustomer.PhoneNumber = customer.PhoneNumber
	existingCustomer.Address = customer.Address

	if err := db.db.Save(&existingCustomer).Error; err != nil {
		return nil, err
	}

	return &existingCustomer, nil
}

func (db *customerRepository) VerifyCredential(email string, password string) interface{} {
	var customer model.Master_Customer
	res := db.db.Where("email = $1", email).Take(&customer)
	if res.Error != nil {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil
		}
		return nil
	}
	return customer
}

func (db *customerRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var customer model.Master_Customer
	return db.db.Where("email = $1", email).Take(&customer)
}

func (db *customerRepository) IsDuplicateNIK(noKtp string) (tx *gorm.DB) {
	var customer model.Master_Customer
	return db.db.Where("noKtp = $1", noKtp).Take(&customer)
}

func (db *customerRepository) FindByNIK(nik string) model.Master_Customer {
	var customer model.Master_Customer
	db.db.Where("no_ktp = $1", nik).Take(&customer)
	return customer
}

func (db *customerRepository) ProfileCustomer(customerID string) model.Master_Customer {
	var customer model.Master_Customer
	db.db.Find(&customer, customerID)
	return customer
}

func HashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}
