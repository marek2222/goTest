package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"time"
)

// CurrencyArray model
type CurrencyArray struct {
	CurrencyList []Currency
}

// AddCurrency model
func (c *CurrencyArray) AddCurrency(currency string, amount int) {
	newc := Currency{Amount: amount}
	newc.XMLName.Local = currency
	c.CurrencyList = append(c.CurrencyList, newc)
}

// Currency model
type Currency struct {
	XMLName xml.Name
	Amount  int `xml:",innerxml"`
}

// Plan model
type Plan struct {
	XMLName           xml.Name      `xml:"plan"`
	PlanCode          string        `xml:"plan_code,omitempty"`
	CreatedAt         *time.Time    `xml:"created_at,omitempty"`
	UnitAmountInCents CurrencyArray `xml:"unit_amount_in_cents"`
	SetupFeeInCents   CurrencyArray `xml:"setup_fee_in_cents"`
}

func main() {
	fmt.Println("Hello, playground")
	p := Plan{PlanCode: "test"}
	p.UnitAmountInCents.AddCurrency("USD", 4000)
	p.SetupFeeInCents.AddCurrency("EUR", 2500)
	p.SetupFeeInCents.AddCurrency("EUR", 3200)
	file, _ := xml.MarshalIndent(p, "", "    ")
	file = []byte(xml.Header + string(file))
	_ = ioutil.WriteFile("notes1.xml", file, 0644)

}
