package dao

import (
	"github.com/nicholaskoldys/currency-tracker/models/currency"
)

/* *
* CurrencyDecimalDAO methods are not database methods but repository methods that will hold logical paths to appropriate data:
* ie. Get methods will attempt to locate repo but can call repo get methods for update.
 */
type CurrencyDecimalDAO interface {
	GetDecimals(queryString string) []*currency.CurrencyDecimal
	GetDecimal(queryString string, typee string, code string) currency.CurrencyDecimal
	PostDecimal(queryString string)
	UpdateDecimal(queryString string)
}

// func GetCurrency(type string, code string) {
// 	/* var deciSlice []*currency.CurrencyDecimal

// 	db.GetCurrencyDecimals(getAllDecimals, &deciSlice)

// 	fmt.Println("CurrencyDecimal: ", deciSlice)
// 	for i, curr := range deciSlice {
// 		fmt.Printf("%d : %v \n", i, curr)
// 	}

// 	return deciSlice */
// }

// func GetCurrencyDecimals() []*currency.CurrencyDecimal {

// 	/* var deciSlice []*currency.CurrencyDecimal

// 	db.GetCurrencyDecimals(getAllDecimals, &deciSlice)

// 	fmt.Println("CurrencyDecimal: ", deciSlice)
// 	for i, curr := range deciSlice {
// 		fmt.Printf("%d : %v \n", i, curr)
// 	}

// 	return deciSlice */
// }