package zz

import (
	"testing"
	"fmt"
)


func TestZzSimple( t *testing.T) {
	z := NewZz(3, "fizz")
	
	var tests = []struct {
		input int
		expected string
	} {
		{1,"1"},
		{2,"2"},
		{3,"fizz"},
		{4,"4"},
		{5,"5"},
		{6,"fizz"},
	}

	for _, tt := range tests {

        testname := fmt.Sprintf("%d,%s", tt.input, tt.expected)
        t.Run(testname, func(t *testing.T) {
			  zService := NewZzService(z)
			  output := zService.Execute(tt.input)
	
            if output != tt.expected {
                t.Errorf("got %s, want %s", output, tt.expected)
            }
        })
    }

}


func TestZzComposite( t *testing.T) {
	z1 := NewZz(3, "fizz")
	z2 := NewZz(5, "buzz")

	z := NewCompositeZz(z1, z2)
	
	var tests = []struct {
		input int
		expected string
	} {
		{1,"1"},
		{2,"2"},
		{15,"fizzbuzz"},
		{4,"4"},
		{5,"5"},
		{30,"fizzbuzz"},
	}

	for _, tt := range tests {

        testname := fmt.Sprintf("%d,%s", tt.input, tt.expected)
        t.Run(testname, func(t *testing.T) {
			  zService := NewZzService(z)
			  output := zService.Execute(tt.input)
	
            if output != tt.expected {
                t.Errorf("got %s, want %s", output, tt.expected)
            }
        })
    }

}
