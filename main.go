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

	fmt.Println(bigfft.MulFFT(big.NewInt(-79897644433432235), big.NewInt(77886632408082)))

}
