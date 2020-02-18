package leboncoin

import (
	"strings"

	"github.com/ejemba/fizzbuzz/application/smartzz"
	"github.com/ejemba/fizzbuzz/domain/zz"
)

// NewLeBonCoinService constructs a new LBC Le Bon Coin application service
func NewLeBonCoinService(z1 zz.Zz, z2 zz.Zz, limit int) LeBonCoinService {
	zService1 := zz.NewZzService(z1)
	zService2 := zz.NewZzService(z2)

	zComposite := zz.NewCompositeZz(z1, z2)
	zCompositeService := zz.NewZzService(zComposite)

	zSmartService := smartzz.NewSmartZzService(zCompositeService, zService1, zService2)

	return &lbcService{zSmartService: zSmartService, limit: limit}
}

// LeBonCoinService is a hig level (Process) service . 
type LeBonCoinService interface {
	Execute() string
}

type lbcService struct {
	zSmartService smartzz.SmartZzService
	limit         int
}

func (s *lbcService) Execute() string {
	msg := make([]string,0)
	for i := 1; i <= s.limit; i++ {
		msg = append(msg, s.zSmartService.Execute(i))
	}

	return strings.Join(msg, ",")
}
