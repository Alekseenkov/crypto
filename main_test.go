package main

import (
	"fmt"
	"math/big"
	"mycrypto/module"
	"reflect"
	"testing"
)

type testValue struct {
	a, m   *big.Int
	jacobi int
}

type testValueComp2 struct {
	p, a, n *big.Int
	rez     []module.AnswerToEquation
}

var testValues = []testValue{
	{big.NewInt(24), big.NewInt(47), 1},
	{big.NewInt(42), big.NewInt(47), 1},
	{big.NewInt(29), big.NewInt(53), 1},
	{big.NewInt(46), big.NewInt(97), -1},
	{big.NewInt(180), big.NewInt(307), -1},
	{big.NewInt(572), big.NewInt(971), -1},
}

///////////////////////////////////
var rez2 = []module.AnswerToEquation{
	module.NewAnswerToEquation(big.NewInt(2), big.NewInt(7)),
	module.NewAnswerToEquation(big.NewInt(5), big.NewInt(7)),
}

var testValuesComp2 = []testValueComp2{
	{big.NewInt(7), big.NewInt(4), big.NewInt(3), rez2},
}

func TestJacobi(t *testing.T) {
	fmt.Println("Test for Jacobi")

	for _, test := range testValues {
		res, err := module.Jacobimy(test.a, test.m)

		if err == false {
			if res != test.jacobi {
				t.Error("For", test.a, "and", test.m,
					"expected", test.jacobi,
					"got", res,
				)
			}
		} else {
			t.Error("Function returned error for:",
				test.a, "and", test.m,
				";Error: ", err,
			)
		}
	}
}

func TestSolvingComparisons2Degree(t *testing.T) {

	fmt.Println("Test for SolvingComparisons2Degree")

	for _, test := range testValuesComp2 {
		var s = []module.AnswerToEquation{}
		s = module.SolvingComparisons2Degree(test.p, test.a, test.n)

		if reflect.DeepEqual(s, test.rez) == false {
			t.Error(" error")
		}

	}
}
