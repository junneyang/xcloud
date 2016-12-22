package common

import (
	"fmt"
	"reflect"
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

type PeopleProtocol interface {
	SayHello() (ret string)
}

type PeopleProtocolNew interface {
	PeopleProtocol
}

func (h Human) SayHello() (ret string) {
	return fmt.Sprintf("name: %v, age: %v", h.name, h.age)
}

// StudentNew Redefine The Method SayHello
func (s StudentNew) SayHello() (ret string) {
	return fmt.Sprintf("name: %v, age: %v, spec: %v", s.name, s.age, s.spec)
}

func Test_Interface(t *testing.T) {
	h := Human{name: "junneyang", age: 28}
	var p PeopleProtocol
	p = h
	t.Log(p.SayHello())

	stu_new := StudentNew{h, "Computer Science"}
	p = stu_new
	t.Log(p.SayHello())

	//	h = stu_new
	//	t.Log(h.SayHello())

	s := make([]PeopleProtocol, 2)
	s[0], s[1] = h, stu_new
	t.Log(s[0].SayHello())
	t.Log(s[1].SayHello())

	t.Log(reflect.TypeOf(s[0]))
	t.Log(reflect.TypeOf(s[1]))

	v1, ok := s[0].(Human)
	t.Log(v1, ok)
	v2, ok := s[1].(Human)
	t.Log(v2, ok)
	v3, ok := s[1].(StudentNew)
	t.Log(v3, ok)

	switch v := s[1].(type) {
	case Human:
		t.Log("Human->", v)
	case StudentNew:
		t.Log("StudentNew->", v)
	default:
		t.Log("Don't know it's type")
	}

	var p_new PeopleProtocolNew
	p_new = stu_new
	t.Log(p_new.SayHello())

	//	human := reflect.New(reflect.TypeOf(h)).Interface()
	human := reflect.New(reflect.ValueOf(h).Type()).Interface()
	t.Log(human)
	hh := human.(*Human)
	hh.name = "reflect_name"
	hh.age = 88
	t.Log(*hh)
	t.Log(hh.SayHello())
}
