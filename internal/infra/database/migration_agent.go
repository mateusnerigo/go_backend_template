package database

func MigrationAgent() {
	db, _ := Client()
	
	db.AutoMigrate(
		&models.User{},
	)
}