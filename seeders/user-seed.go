package seeders

import (
	"log"
	"techincal-test/database"
	"techincal-test/structs"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	users := []structs.User{
		{Email: "user1@mail.com", Password: "user1"},
		{Email: "user2@mail.com", Password: "user2"},
	}

	for _, user := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			panic(err)
		}
		user.Password = string(hashedPassword)

		result := database.DB.FirstOrCreate(&user)
		if result.Error != nil {
			log.Printf("Failed to seed user %s: %v", user.Email, result.Error)
		} else {
			log.Printf("User %s seeded successfully", user.Email)
		}
	}
}
