package services

import (
	"golang-pinjol/dto"
	"golang-pinjol/model"
	"golang-pinjol/repository"
	"log"

	"github.com/mashingan/smapping"
)

type PekerjaanNasabahService interface {
	AddNasabahJobsService(jobs dto.CreatePekerjaanNasabahDTO) (*model.Master_Jobs_Nasabah, error)
	NasabahJobsUpdateService(jobs dto.UpdatePekerjaanNasabahDTO) (*model.Master_Jobs_Nasabah, error)
	SearchNasabahJobsByIdService(id int) (*model.Master_Jobs_Nasabah, error)
	DeleteNasabahJobsService(id int) error
}

type pekerjaanNasabahService struct {
	jobsRepository repository.RepositoryNasabahJobs
}

func NewPekerjaanNasabahService(jp repository.RepositoryNasabahJobs) PekerjaanNasabahService {
	return &pekerjaanNasabahService{
		jobsRepository: jp,
	}
}

func (pns *pekerjaanNasabahService) AddNasabahJobsService(jobs dto.CreatePekerjaanNasabahDTO) (*model.Master_Jobs_Nasabah, error) {
	var jobsNasabah model.Master_Jobs_Nasabah
	err := smapping.FillStruct(&jobsNasabah, smapping.MapFields(&jobs))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	addJobs, err := pns.jobsRepository.AddNasabahJobs(&jobsNasabah)
	if err != nil {
		log.Printf("error add nasabah %v", err)
	}

	return addJobs, nil
}

func (pns *pekerjaanNasabahService) NasabahJobsUpdateService(jobs dto.UpdatePekerjaanNasabahDTO) (*model.Master_Jobs_Nasabah, error) {
	var jobsUpdate model.Master_Jobs_Nasabah
	err := smapping.FillStruct(&jobsUpdate, smapping.MapFields(&jobs))
	if err != nil {
		log.Printf("error map %v", err)
	}

	updateJobs, err := pns.jobsRepository.NasabahJobsUpdate(jobs.Id, &jobsUpdate)
	if err != nil {
		log.Printf("failed to update nasabah jobs %v", err)
	}

	return updateJobs, nil
}

func (pns *pekerjaanNasabahService) SearchNasabahJobsByIdService(id int) (*model.Master_Jobs_Nasabah, error) {
	return pns.jobsRepository.SearchNasabahJobsById(id)
}

func (pns *pekerjaanNasabahService) DeleteNasabahJobsService(id int) error {
	return pns.jobsRepository.DeleteNasabahJobs(id)
}
