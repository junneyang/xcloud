package common

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func SayHello(name string) (ret string) {
	return fmt.Sprintf("hello, my name is %s. 欢迎您 %s", name, name)
}

func AddStr(a, b string) (ret string) {
	return fmt.Sprintf("a: %s, b: %s", a, b)
}

func Sum() (sum int64, err error) {
	var PARAMA, PARAMB string = os.Getenv("PARAMA"), os.Getenv("PARAMB")
	fmt.Println(AddStr(PARAMA, PARAMB))

	a, err := strconv.ParseInt(PARAMA, 10, 64)
	if err != nil {
		return sum, errors.Wrapf(err, "ParseInt Failed, PARAM: %v", PARAMA)
	}

	b, err := strconv.ParseInt(PARAMB, 10, 64)
	if err != nil {
		return sum, errors.Wrapf(err, "ParseInt Failed, PARAM: %v", PARAMB)
	}

	sum = a + b
	return sum, err
}
