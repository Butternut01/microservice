package main

import (
	"fmt"
	"log"
	"net/http"
	"premium_microservice/database"
	"premium_microservice/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	router := gin.Default()

	router.POST("/payment", CreateTransaction)
	router.POST("/process-payment/:id", ProcessPayment)

	log.Println("Микросервис запущен на порту 8081")
	router.Run(":8081")
}

// Создание транзакции
func CreateTransaction(c *gin.Context) {
	var request struct {
		CartItems []struct {
			ID    string  `json:"id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		} `json:"cartItems"`
		Customer struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"customer"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(request)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	totalPrice := 0.0
	for _, item := range request.CartItems {
		totalPrice += item.Price
	}

	transaction := models.Transaction{
		CartID:     request.Customer.ID,
		CustomerID: request.Customer.ID,
		Email:      request.Customer.Email,
		Status:     "pending",
		TotalPrice: totalPrice,
	}

	database.DB.Create(&transaction)
	c.JSON(http.StatusOK, gin.H{"transaction_id": transaction.ID, "status": transaction.Status})
}

// Обработка платежа
func ProcessPayment(c *gin.Context) {
	var payment struct {
		CardNumber     string `json:"cardNumber"`
		ExpirationDate string `json:"expirationDate"`
		CVV            string `json:"cvv"`
		Name           string `json:"name"`
		Address        string `json:"address"`
	}

	id := c.Param("id")
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment data"})
		return
	}

	var transaction models.Transaction
	if err := database.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if payment.CardNumber == "0000 0000 0000 0000" {
		transaction.Status = "declined"
	} else {
		transaction.Status = "paid"
	}

	database.DB.Save(&transaction)
	c.JSON(http.StatusOK, gin.H{"transaction_id": transaction.ID, "status": transaction.Status})
}
