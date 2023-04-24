package Database

import (
	"where_my_pet_at/server/Models"

	"github.com/jinzhu/gorm"
)

// MigrateDatabaseSchema migrates the database schema
func MigrateDatabaseSchema(db *gorm.DB) {
	// Migrate database schema
	db.AutoMigrate(&Models.Pet{})
	db.AutoMigrate(&Models.User{})
	db.AutoMigrate(&Models.Characteristic{})
	db.AutoMigrate(&Models.Image{})
	db.AutoMigrate(&Models.Location{})
	db.AutoMigrate(&Models.UserPet{})
	db.AutoMigrate(&Models.PetLocation{})
	db.AutoMigrate(&Models.PetImage{})
	db.AutoMigrate(&Models.PetLocation{})
	db.AutoMigrate(&Models.PetCharacteristic{})

	//create data only if database is empty
	users := []Models.User{
		{Name: "admin", Password: "admin", Email: "admin@localhost"},
		{Name: "teste", Password: "teste", Email: "teste@localhost"},
		{Name: "mais um", Password: "mais um", Email: "maisum@localhost"},
	}

	var usersCount int
	db.Model(&Models.User{}).Count(&usersCount)
	if usersCount <= 0 {

		for _, user := range users {
			db.Create(&user)
		}

	}

	var petsCount int
	db.Model(&Models.User{}).Count(&petsCount)
	if petsCount <= 0 {
		pets := []Models.Pet{
			{Name: "Fido", Description: "Cachorro perdido", Age: "2", LostDate: "2019-01-01", FoundDate: "", Image_url: "https://www.petz.com.br/blog/wp-content/uploads/2019/05/cachorro-independente-1.jpg", Owner: users[0]},
			{Name: "Gatinho", Description: "Gatinho perdido", Age: "1", LostDate: "2022-03-10", FoundDate: "", Image_url: "https://www.petz.com.br/blog/wp-content/uploads/2021/11/enxoval-para-gato-Copia.jpg", Owner: users[0]},
			{Name: "Passarinho", Description: "Passarinho perdido", Age: "5", LostDate: "2023-03-10", FoundDate: "", Image_url: "https://i.pinimg.com/736x/95/c2/c9/95c2c9c00992a7d283d02818e470ad56--tim-beta-goldfinch.jpg", Owner: users[1]},
			{Name: "Jaguatirica", Description: "Jaguatirica perdida", Age: "3", LostDate: "", FoundDate: "2021-08-17", Image_url: "https://procarnivoros.org.br/wp-content/uploads/2020/06/jaguatirica-Leopardus-pardalis-adriano-gambarini.jpg", Owner: users[1]},
			{Name: "Tartaruga", Description: "Tartaruga perdida", Age: "24", LostDate: "2020-05-28", FoundDate: "", Image_url: "https://img.olhardigital.com.br/wp-content/uploads/2022/06/jonathan-tartaruga.jpg", Owner: users[2]},
		}

		for _, pet := range pets {
			db.Create(&pet)
		}
	}
}
