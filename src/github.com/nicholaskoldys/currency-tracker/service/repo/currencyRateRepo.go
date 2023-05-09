package repo

import (
	"fmt"

	"github.com/nicholaskoldys/currency-tracker/models/currency"
	"github.com/nicholaskoldys/currency-tracker/service/db"
	"github.com/nicholaskoldys/currency-tracker/util"
)

type RateRepository interface {
	GetRates(queryString string) []*currency.CurrencyRate
	GetRate(queryString string) currency.CurrencyRate
	PostRate(queryString string)
	UpdateRate(queryString string)
}

const getAllRates string = "SELECT * FROM currencyRates"

var iRates *[]*currency.CurrencyRate

func init() {
	dbCurrentRates := GetRates(getAllRates)
	iRates = &dbCurrentRates
	fmt.Printf("%v \n", iRates)
}

func GetRates(queryString string) []*currency.CurrencyRate {
	db := db.GetDB()
	if db == nil {
		util.Debug(9, "Database returned nothing")
		panic("no database")
	}

	stmt, err := db.Prepare(queryString)
	if err != nil {
		util.Debug(9, err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		util.Debug(9, err.Error())
	}

	structSlice := []*currency.CurrencyRate{}

	for rows.Next() {
		cD := currency.CurrencyRate{}
		if err:= rows.Scan( cD.GetStructFields()... ); err != nil {
			util.Debug(9, err.Error())
		}

		structSlice = append(structSlice, &cD)
	}

	return structSlice
}

func GetRate(queryString string) currency.CurrencyRate {
	db := db.GetDB()
	if db == nil {
		util.Debug(9, "Database returned nothing")
		panic("no database")
	}

	stmt, err := db.Prepare(queryString)
	if err != nil {
		util.Debug(9, err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		util.Debug(9, err.Error())
	}

	var currencyRate currency.CurrencyRate

	for rows.Next() {
		currencyRate = currency.CurrencyRate{}
		if err:= rows.Scan( currencyRate.GetStructFields()... ); err != nil {
			util.Debug(9, err.Error())
		}
	}

	return currencyRate
}