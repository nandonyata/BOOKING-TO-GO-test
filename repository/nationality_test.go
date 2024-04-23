package repository_test

import (
	"testing"

	"github.com/nandonyata/BOOKING-TO-GO-test/config"
	"github.com/nandonyata/BOOKING-TO-GO-test/entity"
	"github.com/nandonyata/BOOKING-TO-GO-test/repository"
)

func TestNationalityRepositoryCreate(t *testing.T) {
	config.ConnectDB()
	newRepo := repository.NationalityRepository{Database: config.DB}

	t.Run("Create Nationality", func(t *testing.T) {
		_, err := newRepo.Create(&entity.Nationality{
			Nationality_name: "INDONESIA",
			Nationality_code: "ID",
		})

		if err != nil {
			t.Fatalf("Error Create Nationality : %v\n", err)
		}

		t.Log("Success Create Nationality")
	})
}
