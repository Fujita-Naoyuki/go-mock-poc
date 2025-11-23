package xxxservice

import (
	dynamo "example/mock/repository"
	"fmt"
)

// プロダクトコードでxxxServiceeImplemを利用するための準備（その1）
// ユニットテストでMockを利用するための準備（その2）
var x XxxService

// プロダクトコードでxxxServiceeImplemを利用するための準備（その2）
// プロダクトコードからはGetXxxserviceを呼び出す。その際にinit()が呼び出される。
func init() {
	x = &xxxServiceeImplem{}
}

// プロダクトコードでxxxServiceeImplemを利用するための準備（その3）
func GetXxxservice() XxxService {
	return x
}

// ユニットテストでMockを利用するための準備（その2）
// プロダクトコードではxxxServiceeImplemしか使わせたくないので利用しないこと！
func SetXxxservice(a XxxService) {
	x = a
}

// Mockかするにあたりインターフェースに依存する必要があるためインターフェース作成
//
//go:generate moq -rm -out foo_moq.go . XxxService
type XxxService interface {
	Xxx() bool // DIする前だったらxxxservice.Xxx()と呼び出していたメソッドをここに記載
}

// インターフェースを実装するにあたりベースとなる構造体
// ★UTではxxxServiceeMockなど別構造体を準備する
type xxxServiceeImplem struct{}

// ↑構造体におけるXxx()の実装
// ★UTでは「func (impl *xxxServiceeMock) Xxx() {」のようなメソッドを作ることで処理を書き換えることができるようになる。
func (impl *xxxServiceeImplem) Xxx() bool {
	fmt.Println("Implem call")
	dynamo.Callxxxx()
	return true
}
