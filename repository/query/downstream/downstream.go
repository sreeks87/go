package downstream

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/sreeks87/repository/query/domain"
)

type BeChallenge struct {
	URL string
}

func NewBeChallenge(u string) domain.Downstream {
	return &BeChallenge{
		URL: u,
	}
}

func (b *BeChallenge) Get(query string) (*domain.Response, error) {
	finalurl := b.URL + "?query=" + query
	resp, err := http.Get(finalurl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("could not get success from downstream")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var r domain.Response
	json.Unmarshal(body, &r)
	return &r, nil
}
