package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	// ctx := context.WithValue(context.Background(), "username", "aditya")
	userId, err := fetchUserID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v -> %+v\n", time.Since(start), userId)
}

// context can also be used like this
// func fetchUserID(ctx context.Context) (string, error) {}

func fetchUserID() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	val := ctx.Value("username")
	fmt.Println("the value = ", val)

	type result struct {
		userId string
		err    error
	}
	resutlch := make(chan result, 1)

	go func() {
		res, err := thirdPartyHTTPCall()
		resutlch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resutlch:
		return res.userId, res.err
	}
}

func thirdPartyHTTPCall() (string, error) {
	// time.Sleep(time.Millisecond * 500)
	return "user id 1", nil
}
