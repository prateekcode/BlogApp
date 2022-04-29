package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/prateekcode/blogapp/api/models"
)

var users = []models.User{
	{
		Nickname: "Prateek Rai",
		Email:    "pratiekray@gmail.com",
		Password: "password",
	},
	{
		Nickname: "Prateek Code",
		Email:    "codes.prateek@gmail.com",
		Password: "password123",
	},
}

var posts = []models.Post{
	{
		Title:   "Test title 1",
		Content: "Hello World test",
	},
	{
		Title:   "Test title 2",
		Content: "Hey hello",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attatching foreign key error: %v", err)
	}
	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model((&models.Post{})).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed the posts table: %v", err)
		}
	}
}
