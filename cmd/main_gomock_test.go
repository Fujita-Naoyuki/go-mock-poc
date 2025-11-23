package main

import (
	mock_xxxservice "example/mock/mock"
	xxxservice "example/mock/service"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestGoMock_main(t *testing.T) {
	ctrl := gomock.NewController(t) // gomockを使うときのおまじない
	back := xxxservice.GetXxxservice()
	defer goend(back)

	app := mock_xxxservice.NewMockXxxService(ctrl)
	app.EXPECT().Xxx().Return().AnyTimes() // AnyTimes()がついているのは{name: "mock"},{name: "imple"}で2回呼び出しされる関係(このテストコードの書き方と相性悪いなぁ)
	xxxservice.SetXxxservice(app)

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

func goend(backup xxxservice.XxxService) {
	xxxservice.SetXxxservice(backup)
}

// Mock処理をしていないので登録されているDIのまま動く
func TestGo_main(t *testing.T) {

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
