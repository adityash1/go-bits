// package main

// import "fmt"

// type NumberStorer interface {
// 	GetAll() ([]int, error)
// 	Put(int) error
// }

// type MongoDBNumberStore struct {
// 	// implement db connection
// }

// func (m MongoDBNumberStore) GetAll() ([]int, error) {
// 	return []int{1, 2, 3}, nil
// }

// func (m MongoDBNumberStore) Put(number int) error {
// 	fmt.Printf("put number data into storage")
// 	return nil
// }

// type ApiServer struct {
// 	numberStore NumberStorer
// }

// func main() {
// 	apiServer := ApiServer{
// 		numberStore: MongoDBNumberStore{},
// 	}

// 	numbers, err := apiServer.numberStore.GetAll()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("getall worked and numbers are %v", numbers)

// 	if err = apiServer.numberStore.Put(10); err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("put worked")

// }
