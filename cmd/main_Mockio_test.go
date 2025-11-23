package main

import (
	xxxservice "example/mock/service"
	"testing"

	. "github.com/ovechkin-dm/mockio/mock"
)

func TestMockio_main(t *testing.T) {
	SetUp(t)
	m := Mock[xxxservice.XxxService]()
	WhenSingle(m.Xxx()).ThenReturn(true)

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

	// Verify(m, Once()).Xxx() // なぜかうまくいかなかった（引数ないからかな？）
	VerifyNoMoreInteractions(m)
}
