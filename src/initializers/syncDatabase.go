package initializers

import (
	"log"
	"sample/src/pkg/entities"
)

func SyncDatabase() {
	if err := DB.Migrator().AutoMigrate(&entities.Permission{}); err != nil {
		log.Fatalf("migration failed for %T: %v", &entities.Permission{}, err)
	}

	if err := DB.Migrator().AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("migration failed for %T: %v", &entities.User{}, err)
	}

}
