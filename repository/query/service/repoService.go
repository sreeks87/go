package service

import (
	"github.com/sreeks87/repository/query/domain"
)

type RepoSvc struct {
	Database domain.DB
	DS       domain.Downstream
}

func NewRepoSvc(db domain.DB, ds domain.Downstream) domain.Service {
	return &RepoSvc{
		Database: db,
		DS:       ds,
	}
}

func (svc *RepoSvc) Fetch(param string) (*domain.Response, error) {

	return nil, nil
}

func (svc *RepoSvc) Stat() (*domain.StatResponse, error) {
	return svc.Database.GetStat()
}

func (svc *RepoSvc) Call(param string) {
	svc.DS.Get(param)
}
