package module

import (
	"fmt"
	"math/big"
	//"os"
)

func comparison1(t_ *big.Int, e *big.Int, n *big.Int) *big.Int { //-> X *big.Int

	/// T X (=) E (mod N )   Решить сравнение 111x є 75(mod 322). решение сравнений первой степени

	var vector = []*big.Int{big.NewInt(0), big.NewInt(0)}  // q n {0,0}
	var vector2 = []*big.Int{big.NewInt(0), big.NewInt(1)} // P n {0,1}

	var T = new(big.Int).Set(t_)
	var E = new(big.Int).Set(e)
	var N = new(big.Int).Set(n)
	var X = new(big.Int)

	X.SetString("0", 10)

	c := new(big.Int)
	t := new(big.Int)
	var N_ = N

	for T.String() != "0" {
		t = T
		c, T = divMod(N_, T)
		N_ = t
		//fmt.Println(c, " ", T)
		vector = append(vector, c)
	}

	for iter := 2; iter < len(vector); iter++ {

		var buf = new(big.Int)
		buf.Mul(vector[iter], vector2[iter-1])
		vector2 = append(vector2, buf.Add(buf, vector2[iter-2]))
	}

	X.Mul(vector2[len(vector2)-2], E) //получение произведения x*E

	if (len(vector2)-3)%2 == 1 {
		X = X.Mul(X, big.NewInt(-1))
	}

	X.Mod(X, N) // if N= 0 panic

	return big.NewInt(0).Set(X)

}

func InterfaceComparison1() {
	var T = new(big.Int)
	var E = new(big.Int)
	var N = new(big.Int)
	var X = new(big.Int)

	fmt.Println(" X (=) E (mod N )")

	T = NewBigIntSetString("Введите Т")
	E = NewBigIntSetString("Введите Е")
	N = NewBigIntSetString("Введите N")

	X = comparison1(T, E, N)

	fmt.Println("x =", X)

}
