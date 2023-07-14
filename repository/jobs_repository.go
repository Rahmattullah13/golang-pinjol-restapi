package repository

import (
	"golang-pinjol/model"

	"gorm.io/gorm"
)

type CustomerJobsRepository interface {
	AddCustomerJobs(jobs *model.Master_Jobs_Customer) (*model.Master_Jobs_Customer, error)
	UpdateCustomerJobs(id int, jobs *model.Master_Jobs_Customer) (*model.Master_Jobs_Customer, error)
	SearchCustomerJobsById(id int) (*model.Master_Jobs_Customer, error)
	DeleteCustomerJobs(id int) error
}

type customerJobsRepository struct {
	db *gorm.DB
}

func NewCustomerJobsRepository(db *gorm.DB) CustomerJobsRepository {
	return &customerJobsRepository{
		db: db,
	}
}

func (db *customerJobsRepository) AddCustomerJobs(jobs *model.Master_Jobs_Customer) (*model.Master_Jobs_Customer, error) {
	if err := db.db.Create(jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (db *customerJobsRepository) UpdateCustomerJobs(id int, jobs *model.Master_Jobs_Customer) (*model.Master_Jobs_Customer, error) {
	if err := db.db.Model(&model.Master_Jobs_Customer{ID: id}).Updates(jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (db *customerJobsRepository) SearchCustomerJobsById(id int) (*model.Master_Jobs_Customer, error) {
	var jobs model.Master_Jobs_Customer
	if err := db.db.Find(&jobs, id).Error; err != nil {
		return nil, err
	}
	return &jobs, nil
}

func (db *customerJobsRepository) DeleteCustomerJobs(id int) error {
	if err := db.db.Where("id = $1", id).Delete(&model.Master_Jobs_Customer{}).Error; err != nil {
		return err
	}
	return nil
}
