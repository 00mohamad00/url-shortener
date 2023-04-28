package urlshortener

type Service interface {
	GetUrl(token string) (string, error)
	AddUrl(url string) (string, error)
}
