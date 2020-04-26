package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFarm(t *testing.T) {
	t.Run("can create a farm", func(t *testing.T) {
		farm := Farm{
			ID:   "id",
			Name: "Animal Farm",
		}

		assert.Equal(t, "id", farm.ID)
		assert.Equal(t, "Animal Farm", farm.Name)
	})
}

func TestAnimal(t *testing.T) {
	t.Run("can create an animal", func(t *testing.T) {
		animal := Animal{
			ID:       "id",
			Type:     "Pig",
			Location: "here",
			FarmID:   "1",
		}

		assert.Equal(t, "id", animal.ID)
		assert.Equal(t, "Pig", animal.Type)
		assert.Equal(t, "here", animal.Location)
		assert.Equal(t, "1", animal.FarmID)
	})
}
