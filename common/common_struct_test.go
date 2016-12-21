package common

import (
	"testing"
)

type Human struct {
	name string
	age  int
}

type Student struct {
	h    Human
	spec string
}

type StudentNew struct {
	Human
	spec string
}

func Test_If(t *testing.T) {
	h := Human{name: "junneyang", age: 28}
	t.Log(h)
	h.age += 10
	t.Log(h)

	stu := Student{h: h, spec: "Computer Science"}
	t.Log(stu)
	stu.h.name = "junneyang2"
	t.Log(stu)
	t.Log(h)

	stu_new := StudentNew{h, "Computer Science"}
	t.Log(stu_new)
	stu_new.age = 88
	t.Log(stu_new)
	t.Log(h)

	//	h = stu_new
	//	h = stu
	//	t.Log(h)
}
