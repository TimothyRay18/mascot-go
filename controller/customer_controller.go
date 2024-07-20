package controller

import (
	"fmt"
	"mascot/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func (r *CustomerRepository) GetAllCustomers(context *fiber.Ctx) error {
	customers := &[]models.Customer{}

	err := r.DB.Find(customers).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "customers fetched successfully",
			"data":    customers,
		})
	return nil
}

func (r *CustomerRepository) AddNewCustomer(context *fiber.Ctx) error {
	customer := models.Customer{}

	err := context.BodyParser(&customer)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&customer).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not add customer"})
		return err
	}

	data := map[string]interface{}{
		"name":  customer.Name,
		"cis":   customer.Cis,
		"bcaId": customer.BcaId}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "customer has been added",
		"data":    data})
	return nil
}

func (r *CustomerRepository) GetCustomerById(context *fiber.Ctx) error {
	id := context.Params("id")
	customer := &models.Customer{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(customer).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "id not found in database",
		})
		fmt.Println("ASDFA")
		return nil
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "customer id fetched successfully",
		"data":    customer,
	})
	return nil
}

func (r *CustomerRepository) UpdateCustomer(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	fmt.Println("the ID is ", id)

	updateCustomer := models.Customer{}
	if err := context.BodyParser(&updateCustomer); err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "request failed",
		})
		return err
	}

	customer := &models.Customer{}
	if err := r.DB.Where("id = ?", id).First(customer).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "id not found in database",
		})
		return nil
	}

	if updateCustomer.BcaId != nil {
		customer.BcaId = updateCustomer.BcaId
	}
	if updateCustomer.Cis != nil {
		customer.Cis = updateCustomer.Cis
	}
	if updateCustomer.Name != nil {
		customer.Name = updateCustomer.Name
	}

	if err := r.DB.Save(customer).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "failed to save customer to db",
		})
		return nil
	}

	data := map[string]interface{}{
		"name":  customer.Name,
		"cis":   customer.Cis,
		"bcaId": customer.BcaId}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message":  "Customer " + id + " updated successfully",
		"customer": data,
	})
	return nil
}

func (r *CustomerRepository) DeleteCustomer(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	fmt.Println("the ID to delete is ", id)

	customer := &models.Customer{}
	err := r.DB.Delete(customer, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete customer",
		})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "customer delete successfully",
	})
	return nil
}
