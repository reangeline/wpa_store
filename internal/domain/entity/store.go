package entity

import (
	"errors"

	pkg_entity "github.com/reangeline/wpa_store/pkg/entity"
)

type Store struct {
	IDStore     pkg_entity.ID `json:"id_store"`
	Name        string        `json:"name"`
	PhoneNumber string        `json:"phone_number"`
}

func NewStore(name, phoneNumber string) (*Store, error) {
	store := &Store{
		Name:        name,
		PhoneNumber: phoneNumber,
	}

	store.AddId()

	err := store.IsValid()
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) AddId() {
	s.IDStore = pkg_entity.NewID()
}

func (s *Store) IsValid() error {
	if s.Name == "" {
		return errors.New("store name is required")
	}

	if s.PhoneNumber == "" {
		return errors.New("store phone number is required")
	}

	return nil
}
