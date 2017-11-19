package main

import (
	"fmt"
	"time"
	xvc "github.com/devmahno/vcashrpcgo"
)

func main() {
	fmt.Printf("\n* Call rpc getinfo\n")
	show_data(xvc.RpcGetInfo())

	fmt.Printf("\n* Call rpc getbalance\n")
	show_data(xvc.RpcGetBalance())

	//show_data(xvc.RpcListTransactions("*", "10", "0"))

	start := time.Now()

	fmt.Printf("* check_received \n")
	fmt.Printf("%+v\n", xvc.CheckReceived("SOME_XVC_ADDRESS"))

	elapsed := time.Since(start)
	fmt.Printf("check_received took %s\n", elapsed)

	//ticker1 := time.NewTicker(2 * time.Second)
	//quit := make(chan struct{})
	//for {
	//	select {
	//	case tick := <- ticker1.C:
	//		// do stuff
	//		fmt.Printf("Tick %s\n", tick)
	//	case <- quit:
	//		ticker1.Stop()
	//		return
	//	}
	//}

	ticker := time.NewTicker(time.Second * 5)
	//go func() {
	//	for t := range ticker.C {
	//		fmt.Println("Tick at", t)
	//		fmt.Printf("* getdifficulty \n")
	//		show_data(xvc.RpcGetDifficulty())
	//		fmt.Printf("\n")
	//	}
	//}()
	time.Sleep(time.Second * 2)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func show_data(data map[string]interface{}) {
	for k, v := range data {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println("List", i, u)
			}
		case map[string]interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println("Map", i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
