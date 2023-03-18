package main

import (
	"fmt"
	"math/big"
	"mycrypto/fftmy"
)

func main() {

	//module.InterfaceComparison1()
	//module.InterfaceComparison2()
	//module.InterfaceDiophantine()
	//module.InterfaceJacobi()

	var T = new(big.Int)
	var E = new(big.Int)
	var X = new(big.Int)

	fmt.Println(" a*b = x ")

	T = big.NewInt(5)
	E = big.NewInt(5)

	X = fftmy.MulFFT(T, E)

	fmt.Println("x =", X)

}
