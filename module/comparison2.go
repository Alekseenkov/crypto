package module

import (
	"fmt"
	"math/big"
)

type AnswerToEquation struct { // x  = a (mod p)
	a, p *big.Int
}

func NewAnswerToEquation(a *big.Int, p *big.Int) AnswerToEquation {

	var AN = AnswerToEquation{a: big.NewInt(0).Set(a), p: big.NewInt(0).Set(p)}

	return AN
}

func (Eq AnswerToEquation) Print() {
	fmt.Println("x = ", Eq.a.String(), "mod(", Eq.p.String(), ")")
}

func enumerationOptions(p *big.Int) (*big.Int, *big.Int) { //p = 2^k * h + 1

	var p_ = new(big.Int)
	var h = new(big.Int)
	var temp = new(big.Int)
	var exp2_ = new(big.Int)

	p_.Sub(p, big.NewInt(1)) // вычитаем единицу

	for iter := big.NewInt(0); iter.Int64() < 100; iter.Add(iter, big.NewInt(1)) { // ограничили перебор 2 степени до сотой

		exp2_.Exp(big.NewInt(2), iter, big.NewInt(0))
		h.Div(p_, exp2_)

		var isParityCheck = (temp.Mod(h, big.NewInt(2)).String() == "1") ///  проверка что  h нечетно
		var isZeroRemainderCheck = (temp.Mod(p_, exp2_).String() == "0") //проверка что  проверка что остаток = 0

		if isParityCheck && isZeroRemainderCheck {
			return iter, h
		}
	}

	return big.NewInt(0), big.NewInt(0) // лучше вызвать панику
}

func SolvingComparisons2Degree(p, a, n *big.Int) []AnswerToEquation {

	var sliseAnswer = []AnswerToEquation{}

	var a1 = new(big.Int)
	var a2 = new(big.Int)
	var n1 = new(big.Int)

	var k = new(big.Int)
	var h = new(big.Int)

	k, h = enumerationOptions(p) //перевое действие алгоритма

	var temp = new(big.Int)

	temp.Add(h, big.NewInt(1))
	temp.Div(temp, big.NewInt(2))

	a1.Exp(a, temp, p) // второе действие алгоритма

	a2.Exp(a, big.NewInt(-1), p)
	n1 = n1.Exp(n, h, p)
	var n2 = big.NewInt(1)
	var j = big.NewInt(0)

	var t = new(big.Int)
	var b = new(big.Int)
	var c = new(big.Int)

	if k.Cmp(big.NewInt(1)) == 1 {

		for i := 0; i < int(t.Sub(k, big.NewInt(-2)).Int64()); i++ { /// третье действие алгоритма

			b.Mul(a1, n2)
			b.Mod(b, p) /// 3.1

			c.Mul(a2, b.Exp(b, big.NewInt(2), big.NewInt(0)))
			c.Mod(b, p) /// 3.1

			// наименьший вычет
			var d = new(big.Int)
			var exp_ = new(big.Int)

			exp_.Sub(k, big.NewInt(-2))
			exp_.Sub(exp_, big.NewInt(int64(i)))

			exp_.Exp(big.NewInt(2), exp_, big.NewInt(0))
			d.Exp(c, exp_, p)

			if d.String() == "1" {
				j.SetInt64(0)
			}
			if d.String() == "-1" {
				j.SetInt64(1)
			}

			var n2 = new(big.Int)

			exp_.Exp(big.NewInt(2), j, big.NewInt(0))
			exp_.Mul(exp_, j)

			exp_.Exp(n1, exp_, big.NewInt(0))
			n2.Mul(n2, n1)
			n2.Mod(n2, p)
			//fmt.Println("x =  (+-)", a1, " ", n2, "mod(", p, ")")
			///
			var a_rez = new(big.Int)

			a_rez.Mul(a1, n2)
			a_rez.Mod(a_rez, p)
			sliseAnswer = append(sliseAnswer, AnswerToEquation{a: big.NewInt(0).Set(a_rez), p: big.NewInt(0).Set(p)})

			a_rez.Mul(a1, big.NewInt(-1))
			a_rez.Mul(a_rez, n2)
			a_rez.Mod(a_rez, p)
			sliseAnswer = append(sliseAnswer, AnswerToEquation{a: big.NewInt(0).Set(a_rez), p: big.NewInt(0).Set(p)})
			///
		}

	} else {

		var a_rez = new(big.Int)

		a_rez.Mul(a1, n2)
		a_rez.Mod(a_rez, p)
		sliseAnswer = append(sliseAnswer, AnswerToEquation{a: big.NewInt(0).Set(a_rez), p: big.NewInt(0).Set(p)})

		a_rez.Mul(a1, big.NewInt(-1))
		a_rez.Mul(a_rez, n2)
		a_rez.Mod(a_rez, p)
		sliseAnswer = append(sliseAnswer, AnswerToEquation{a: big.NewInt(0).Set(a_rez), p: big.NewInt(0).Set(p)})
		//fmt.Println("x =  (+-)", a1, " ", n2, "mod(", p, ")")
	}
	return sliseAnswer
}

func InterfaceComparison2() {

	var p = new(big.Int)
	var a = new(big.Int)
	var n = new(big.Int)

	fmt.Println("_________")

	p = NewBigIntSetString("Введите P")
	a = NewBigIntSetString("Введите A")
	n = NewBigIntSetString("Введите N")

	var jacobi1, jacobi2 int
	var err1, err2 bool

	jacobi1, err1 = Jacobimy(a, p)
	jacobi2, err2 = Jacobimy(n, p)

	if (jacobi1 == 1) && (jacobi2 == -1) && (err1 == false) && (err2 == false) {

		var slise = []AnswerToEquation{}

		slise = SolvingComparisons2Degree(p, a, n)
		//slise = SolvingComparisons2Degree(big.NewInt(7), big.NewInt(4), big.NewInt(3))

		for i := 0; i < len(slise); i++ {
			slise[i].Print()
		}
	} else {
		fmt.Println("Jacobi eror  (a/p)  = 1  (n/p)  =  -1 ")
	}

}
