package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	addr := common.HexToAddress("")
	port := ""

	cli, err := ethclient.Dial(
		"http://localhost:" + port,
	)
	if err != nil {
		log.Fatalf("rpc.Dial %v", err)
	}
	att, errA := NewEvtAttendance(addr, cli)
	if err != nil {
		log.Fatalf("NewAttendance: %v", errA)
	}

	count, errC := att.GetEventCount(&bind.CallOpts{})
	if errC != nil {
		log.Fatalf("GetEventCount %v", errC)
	}
	fmt.Println(count)

}
