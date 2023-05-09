package db

/* func GetTyped[mwFs models.ModelWrapper[models.IFieldedStruct]](queryString string, structSlice *[]*mwFs, cm models.ICurrencyModel[mwFs]) {
	db := getDB()
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

	// ! PROBLEM - when cm.GetCopy() is used, it reverts to the base type, its no longer an interface.
	// ! With type safe checking Golang Doesn't think it will fit into the spot because its type has been
	// ! reverted back to Type=any
	for rows.Next() {
		newCm := cm.GetCopy()
		// capableInter := (any(&newCm)).(models.ICurrencyModel[T])
		// rowAssign[mwFs](rows, structSlice, newCm)
		// fmt.Printf("%#v, %v, %v \n", newCm, *newCm, capableInter)
		fmt.Printf("%#v, %v, %v \n", newCm, newCm, *newCm)
	}
} */

/* func GetQuery(queryString string, structSlice interface{}) {

	db := getDB()
	if db == nil {
		util.Debug(9, "Database returned nothing")
		panic("no database")
	}

	fmt.Println("Opened Connection");

	stmt, err := db.Prepare(queryString)
	if err != nil {
		util.Debug(9, err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		util.Debug(9, err.Error())
	}

	if err := scanToStructs( rows, structSlice ); err != nil {
		util.Debug(9, err.Error())
	}
	// for rows.Next() {
	// 	// var newCd currency.CurrencyDecimal
	// 	// var typeStr string

	// 	fmt.Println("   Object: ", obj)
	// 	// if err := rows.Scan(&newCd.Id, &typeStr, &newCd.Code, &newCd.DecimalShift, &newCd.InUse, &newCd.CreatedAt); err != nil {
	// 	if err := rows.Scan( objLocations... ); err != nil {
	// 		util.Debug(9, err.Error())
	// 	}
	// 	// newCd.SetDecimalType(typeStr)

	// 	fmt.Println("   Current Object: ", obj)
	// }
}

func scanToStructs(rows *sql.Rows, structSlice interface{}) error {

	var getFieldByTag = func(rStruct reflect.Value, tagLookup string) reflect.Value {
		for i, field := range reflect.VisibleFields(rStruct.Type()) {
			if field.Tag.Get("db") == tagLookup {
				return rStruct.Field(i)
			}
		}
		util.Debug(9, "scanStructs: Returning Empty Value")
		return reflect.Value{}
	}

	// Check if the results interface is a pointer to a slice
	resultValue := reflect.ValueOf(structSlice)
	if resultValue.Kind() != reflect.Ptr || resultValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result must be a pointer to a slice")
	}

	// Check if the results interface is made of a struct type
	sliceValue := resultValue.Elem()
	sliceElementType := sliceValue.Type().Elem()
	if sliceElementType.Kind() != reflect.Struct {
		return fmt.Errorf("slice element type must be a struct")
	}

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	for rows.Next() {
		structValue := reflect.New(sliceElementType).Elem()
		values := make([]interface{}, len(columns))

		for i, column := range columns {
			fieldValue := getFieldByTag(structValue, column)

			if fieldValue.IsValid() && fieldValue.CanSet() {
				values[i] = fieldValue.Addr().Interface()
			} else {
				var dummy interface{}
				values[i] = &dummy
			}
		}

		// Take the values array and apply the scanned rows
		if err:= rows.Scan(values...); err != nil { return err }

		// Active apply new struct to pointer slice
		sliceValue.Set( reflect.Append(sliceValue, structValue) )
	}
	return rows.Err()
} */

/* func rowAssign[T *models.IFieldedStruct](rows *sql.Rows, structSlice *[]*models.ModelWrapper[T], obj *models.ModelWrapper[T]) {

	if err:= rows.Scan( obj.GetStructFields()... ); err != nil {
		util.Debug(9, err.Error())
	}

	obj.PrintFields()
	// *structSlice = append(*structSlice, obj.Ptr())
} */

/* func GetTyped[T *models.IFieldedStruct](queryString string, structSlice *[]*models.ModelWrapper[T], w models.ModelWrapper[T]) {
	db := getDB()
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

	for rows.Next() {
		newCm := w.GetCopy()
		rowAssign[T](rows, structSlice, newCm)
		fmt.Printf("%v \n", *newCm.Model)
	}
} */