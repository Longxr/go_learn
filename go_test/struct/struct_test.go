package struct_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Result struct {
	Status int
}

func newResult(status int) Result {
	return Result{
		Status: status,
	}
}

func TestStructJson(t *testing.T) {
	data := []byte(`{"status": 200}`)
	result := &Result{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("result=%+v\n", result)
	fmt.Printf("result address %p\n", result)
}
