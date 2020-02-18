package leboncoin

import (
	"testing"
	"fmt"
	"github.com/ejemba/fizzbuzz/domain/zz"
)


func TestLeBonCoinService(t *testing.T) {
	
	z1 := zz.NewZz(3, "fizz")
	z2 := zz.NewZz(5, "buzz")
	
	
	var tests = []struct {
		limit int
		expected string
	} {
		{1,"1"},
		{2,"1,2"},
		{3,"1,2,fizz"},
		{4,"1,2,fizz,4"},
		{5,"1,2,fizz,4,buzz"},
		{6,"1,2,fizz,4,buzz,fizz"},
		{7,"1,2,fizz,4,buzz,fizz,7"},
		{8,"1,2,fizz,4,buzz,fizz,7,8"},
		{15,"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz"},
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%s", tt.limit, tt.expected)
        t.Run(testname, func(t *testing.T) {
			  lbcService := NewLeBonCoinService(z1 , z2, tt.limit)
			  
			  output := lbcService.Execute()
			  
            if output != tt.expected {
                t.Errorf("got %s, want %s", output, tt.expected)
            }
        })
    }
	
	
}
