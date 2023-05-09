package currency

import (
	"fmt"
	"time"
)

/* type DecimalType string

const (
	Rate    DecimalType = "rate"
	Amount  DecimalType = "amount"
	unknown DecimalType = "unknown"
)

func (dt DecimalType) String() string {
	switch dt {
	case Rate:
		return "rate"
	case Amount:
		return "amount"
	}
	return "unknown"
}

func decimalTypeConvert(dt string) DecimalType {
	switch dt {
	case "rate":
		return Rate
	case "amount":
		return Amount
	}
	return unknown
}*/

type CurrencyDecimal struct {
	Id           int    `db:"id"`
	Type         string `db:"type"`
	Code         string `db:"code"`
	DecimalShift int    `db:"decimalShift"`
	InUse        bool   `db:"inUse"`
	CreatedAt    string `db:"createdAt"`
}

func CreateCurrencyDecimal(Type string, Code string, DecimalShift int) *CurrencyDecimal {
	cd := CurrencyDecimal{
		Id:           0,
		Type:         Type,
		Code:         Code,
		DecimalShift: DecimalShift,
		InUse:        true,
		CreatedAt:    time.Now().UTC().Format(time.DateTime),
	}
	return &cd
}

func (cd *CurrencyDecimal) GetStructFields() []any {
	fields := []any{
		&cd.Id,
		&cd.Type,
		&cd.Code,
		&cd.DecimalShift,
		&cd.InUse,
		&cd.CreatedAt,
	}
	return fields
}

func (cd *CurrencyDecimal) PrintFields() {
	fmt.Printf("CurrencyDecimal Fields: %v %v %v %v %v\n", 
		cd.Id, 
		cd.Type, 
		cd.Code, 
		cd.DecimalShift, 
		cd.CreatedAt, 
	)
}