package domain

type RepoDetails struct {
	RepositoryName string `json:"repository_name"`
	RepositoryURL  string `json:"repository_url"`
}

type Response struct {
	Items []RepoDetails `json:"items"`
}

type Statistics struct {
	Query string `json"query"`
	Count string `json:"count"`
}

type StatResponse struct {
	Items []Statistics `json:"items"`
}

type Service interface {
	Fetch(string) (*Response, error)
	Stat() (*StatResponse, error)
}

type Downstream interface {
	Get(string) (*Response, error)
}

type DB interface {
	GetStat() (*StatResponse, error)
	UpdateStat(string) error
}
