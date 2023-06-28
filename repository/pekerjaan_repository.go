package repository

import (
	"golang-pinjol/model"

	"gorm.io/gorm"
)

type RepositoryNasabahJobs interface {
	AddNasabahJobs(jobs *model.Master_Jobs_Nasabah) (*model.Master_Jobs_Nasabah, error)
	NasabahJobsUpdate(id int, jobs *model.Master_Jobs_Nasabah) (*model.Master_Jobs_Nasabah, error)
	SearchNasabahJobsById(id int) (*model.Master_Jobs_Nasabah, error)
	DeleteNasabahJobs(id int) error
}

type connectionNasabahJob struct {
	db *gorm.DB
}

func NewRepositoryNasabahJobs(db *gorm.DB) RepositoryNasabahJobs {
	return &connectionNasabahJob{
		db: db,
	}
}

func (db *connectionNasabahJob) AddNasabahJobs(jobs *model.Master_Jobs_Nasabah) (*model.Master_Jobs_Nasabah, error) {
	if err := db.db.Create(jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (db *connectionNasabahJob) NasabahJobsUpdate(id int, jobs *model.Master_Jobs_Nasabah) (*model.Master_Jobs_Nasabah, error) {
	if err := db.db.Model(&model.Master_Jobs_Nasabah{ID: id}).Updates(jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (db *connectionNasabahJob) SearchNasabahJobsById(id int) (*model.Master_Jobs_Nasabah, error) {
	var jobs model.Master_Jobs_Nasabah
	if err := db.db.First(&jobs, id).Error; err != nil {
		return nil, err
	}
	return &jobs, nil
}

func (db *connectionNasabahJob) DeleteNasabahJobs(id int) error {
	if err := db.db.Where("id = $1", id).Delete(&model.Master_Jobs_Nasabah{}).Error; err != nil {
		return err
	}
	return nil
}
