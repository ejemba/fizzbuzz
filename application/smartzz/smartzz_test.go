package smartzz

import (
	"testing"
	"fmt"
	"github.com/ejemba/fizzbuzz/domain/zz"
)


func TestZzSmartService( t *testing.T) {
	z1 := zz.NewZz(3, "fizz")
	z2 := zz.NewZz(5, "buzz")

	z := zz.NewCompositeZz (z1, z2)

	zService3 := zz.NewZzService(z1)
	zService5 := zz.NewZzService(z2)
	zCompoService  := zz.NewZzService(z)
	
	zSmartService := NewSmartZzService(zCompoService , zService3 , zService5)

	
	var tests = []struct {
		input int
		expected string
	} {
		{1,"1"},
		{2,"2"},
		{3,"fizz"},
		{15,"fizzbuzz"},
		{4,"4"},
		{5,"buzz"},
		{30,"fizzbuzz"},
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%s", tt.input, tt.expected)
        t.Run(testname, func(t *testing.T) {
			  
			  output := zSmartService.Execute(tt.input)
			  
	
            if output != tt.expected {
                t.Errorf("got %s, want %s", output, tt.expected)
            }
        })
    }

}
