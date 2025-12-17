package bigNumbers

import (
	"errors"
	"math/big"
)

type bigNumbers struct {
	a *big.Int
	b *big.Int
}

func NewBigNumbers(a, b *big.Int) *bigNumbers {
	return &bigNumbers{
		a: new(big.Int).Set(a),
		b: new(big.Int).Set(b),
	}
}

func (n *bigNumbers) Add() *big.Int {
	return new(big.Int).Add(n.a, n.b)
}

func (n *bigNumbers) Sub() *big.Int {
	return new(big.Int).Sub(n.a, n.b)
}

func (n *bigNumbers) Mul() *big.Int {
	return new(big.Int).Mul(n.a, n.b)
}

func (n *bigNumbers) DivRat() (*big.Rat, error) {
	if n.b.Sign() == 0 {
		return nil, errors.New("division by zero")
	}
	div := new(big.Rat).SetFrac(n.a, n.b)
	return div, nil
}
