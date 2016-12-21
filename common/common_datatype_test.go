package common

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func Test_SayHello(t *testing.T) {
	t.Log(SayHello("junneyang"))
}

func Test_DataType(t *testing.T) {
	var flag bool
	t.Log(flag)
	a := 123
	t.Log(reflect.TypeOf(a))
	b := 123.321
	t.Log(reflect.TypeOf(b))
	c := 123 + 321i
	t.Log(reflect.TypeOf(c))
	var str string
	t.Log(str == "")
	var cc complex128
	t.Log(cc == 0)
	//	t.Log(cc == nil)
	var up uintptr
	t.Log(up)
	t.Log(reflect.TypeOf(up))
}

func Test_Str(t *testing.T) {
	str := "Hello"
	//	str[0] = 'X'
	str_b := []byte(str)
	str_b[0] = 'X'
	t.Log(str)
	t.Log(string(str_b))

	str = "Hello"
	str_s := str[:]
	t.Log(reflect.TypeOf(str_s))
	//	str_s[0] = 'X'
	t.Log(str)
	t.Log(string(str_s))

	str = "您好,HELLO"
	str_array := []byte(str)
	str_array[0] = 'X'
	t.Log(str)
	t.Log(string(str_array))

	str = "您好,HELLO"
	str_rune := []rune(str)
	str_rune[0] = 'X'
	t.Log(str)
	t.Log(string(str_rune))

	t.Log("X" + str[1:])
	t.Log("Hello" + " " + "World")
}

func Test_Str_I18N(t *testing.T) {
	str := "您好,HELLO"
	t.Log(len(str))
	t.Log(len([]rune(str)))
	t.Log(utf8.RuneCount([]byte(str)))
	t.Log(utf8.RuneCountInString(str))
}

func Test_Array(t *testing.T) {
	arr := [5]int64{1, 2, 3, 4, 5}
	arr[0] = 88
	t.Log(arr)
	t.Log(reflect.TypeOf(arr))
	for k, v := range arr {
		t.Log(k, "=", v)
	}

	arra := [...]int64{}
	t.Log(arra)
	//	t.Log(arra == nil)
	t.Log(arra == [0]int64{})

	//	var a [...]int64
	//	t.Log(a)
	//	t.Log(reflect.TypeOf(a))
	//	t.Log(a == nil)
}

func Test_Slice(t *testing.T) {
	arr := [5]int64{1, 2, 3, 4, 5}
	t.Log(arr)
	t.Log(reflect.TypeOf(arr))
	slice := arr[:]
	t.Log(slice)
	t.Log(reflect.TypeOf(slice))
	slice[0] = 88
	t.Log(arr)
	t.Log(slice)
	for k, v := range arr {
		t.Log(k, "=", v)
	}

	sliceb := arr[0:2]
	sliceb = append(sliceb, 99)
	t.Log(arr)
	t.Log(sliceb)

	slicec := make([]int64, 5)
	t.Log(slicec)
	t.Log(reflect.TypeOf(slicec))
	t.Log(len(slicec))
	t.Log(cap(slicec))
	slicec = append(slicec, 99)
	t.Log(len(slicec))
	t.Log(cap(slicec))

	sliced := new([]int64)
	t.Log(sliced)
	t.Log(*sliced)
	t.Log(reflect.TypeOf(sliced))
	t.Log(reflect.TypeOf(*sliced))
	t.Log(sliced == nil)
	t.Log(*sliced == nil)

	var s []int64
	t.Log(s)
	t.Log(s == nil)

	ss := make([]int64, 2, 6)
	t.Log(ss)
	t.Log(ss[0])
	t.Log(ss == nil)
}

func Test_Map(t *testing.T) {
	var map_t map[string]int64
	t.Log(map_t == nil)

	// No Meanging For Length Specification
	mapa := make(map[string]int64, 2)
	t.Log(mapa)
	t.Log(reflect.TypeOf(mapa))
	t.Log(mapa == nil)
	t.Log(len(mapa))
	//	t.Log(cap(mapa))
	mapa["A"] = 123
	mapa["B"] = 456
	t.Log(mapa)
	t.Log(mapa["B"])

	for k, v := range mapa {
		t.Log(k, "=", v)
	}
	t.Log(len(mapa))
	delete(mapa, "B")
	t.Log(len(mapa))
	t.Log(mapa)
	t.Log(mapa["B"])
	v, ok := mapa["C"]
	t.Log(v, ok)
}
