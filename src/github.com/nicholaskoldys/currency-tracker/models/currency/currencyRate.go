package currency

import "fmt"

type CurrencyRate struct {
	Id         int
	CodeBase   string
	CodeQuote  string
	MiddleRate int
	MonthRate  int
	YearAvg    int
	Trend      float64
	InputCount int
	DateOf     string
}

func CreateCurrencyRate(CodeBase string, CodeQuote string, MiddleRate int, MonthRate int, YearAvg int, Trend float64, InputCount int, DateOf string) *CurrencyRate {
	cr := CurrencyRate{
		Id:         0,
		CodeBase:   CodeBase,
		CodeQuote:  CodeQuote,
		MiddleRate: MiddleRate,
		MonthRate:  MonthRate,
		YearAvg:    YearAvg,
		Trend:      Trend,
		InputCount: InputCount,
		DateOf:     DateOf,
	}
	return &cr
}

func (cr *CurrencyRate) GetStructFields() []any {
	fields := []any{
		&cr.Id,
		&cr.CodeBase,
		&cr.CodeQuote,
		&cr.MiddleRate,
		&cr.MonthRate,
		&cr.YearAvg,
		&cr.Trend,
		&cr.InputCount,
		&cr.DateOf,
	}
	return fields
}

func (cr *CurrencyRate) PrintFields() {
	fmt.Printf("CurrencyRate Fields: %v %v %v %v %v %v %v %v %v\n", 
		cr.Id,
		cr.CodeBase,
		cr.CodeQuote,
		cr.MiddleRate,
		cr.MonthRate,
		cr.YearAvg,
		cr.Trend,
		cr.InputCount,
		cr.DateOf,
	)
}