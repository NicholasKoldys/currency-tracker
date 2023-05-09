package repo

// type currencyRepo []currency.Currency

// var currencyRepoInstance currencyRepo
// var lock = &sync.Mutex{}

// func createRepo() currencyRepo {

// 	// return []currency.Currency{ *currency1 }
// }

// func GetCurrencyRepository() *currencyRepo {
// 	if currencyRepoInstance == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if currencyRepoInstance == nil {
// 			util.Debug(9, "Creating 'currencyRepoInstance'.")
// 			currencyRepoInstance = createRepo()
// 		}
// 	}
// 	return &currencyRepoInstance
// }