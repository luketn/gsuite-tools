package main

import (
	"github.com/mayur-tolexo/mold"
	"path/filepath"
)

//Invoice details
type Invoice struct {
	InvoiceNo   string
	InvoiceDate string
	Currency    string
	AmountDue   float64
	Items       []ItemDetail
	Total       float64
}

//ItemDetail : Item details
type ItemDetail struct {
	Name   string
	Desc   string
	Amount float64
	Qty    int
	Total  float64
}

func CreateInvoice(rootPath string, invoice Invoice) ([]byte, error) {
	mold, _ := mold.NewHTMLTemplate()
	mold.HTMLPath = filepath.Join(rootPath, "invoice.html")
	if err := mold.Execute(invoice); err == nil {
		return mold.Bytes(), nil
	} else {
		return nil, err
	}
}
