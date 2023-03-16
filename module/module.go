package module

import (
	"fmt"
	"math/big"
	"os"
)

func NewBigIntSetString(str string) *big.Int {
	var add = new(big.Int)

	for {
		var ok = false
		var strBigint string
		fmt.Println(str)
		fmt.Fscan(os.Stdin, &strBigint)
		add, ok = add.SetString(strBigint, 10)
		if ok {
			break
		}
		add = new(big.Int)
		strBigint = ""
	}
	return add
}

func divMod(a *big.Int, b *big.Int) (*big.Int, *big.Int) { // -> div int   mod int

	div_ := new(big.Int)
	mod_ := new(big.Int)

	if b.String() == "0" {
		return div_.SetInt64(0), a
	}

	return div_.Div(a, b), mod_.Mod(a, b)
}
