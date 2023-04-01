package main

import (
	"context"
	"fmt"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "api-key", "my-super-api-key")
}

func doSomethingCool(ctx context.Context) {
	apikey := ctx.Value("api-key")
	fmt.Println(apikey)

}

func main() {
	fmt.Println("Go context")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	doSomethingCool(ctx)
}
