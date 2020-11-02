package prototype

import (
	"encoding/json"
	"fmt"
	"testing"
)

type (
	IClone interface {
		Clone() IClone
	}

	Type1 struct {
		name  string
		slice []int
	}

	Type2 struct {
		name  string
		slice []int
	}
)

// 深拷贝 在type1结构中 tc.slice 是切片类型该类型为引用拷贝，需要进行转换为值拷贝的数据类型后再进行转换达到深拷贝的原型
func (t1 *Type1) Clone() IClone {
	tc := *t1
	var sl = make([]int, len(tc.slice), cap(tc.slice))
	s, err := json.Marshal(tc.slice)
	if err != nil {
		fmt.Println("json 序列化失败", err.Error())
	}
	json.Unmarshal(s, &sl)
	tc.slice = sl
	return &tc
}

// 浅拷贝 在go语言中struct 是值拷贝，获取到值struct 的值赋值给一个新的变量就clone成功
func (t2 *Type2) Clone() IClone {
	tc := *t2
	return &tc
}

func NewType1() IClone {
	s := make([]int, 10, 10)
	s[0] = 123
	return &Type1{
		name:  "type1",
		slice: s,
	}
}

func NewType2() IClone {
	s := make([]int, 10, 10)
	s[0] = 123
	return &Type2{
		name:  "type1",
		slice: s,
	}
}

func TestDeepClone(t *testing.T) {
	t1 := NewType1()
	ct1, ok := t1.Clone().(*Type1)
	if ok {
		ct1.slice[0] = 321
		fmt.Println(t1, ct1)
	} else {
		fmt.Println("类型断言失败")
	}

}

func TestShallowClone(t *testing.T) {
	t2 := NewType2()
	ct2, ok := t2.Clone().(*Type2)
	if ok {
		ct2.slice[0] = 321
		fmt.Println(t2, ct2)
	} else {
		fmt.Println("类型断言失败")
	}
}
