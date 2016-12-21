package common

import (
	"testing"
)

func Test_If(t *testing.T) {

	// Use i in this if-else block
	if i := 9; i > 10 {
		t.Log(i, "i > 10")
	} else if i == 10 {
		t.Log(i, "i = 10")
	} else {
		t.Log(i, "i < 10")
		goto HELL
	}

HELL:
	t.Log("Welcome to the hell of floor -18!")
}

func Test_For(t *testing.T) {
	arr := [...]int64{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(i, "=>", arr[i])
	}

	sum := 0
	for {
		sum += 1
		t.Log(sum)
		if sum >= 10 {
			break
		}
	}

	for k, v := range arr {
		t.Log(k, "=>", v)
	}
}

func Test_Switch(t *testing.T) {
	i := 5
	switch i {
	case 2:
		t.Log("i = 2")
	case 3, 4:
		t.Log("i = 3, 4")
	case 5:
		t.Log("i = 5")
		fallthrough
	default:
		t.Log("DEFAULT")
	}
}
