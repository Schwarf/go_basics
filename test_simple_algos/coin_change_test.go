package test_simple_algos

import (
	"github.com/Schwarf/go_basics/simple_algos"
	"testing"
)

func TestCoinChangeTwo(t *testing.T) {
	coins := make([]int, 3)
	coins[0] = 1
	coins[1] = 2
	coins[2] = 5
	target := 10
	if result := simple_algos.MinimizeCoinChange(coins, target); result != 2 {
		t.Errorf("Expected result is 2 but is %d", result)
	}
}

func TestCoinChangeZero(t *testing.T) {
	coins := make([]int, 3)
	coins[0] = 1
	coins[1] = 2
	coins[2] = 5
	target := 0
	if result := simple_algos.MinimizeCoinChange(coins, target); result != 0 {
		t.Errorf("Expected result is 0 but is %d", result)
	}
}

func TestCoinChangeFour(t *testing.T) {
	coins := make([]int, 3)
	coins[0] = 1
	coins[1] = 2
	coins[2] = 3
	target := 10
	if result := simple_algos.MinimizeCoinChange(coins, target); result != 4 {
		t.Errorf("Expected result is 4 but is %d", result)
	}
}
