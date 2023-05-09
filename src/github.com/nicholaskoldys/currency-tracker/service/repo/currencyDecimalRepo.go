package repo

import (
	"fmt"

	"github.com/nicholaskoldys/currency-tracker/models/currency"
	"github.com/nicholaskoldys/currency-tracker/service/db"
	"github.com/nicholaskoldys/currency-tracker/util"
)

type DecimalRepository interface {
	GetDecimals() []*currency.CurrencyDecimal
	GetDecimal(typee string, code string) currency.CurrencyDecimal
	PostDecimal()
	UpdateDecimal()
}

const getAllDecimals string = "SELECT * FROM currencyDecimals"
const getCurrencyDecimal string = "SELECT * FROM currencyDecimals WHERE type = ? AND code = ?"
const postCurrencyDecimal string = ""
const updateCurrencyDecimal string = ""

var iDecimals *[]*currency.CurrencyDecimal

func init() {
	renewRepo()
}

func renewRepo() {
	dbCurrentDecimals := GetDecimals()
	iDecimals = &dbCurrentDecimals
	fmt.Printf("%v \n", iDecimals)
}

func getDecimals() *[]*currency.CurrencyDecimal {
	if(len(*iDecimals) == 0) {
		renewRepo()
	}
	return iDecimals
}

func GetDecimals() []*currency.CurrencyDecimal {

	db := db.GetDB()
	if db == nil {
		util.Debug(9, "Database returned nothing")
		panic("no database")
	}

	stmt, err := db.Prepare(getAllDecimals)
	if err != nil {
		util.Debug(9, err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		util.Debug(9, err.Error())
	}
	
	structSlice := []*currency.CurrencyDecimal{}

	for rows.Next() {
		cD := currency.CurrencyDecimal{}
		if err := rows.Scan(cD.GetStructFields()...); err != nil {
			util.Debug(9, err.Error())
		}

		structSlice = append(structSlice, &cD)
	}

	return structSlice
}

func getDecimal(typee string, code string) (*currency.CurrencyDecimal, bool) {
	if(len(*iDecimals) == 0) {
		renewRepo()	
	}

	for _, dec := range *iDecimals {
		if dec.Type == typee && dec.Code == code {
			return dec, true
		}
	}

	if gD, isOk := GetDecimal(typee, code); isOk != false {
		return gD, true
	}

	return nil, false
}

func GetDecimal(typee string, code string) (*currency.CurrencyDecimal, bool) {
	db := db.GetDB()
	if db == nil {
		util.Debug(9, "Database returned nothing")
		panic("no database")
	}

	stmt, err := db.Prepare(getCurrencyDecimal)
	if err != nil {
		util.Debug(9, err.Error())
	}

	rows, err := stmt.Query(typee, code)
	if err != nil {
		util.Debug(9, err.Error())
	}

	if rows.Next() {
		currencyDecmial := currency.CurrencyDecimal{}
		if err := rows.Scan(currencyDecmial.GetStructFields()...); err != nil {
			util.Debug(9, err.Error())
		}	
		return &currencyDecmial, true
	}
	
	return nil, false
	
}