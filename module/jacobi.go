package module

import (
	"fmt"
	"math/big"
)

func Jacobimy(a, m *big.Int) (jacobi int, err bool) {

	if big.NewInt(0).Mod(m, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return 0, true ///"м" не может быть нечетным"
	}

	var result = int(1)
	var buf = big.NewInt(0)
	var num = big.NewInt(0).Set(a)
	var den = big.NewInt(0).Set(m)

	for num.Cmp(big.NewInt(0)) != 0 {
		for buf.And(num, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
			num.Rsh(num, 1)
			buf.Mod(den, big.NewInt(8))
			if buf.Cmp(big.NewInt(3)) == 0 || buf.Cmp(big.NewInt(5)) == 0 {
				result = result * -1
			}
		}

		buf.Set(num)
		num.Set(den)
		den.Set(buf)

		if buf.Mod(num, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 && buf.Mod(den, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 {
			result = result * -1
		}

		num.Mod(num, den)
	}

	if den.Cmp(big.NewInt(1)) == 0 {
		return result, false
	}

	return 0, true //Знаменатель не равен единице
}

func InterfaceJacobi() {
	var a = new(big.Int)
	var b = new(big.Int)

	var jacobi int
	var err bool

	fmt.Println(" (a/b) = Jacobi")

	a = NewBigIntSetString("Введите a")
	b = NewBigIntSetString("Введите b")

	jacobi, err = Jacobimy(a, b)

	if err == true {
		fmt.Println(" err ")
	} else {
		fmt.Println("(", a, "/", b, ")=", jacobi)
	}

}
