package seeder

import (
	entity "fp-mbd-amidrive/domain/model/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Just this function is Public
func UserSeeder(db *gorm.DB) error {
	var err error

	// Data Preprocessing
	users := [10]entity.User{}

	for _, dt := range users {
		err = addUser(db, dt)
		if err != nil {
			return err
		}
	}

	return nil
}

// Other function shouldn't be Public
func addUser(db *gorm.DB, user entity.User) error {
	user.ID = uuid.New().String()
	uc := db.Create(&user)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}
