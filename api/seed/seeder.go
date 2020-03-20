package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/junlapong/go-rest-api/api/model"
)

var users = []model.User{
	model.User{
		Nickname: "Steven Victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	model.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []model.Post{
	model.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	model.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	//err := db.Debug().DropTableIfExists(&model.Post{}, &model.User{}).Error
	err := db.Debug().DropTableIfExists(&model.User{}).Error

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	//err = db.Debug().AutoMigrate(&model.User{}, &model.Post{}).Error
	err = db.Debug().AutoMigrate(&model.User{}).Error

	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// err = db.Debug().Model(&model.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	//for i, _ := range users {
	for i := range users {
		err = db.Debug().Model(&model.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		// posts[i].AuthorID = users[i].ID

		// err = db.Debug().Model(&model.Post{}).Create(&posts[i]).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
	}
}
