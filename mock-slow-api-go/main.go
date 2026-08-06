package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Println("Doing something")
}

// w http.ResponseWriter, r *http.Request
func longTaskHandler(ctx context.Context) {
	// ctx := r.Context() // get the context that might have been passed

	// simulate a long running task like db fetch or some io
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Job finished successfully!")
	case <-ctx.Done():
		fmt.Println("Context ended, returning!: ", ctx.Err().Error())
	}
	// time.Sleep(2 * time.Second)
	// fmt.Println(ctx.Value("name"))
}

func main() {
	// fmt.Println("Hello World!")
	// ctx := context.TODO() 	// empty context with undefined intention
	bgCtx := context.Background() // same as above except the intent is defined i.e. we are trying to build a context with this context as backgroung/root ??
	// ctx := context.WithValue(bgCtx, "name", "ahbar")
	ctx, cancel := context.WithDeadline(bgCtx, time.Now().Add(1*time.Second))
	defer cancel()

	// doSomething(ctx)
	longTaskHandler(ctx)
}
