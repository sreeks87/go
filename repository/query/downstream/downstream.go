package downstream

import "github.com/sreeks87/repository/query/domain"

type BeChallenge struct {
	URL string
}

func NewBeChallenge(u string) domain.Downstream {
	return &BeChallenge{
		URL: u,
	}
}

func (b *BeChallenge) Get(query string) (*domain.Response, error) {
	return nil, nil
}
