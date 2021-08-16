package domain

type ILogger interface {
	Log(string, string) error
}
