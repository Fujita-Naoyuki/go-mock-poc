package main

import (
	xxxservice "example/mock/service"
	"fmt"
	"testing"
)

type MockService struct{}

func (impl *MockService) Xxx() bool {
	fmt.Println("mock call")
	return true
}

func TestMock_main(t *testing.T) {
	back := xxxservice.GetXxxservice()
	defer end(back)
	xxxservice.SetXxxservice(&MockService{})

	// テストそのもの
	tests := []struct {
		name string
	}{
		{name: "mock"},
		{name: "imple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func end(backup xxxservice.XxxService) {
	xxxservice.SetXxxservice(backup)
}

// Mock処理をしていないので登録されているDIのまま動く
func Test_main(t *testing.T) {

	// テストそのもの
	tests := []struct {
		name string
	}{
		{name: "imple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
