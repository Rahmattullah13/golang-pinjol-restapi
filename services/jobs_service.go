package services

import (
	"golang-pinjol/dto"
	"golang-pinjol/model"
	"golang-pinjol/repository"
	"log"

	"github.com/mashingan/smapping"
)

type JobsCustomerService interface {
	AddCustomerJobsService(jobs dto.CreateJobsCustomerDTO) (*model.Master_Jobs_Customer, error)
	UpdateCustomerJobsService(jobs dto.UpdateJobsCustomerDTO) (*model.Master_Jobs_Customer, error)
	SearchCustomerJobsByIdService(id int) (*model.Master_Jobs_Customer, error)
	DeleteCustomerJobsService(id int) error
}

type jobsCustomerService struct {
	jobsRepository repository.CustomerJobsRepository
}

func NewJobsCustomerService(jp repository.CustomerJobsRepository) JobsCustomerService {
	return &jobsCustomerService{
		jobsRepository: jp,
	}
}

func (pns *jobsCustomerService) AddCustomerJobsService(jobs dto.CreateJobsCustomerDTO) (*model.Master_Jobs_Customer, error) {
	var jobsCustomer model.Master_Jobs_Customer
	err := smapping.FillStruct(&jobsCustomer, smapping.MapFields(&jobs))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	addJobs, err := pns.jobsRepository.AddCustomerJobs(&jobsCustomer)
	if err != nil {
		log.Printf("error add customer %v", err)
	}

	return addJobs, nil
}

func (pns *jobsCustomerService) UpdateCustomerJobsService(jobs dto.UpdateJobsCustomerDTO) (*model.Master_Jobs_Customer, error) {
	var jobsUpdate model.Master_Jobs_Customer
	err := smapping.FillStruct(&jobsUpdate, smapping.MapFields(&jobs))
	if err != nil {
		log.Printf("error map %v", err)
	}

	updateJobs, err := pns.jobsRepository.UpdateCustomerJobs(jobs.Id, &jobsUpdate)
	if err != nil {
		log.Printf("failed to update customer jobs %v", err)
	}

	return updateJobs, nil
}

func (pns *jobsCustomerService) SearchCustomerJobsByIdService(id int) (*model.Master_Jobs_Customer, error) {
	return pns.jobsRepository.SearchCustomerJobsById(id)
}

func (pns *jobsCustomerService) DeleteCustomerJobsService(id int) error {
	return pns.jobsRepository.DeleteCustomerJobs(id)
}
