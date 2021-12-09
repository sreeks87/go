package service

import (
	"errors"

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
	if e := svc.Database.UpdateStat(param); e != nil {
		return nil, errors.New("update failed")
	}
	return svc.DS.Get(param)
}

func (svc *RepoSvc) Stat() (*domain.StatResponse, error) {
	return svc.Database.GetStat()
}
