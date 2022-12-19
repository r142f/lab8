package utils

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func AssertEqual[X comparable](x1 X, x2 X) {
	if x1 != x2 {
		panic(fmt.Sprintf("Utils.AssertEqual failed: %v != %v", x1, x2))
	}
}

func AssertLowerOrEqual[X constraints.Ordered](x1, x2 X) {
	if x1 > x2 {
		panic(fmt.Sprintf("Utils.AssertLowerOrEqual failed: %v > %v", x1, x2))
	}
}

func AssertGreater[X constraints.Ordered](x1, x2 X) {
	if x1 <= x2 {
		panic(fmt.Sprintf("Utils.AssertGreater failed: %v <= %v", x1, x2))
	}
}