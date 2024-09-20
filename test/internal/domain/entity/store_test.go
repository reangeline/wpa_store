package entity_test

import (
	"testing"

	"github.com/reangeline/wpa_store/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {

	// BDD

	// Given
	// When
	// Then

	t.Run("Given params When create a store Then return store entity", func(t *testing.T) {
		// AAA
		// Arrange
		name := "Store 1"
		phone := "08123456789"
		// Act
		store, err := entity.NewStore(name, phone)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, name, store.Name)
		assert.Equal(t, phone, store.PhoneNumber)
		assert.NotEmpty(t, store.IDStore)

	})
}
