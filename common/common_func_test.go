package common

import (
	"testing"
)

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	spec string
}

type Int int

func AppendSlice(s []int64, i int64) (ret []int64) {
	ret = append(s, i)
	return
}

func ModSlice(s []int64, i int64) (ret []int64) {
	s[0] = i
	ret = s
	return
}

func (h Human) ModHuman(name string, age int) (ret Human) {
	h.name = name
	h.age = age
	ret = h
	return
}

func (h *Human) ModHuman_New(name string, age int) {
	h.name = name
	h.age = age
}

//func (h Human) ModHuman_New(name string, age int) {
//	h.name = name
//	h.age = age
//}

func (i Int) Inc() (ret Int) {
	ret = i + 1
	return
}

func (i *Int) Inc_New() {
	(*i) += 10
}

// REF: http://wiki.jikexueyuan.com/project/the-way-to-go/10.6.html
// 类型和作用在它上面定义的方法必须在同一个包里定义，这就是为什么不能在 int、float 或类似这些的类型上定义方法
//func (i *int) Inc_New_New() {
//	(*i) += 10
//}

func Test_If(t *testing.T) {
	s := []int64{1, 2, 3, 4, 5}
	ret := AppendSlice(s, 88)
	t.Log(s)
	t.Log(ret)

	ret = ModSlice(s, 99)
	t.Log(s)
	t.Log(ret)

	h := Human{name: "junneyang", age: 28}
	t.Log(h)
	hh := h.ModHuman("junneyang_new", 88)
	t.Log(h)
	t.Log(hh)

	h.ModHuman_New("junneyang_mod", 99)
	t.Log(h)
	(&h).ModHuman_New("junneyang_mod_mod", 999)
	t.Log(h)

	stu := Student{h, "Computer Science"}
	t.Log(stu)
	t.Log(h)
	stu.ModHuman_New("junneyang_in", 888)
	t.Log(stu)
	t.Log(h)

	i := Int(10)
	t.Log(i)
	t.Log(i.Inc())
	t.Log(i)

	i.Inc_New()
	t.Log(i)

	//	j := 88
	//	j.Inc_New_New()
	//	t.Log(j)
}
