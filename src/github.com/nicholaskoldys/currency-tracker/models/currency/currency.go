package currency

import (
	"math"
)

/* type CurrencyStruct uint8

const (
	CurrencyDecimal CurrencyStruct = iota
	CurrencyRate
	Currency
)

func (cs CurrencyStruct) String()  {
    switch cs {
    case CurrencyDecimal:
        return "CurrencyDecimal"
    case CurrencyRate:
        return "CurrencyRate"
    case Currency:
        return "Currency"
	}
	return "unknown"
} */

/* type CurrencyStruct interface {
    CurrencyDecimal | CurrencyRate | Currency
} */

type Currency struct {
	CurrencyDecimal *CurrencyDecimal //*CurrencyDecimal
	CurrencyRate    *CurrencyRate //*CurrencyRate
	CodeBase        string //*CurrnecyRate.CodeBase
	CodeQuote       string //*CurrnecyRate.CodeQuote
	MiddleRate      float64 //*CurrnecyRate.MiddleRate
	MonthRate       float64 //*CurrnecyRate.MonthRate
	YearAvg         float64 //*CurrnecyRate.YearAvg
	Trend           float64 //*CurrnecyRate.Trend
	DateOf          string //*CurrnecyRate.DateOf
}

type CurrencyDao interface {
    GetCurrency(id int) (*Currency, error)
    CreateCurrency(cur *Currency) error
    UpdateCurrency(cur *Currency) error
    DeleteCurrency(id int) error
}

/* func (cs *CurrencyStruct) PrintFields ( ) {
	fmt.Println(cs);
} */

/* func (s ) ToArray() []interface{} {
    // use reflection to iterate over the struct fields
    sValue := reflect.ValueOf(s).Elem()
    a := make([]interface{}, sValue.NumField())
    for i := 0; i < sValue.NumField(); i++ {
        field := sValue.Field(i)
        // check if the field is exported
        if field.CanInterface() {
            a[i] = field.Interface()
        } else {
            // if the field is not exported, set it to nil
            a[i] = nil
        }
    }
    return a
} */

func CreateCurrency(Decimal *CurrencyDecimal, Rate *CurrencyRate) *Currency {
	cur := Currency{
		CurrencyDecimal: Decimal,
		CurrencyRate:    Rate, 
		CodeBase:        Rate.CodeBase, 
		CodeQuote:       Rate.CodeQuote, 
		MiddleRate:      float64(Rate.MiddleRate) / math.Pow(10.0, float64(Decimal.DecimalShift)), 
		MonthRate:       float64(Rate.MonthRate) / math.Pow(10.0, float64(Decimal.DecimalShift)), 
		YearAvg:         float64(Rate.YearAvg) / math.Pow(10.0, float64(Decimal.DecimalShift)), 
		Trend:           Rate.Trend,  
		DateOf:          Rate.DateOf,
	}
	return &cur
}

