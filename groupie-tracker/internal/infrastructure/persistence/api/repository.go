package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"groupie-tracker/internal/domain"
	"io"
	"net/http"
)

type Repository struct {
	apiDomain string
}

func New(apiDomain string) *Repository {
	return &Repository{
		apiDomain: apiDomain,
	}
}

func (r Repository) ParseJSONFromURL(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.Join(domain.ErrFetchingUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Join(domain.ErrRedingBody, err)
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.Join(domain.ErrUnmarshJson, err)
	}
	return nil
}
