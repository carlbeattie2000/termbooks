package invoicedomain

import (
	financialdomain "tax_calculator/engine/internal/domains/book_keeping_domain/financial_domain"
	itemsdomain "tax_calculator/engine/internal/domains/book_keeping_domain/items_domain"
	salesvalueobjects "tax_calculator/engine/internal/valueobjects/sales_valueobjects"
	"time"
)

type InvoiceItemEntry struct {
	item        itemsdomain.Item
	description string
	quantity    int
	rate        float32
	taxRate     *financialdomain.TaxRate
	discount    float32
	total       float32
}

type Invoice struct {
	customerName      string
	invoiceDate       time.Time
	dueDate           time.Time
	invoiceNumber     int
	reference         int
	items             []*InvoiceItemEntry
	message           string
	termsAndConditons string
	paymentOption     salesvalueobjects.PaymentOptions
	subtotal          float32
	discount          float32
	adjustment        float32
	total             float32
	paymentAmount     float32
	status            salesvalueobjects.InvoiceStatus
	void              bool
}
