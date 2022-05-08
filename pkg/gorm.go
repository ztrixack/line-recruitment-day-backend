package pkg

import (
	"fmt"

	"election-service/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		configs.Database.Host,
		configs.Database.Username,
		configs.Database.Password,
		configs.Database.Database,
		configs.Database.Port,
		configs.Database.SSLMode,
		configs.App.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Exit(err)
	}

	Info("OO", "Database connected!")
	return db
}
