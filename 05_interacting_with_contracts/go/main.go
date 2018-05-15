package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	addr := common.HexToAddress("0x731a10897d267e19B34503aD902d0A29173Ba4B1")
	port := "10545"

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

	if count.Int64() > 0 {
		evtID, organizer, title, attendees, errE := att.GetEvent(&bind.CallOpts{}, big.NewInt(0))
		if errC != nil {
			log.Fatalf("GetEvent %v", errE)
		}
		fmt.Printf("eventid: %d\norganizer: %s\ntitle: %s\nattendees: %v\n", evtID, organizer.String(), title, attendees)
	}

}
