package main

import (
	"fmt"
	"net/http"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
)

func main() {
	fmt.Println(rest_errors.NewRestError("impleted me", http.StatusNotImplemented, "not_implemented", nil))
	fmt.Println(rest_errors.NewUnauthorizedError())
}
