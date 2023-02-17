package models

type Test struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type NoEvmAddress struct {
	Address string `json:"address"`
}
