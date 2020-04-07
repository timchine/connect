package connect

import (
	"log"
)

type Connector interface {
	Connect() error
	Close() error
}

var (
	connErrF = "连接外部程序出错%s"
)

type ExternalProcedure struct {
	connects []Connector
}

func NewExternalProcedure(connects ...Connector) *ExternalProcedure {
	e := &ExternalProcedure{}
	for _, conner := range connects {
		err := conner.Connect()
		if err != nil {
			log.Printf(connErrF, err)
		}
		e.connects = append(e.connects, conner)
	}
	return e
}

func (e ExternalProcedure) Close()  {
	for _, connector := range e.connects {
		if connector != nil {
			_ = connector.Close()
		}
	}
}