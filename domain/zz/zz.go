package zz

import (
	"strconv"
	"fmt"
)

// The Zz interface represents the atomic action with FizzBuzz.
//
// Either Fizz,Buzz or FizzBuzz 
type Zz interface {
	Modulo() int
	Replacement() string	
}

func NewZz (modulo int, replacement string) Zz{
	return &zz{modulo: modulo,replacement: replacement}
}

func NewCompositeZz(z1 Zz, z2 Zz) Zz{
	return &zz{
		modulo: z1.Modulo() * z2.Modulo() ,
		replacement: fmt.Sprintf("%s%s" , z1.Replacement() , z2.Replacement()) ,
	}
}

type zz struct {
	modulo int
	replacement string
}

func (z *zz) Replacement() string {
	return z.replacement
}

func (z *zz) Modulo() int {
	return z.modulo
}


type ZzService interface {
	Execute(input int) string
	Zz() Zz
}

func NewZzService(z Zz) ZzService {
	return &zzService{z:z}
}

// 
type zzService struct {	
	z Zz
}

// Execute 
func (s *zzService) Execute(input int) string  {
	if input % s.z.Modulo() == 0 {
		return s.z.Replacement() 
	}
	return strconv.Itoa(input)
}

// Zz returns the internal Zz
func (s *zzService) Zz() Zz {
	return s.z
}
