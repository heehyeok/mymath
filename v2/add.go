// Package mymath provides generic math utilities.
package mymath

import "golang.org/x/exp/constraints"

//Number is a type constraint that allows integer and float types.
type Number interface {
	constraints.Integer | constraints.Float
}

//Add adds two numbers of type Number and returns the result.
//
//For more info on addition, see:
//https://www.mathsisfun.com/numbers/addition.html

func Add[T Number] (a, b T) T {
	return a + b
}
