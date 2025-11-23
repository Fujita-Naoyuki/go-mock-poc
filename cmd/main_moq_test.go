package main

import (
	xxxservice "example/mock/service"
	"fmt"
	"testing"
)

func TestMoqMock_main(t *testing.T) {
	back := xxxservice.GetXxxservice()
	defer moqend(back)
	xxxservice.SetXxxservice(&xxxservice.XxxServiceMock{
		XxxFunc: func() bool {
			fmt.Println("mock call")
			return true
		},
	})

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

func moqend(backup xxxservice.XxxService) {
	xxxservice.SetXxxservice(backup)
}

// Mock処理をしていないので登録されているDIのまま動く
func TestMoq_main(t *testing.T) {

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
