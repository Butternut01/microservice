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
	pdf.Cell(40, 10, "Фискальный чек")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Номер транзакции: %d", transaction.ID))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Общая сумма: %.2f", transaction.TotalPrice))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Статус: %s", transaction.Status))

	pdf.OutputFileAndClose(filePath)
	return filePath
}
