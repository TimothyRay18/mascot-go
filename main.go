package main

import (
	"log"
	"mascot/controller"
	"mascot/models"
	"mascot/storage"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	customerRepo := &controller.CustomerRepository{DB: db}
	userRepo := &controller.UserRepository{DB: db}

	api := app.Group("/mascot")
	api.Get("/get_all_customers", customerRepo.GetAllCustomers)
	api.Post("/add_new_customer", customerRepo.AddNewCustomer)
	api.Get("/get_customer_by_id/:id", customerRepo.GetCustomerById)
	api.Put("/update_customer/:id", customerRepo.UpdateCustomer)
	api.Delete("/delete_customer/:id", customerRepo.DeleteCustomer)

	api.Post("/add_new_user", userRepo.AddNewUser)
	api.Post("/login", userRepo.Login)
	api.Get("/logout", userRepo.Logout)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateCustomerData(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	err = models.MigrateCustomer(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	err = models.MigrateTransaction(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	err = models.MigrateUserApp(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	app := fiber.New()
	SetupRoutes(app, db)
	app.Listen(":8080")
}
