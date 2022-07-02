package interview

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//mem := new(map[string]any)
	//(*mem)["hh"] = "111"
	//fmt.Println((*mem)["hh"])
	//delete((*mem), "hh")

	mem2 := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range mem2 {
		if v%2 == 1 {
			delete(mem2, k)
		} else {
			mem2[fmt.Sprintf("%d", v*2)] = v * 2
		}
	}
	for k, v := range mem2 {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func Test3(t *testing.T) {
	type x struct {
		val int
	}
	mem2 := map[string]x{
		"a": {
			val: 1,
		},
		"b": {
			val: 1,
		},
	}
	for _, v := range mem2 {
		//mem2[k] = x{val: 2 * v.val}
		tmpV := &v
		tmpV.val = 2
	}
	for k, v := range mem2 {
		fmt.Println(k)
		fmt.Println(v)
	}
}

type jsonA struct {
	a int `json:"a"`
}

func TestJson(t *testing.T) {
	ja := &jsonA{
		a: 1,
	}
	marshal, _ := json.Marshal(ja)
	fmt.Println(string(marshal))
}
