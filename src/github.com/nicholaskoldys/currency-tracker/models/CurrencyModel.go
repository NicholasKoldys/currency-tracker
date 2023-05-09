package models

type ICurrencyModel[W ModelWrapper[IFieldedStruct]] interface {
	GetCopy() *W
	GetStructFields() []any
	PrintFields()
}

type IFieldedStruct interface {
	GetStructFields() []any
	PrintFields()
}

type ModelWrapper[F IFieldedStruct] struct {
	Model *F
}

func (w ModelWrapper[F]) GetCopy() *ModelWrapper[F] {
	new := *w.Model
	return &ModelWrapper[F]{Model: &new}
}

/* func (w ModelWrapper[T]) makeInterfacableCopy() *ModelWrapper[T] {
	new := *w.Model
	return &ModelWrapper[T]{Model: &new}
} */

func (w ModelWrapper[F]) GetStructFields() []any {
	return (*w.Model).GetStructFields()
}

func (w ModelWrapper[F]) PrintFields() {
	(*w.Model).PrintFields()
}

// type IFieldedStruct interface {
// 	GetStructFields() []any
// }

// type ModelWrapper[fs IFieldedStruct] struct {
// 	Model *fs
// }

// func (w ModelWrapper[fs]) Ptr() *ModelWrapper[fs] {
// 	return &w
// }

// func (w ModelWrapper[fs]) GetCopy() *ModelWrapper[fs] {
// 	new := *w.Model
// 	return &ModelWrapper[fs]{Model: &new}
// }

// /* func (w ModelWrapper[T]) makeInterfacableCopy() *ModelWrapper[T] {
// 	new := *w.Model
// 	return &ModelWrapper[T]{Model: &new}
// } */

// func (w ModelWrapper[fs]) GetStructFields() []any {
// 	return (*w.Model).GetStructFields()
// }

/* func (cd *CurrencyDecimal) ResetToInterface() ICurrencyModel[CurrencyDecimal] {

	return func(cm ICurrencyModel[CurrencyDecimal]) *ICurrencyModel[CurrencyDecimal] {
		return &cm
	}(*cd)
} */

// func AccessFields(cm ICurrencyModel) {
// 	for i, field := range cm.GetStructFields() {
// 		fmt.Println(i, field)
// 	}
// }