package module

import (
	"fmt"
	"math/big"
)

func diophantineEquations(b *big.Int, a *big.Int) (x *big.Int, y *big.Int) {

	var vector = []*big.Int{big.NewInt(0), big.NewInt(0)}                 // q n {0,0}
	var vector2 = []*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(1)} // P n {0,0,1}
	var vector3 = []*big.Int{big.NewInt(0), big.NewInt(1)}                // Q n {0,1}

	c := new(big.Int)
	t := new(big.Int)

	for b.String() != "0" {
		t = b
		c, b = divMod(a, b)
		a = t
		vector = append(vector, c)
	}

	for iter := 2; iter < len(vector); iter++ {

		buf1 := new(big.Int)
		buf1.Mul(vector[iter], vector3[iter-1])
		buf1.Add(buf1, vector3[iter-2])
		vector3 = append(vector3, buf1)

		if iter > 2 {
			buf := new(big.Int)
			buf.Mul(vector[iter], vector2[iter-1])
			buf.Add(buf, vector2[iter-2])
			vector2 = append(vector2, buf)
		}
	}

	if (len(vector2)-3)%2 == 1 {

		return vector3[len(vector3)-2].Mul(vector3[len(vector3)-2], big.NewInt(-1)), vector2[len(vector2)-2].Mul(vector2[len(vector2)-2], big.NewInt(-1)) // если степень не четная домножаем на -1
	}

	return vector3[len(vector3)-2], vector2[len(vector2)-2]

}

func InterfaceDiophantine() {

	var a = new(big.Int)
	var b = new(big.Int)
	var c = new(big.Int)

	fmt.Println("введите  параметры уравнения ", "ax-by = c ")
	a = NewBigIntSetString("Введите a")
	b = NewBigIntSetString("Введите b")
	c = NewBigIntSetString("Введите c")

	if c.String() == "0" {
		fmt.Println("нет решений с = 0")
	}

	var flag1 = new(big.Int)
	var flag2 = new(big.Int)

	flag1 = flag1.Mod(a, c)
	flag2 = flag2.Mod(b, c)

	if (flag1.String() != "0") || (flag2.String() != "0") {

		fmt.Println("нет решений  в целых числах ")
	}

	if (flag1.String() == "0") && (flag2.String() == "0") {
		x := new(big.Int)
		y := new(big.Int)

		x, y = diophantineEquations(a.Div(a, c), b.Div(b, c))

		fmt.Println("x =", x, "  y =", y)
	}
}
