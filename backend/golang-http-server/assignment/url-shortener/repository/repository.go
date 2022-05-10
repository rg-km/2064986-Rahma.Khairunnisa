package repository

import (
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {
	if r.Data[path] != "" {
		return &entity.URL{}, errors.New("url not found")
	}
	return &entity.URL{} , nil // TODO: replace this
}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	r.mu.Lock()
	rest := entity.GetRandomShortURL(longURL)
	defer r.mu.Unlock()
	return &entity.URL{
		LongURL: rest,
	}, nil
	//&entity.URL{} , nil // TODO: replace this
}

//func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	//&entity.URL{} , nil // TODO: replace this
//}
