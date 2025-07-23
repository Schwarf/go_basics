package test_simple_algos

import (
	"github.com/Schwarf/go_basics/simple_algos"
	"testing"
)

func TestLongestCommonSubsequenceEmpty(t *testing.T) {
	s1 := ""
	s2 := ""
	if size := simple_algos.LongestCommonSubsequence(s1, s2); size != 0 {
		t.Errorf("Expected size is 0 but is %d", size)
	}
}

func TestLongestCommonSubsequenceSameFirst2(t *testing.T) {
	s1 := "abc"
	s2 := "abd"
	if size := simple_algos.LongestCommonSubsequence(s1, s2); size != 2 {
		t.Errorf("Expected size is 2 but is %d", size)
	}
}

func TestLongestCommonSubsequenceDistributed3(t *testing.T) {
	s1 := "abec"
	s2 := "abdfc"
	if size := simple_algos.LongestCommonSubsequence(s1, s2); size != 3 {
		t.Errorf("Expected size is 3 but is %d", size)
	}
}

func TestLongestCommonSubsequenceDistributed4(t *testing.T) {
	s1 := "561abec"
	s2 := "8912abdfc"
	if size := simple_algos.LongestCommonSubsequence(s1, s2); size != 4 {
		t.Errorf("Expected size is 3 but is %d", size)
	}
}
