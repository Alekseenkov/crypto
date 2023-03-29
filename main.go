package main

import (
	"fmt"
	"main/bigfft"
	"math/big"
	//"mycrypto/fftmy"
	//"main/bigfft"
)

func main() {

	//module.InterfaceComparison1()
	//module.InterfaceComparison2()
	//module.InterfaceDiophantine()
	//module.InterfaceJacobi()

	var str = "3333338888888"
	var a = big.NewInt(0)
	a.SetString(str, 10)

	str = ""
	str = "44444447777777"
	var b = big.NewInt(0)
	b.SetString(str, 10)

	fmt.Println(bigfft.Mulmy(a, b))

}
