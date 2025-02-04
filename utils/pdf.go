package utils

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"premium_microservice/models"
)

func GenerateReceipt(transaction models.Transaction) string {
	filePath := fmt.Sprintf("receipts/receipt_%d.pdf", transaction.ID)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Fiscal check")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Trunsaction id: %d", transaction.ID))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("total amount: %.2f", transaction.TotalPrice))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Status: %s", transaction.Status))
	pdf.Ln(10)
	pdf.Cell(40, 10, "Thank you for your purchase!")
	pdf.Ln(10)
	pdf.Cell(40, 10, "Enjoy your premium subscription with MyEcho.")

	pdf.OutputFileAndClose(filePath)
	return filePath
}
