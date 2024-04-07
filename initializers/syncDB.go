package initializers

import "goapi/models"

func SyncDB() {
	DB.AutoMigrate(&models.Useracl{})
}
