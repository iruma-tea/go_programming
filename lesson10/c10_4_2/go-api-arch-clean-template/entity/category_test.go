package entity_test

import (
	"go-api-arch-clean-template/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory(t *testing.T) {
	category := entity.Category{
		ID:   1,
		Name: "sports",
	}

	assert.Equal(t, 1, category.ID)
	assert.Equal(t, "sports", string(category.Name))
}
