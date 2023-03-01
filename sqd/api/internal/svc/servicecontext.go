package svc

import (
	"YxEmr/sqd/api/internal/config"
	"YxEmr/sqd/rpc/add/adder"
	"YxEmr/sqd/rpc/cha/chaer"
	"YxEmr/sqd/rpc/del/deler"
	"YxEmr/sqd/rpc/per/perer"
	"YxEmr/sqd/rpc/reg/reger"
	"YxEmr/sqd/rpc/rep/reper"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Adder  adder.Adder
	Deler  deler.Deler
	Reger  reger.Reger
	Chaer  chaer.Chaer
	Perer  perer.Perer
	Reper  reper.Reper
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Adder:  adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Deler:  deler.NewDeler(zrpc.MustNewClient(c.Del)),
		Reger:  reger.NewReger(zrpc.MustNewClient(c.Reg)),
		Chaer:  chaer.NewChaer(zrpc.MustNewClient(c.Cha)),
		Perer:  perer.NewPerer(zrpc.MustNewClient(c.Per)),
		Reper:  reper.NewReper(zrpc.MustNewClient(c.Rep)),
	}
}
