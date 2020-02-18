package smartzz

import (
	"strconv"	
	"github.com/ejemba/fizzbuzz/domain/zz"
)



// Smart Service is a top service that use other services to bring some intelligence.
//
// The first service will take precedence over the other
type SmartZzService interface {
	Execute(input int) string
}

func NewSmartZzService(services ...zz.ZzService ) SmartZzService {
	return &smartZzService{services : services}
}

// 
type smartZzService struct {	
	services []zz.ZzService
}

// Execute will return the first  
func (s *smartZzService) Execute(input int) string  {
	for _ , s := range s.services  {
		if input % s.Zz().Modulo() == 0 {
			return s.Zz().Replacement()
		}
	}
	return strconv.Itoa(input)
}
