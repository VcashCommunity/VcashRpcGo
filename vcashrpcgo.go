package vcashrpcgo

import (
	"fmt"
	"github.com/teambition/jsonrpc-go"
	"encoding/json"
	"net/http"
	"log"
	"strings"
	"io/ioutil"
)

type RpcPayload struct {
	id     int
	method string
	params []string
	url    string
}

type RpcCheckResponse struct {
	Amount       float64
	HouseAddress string
	UserAddress  string
	Received     bool
}

func RpcGetInfo() map[string]interface{} {
	// getinfo
	params := []string{}
	pay := RpcPayload{id: 1, method: "getinfo", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcGetBalance() map[string]interface{} {
	// getbalance
	params := []string{}
	pay := RpcPayload{id: 1, method: "getbalance", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcGetNewAddress() map[string]interface{} {
	// getnewaddress
	params := []string{}
	pay := RpcPayload{id: 1, method: "getnewaddress", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcListSinceBlock(hash string) map[string]interface{} {
	// listtransactions
	// ex: listsinceblock 0000000000009fe250400eed6cb80ecf09cd90642b0a12019d097c60ac462dcb
	params := []string{hash}
	pay := RpcPayload{id: 1, method: "listsinceblock", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcListTransactions(account string, count string, from string) map[string]interface{} {
	// listtransactions
	// ex: account = *, count = 80, from = 0
	params := []string{account, count, from}
	pay := RpcPayload{id: 1, method: "listtransactions", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcListReceivedByAddress() map[string]interface{} {
	// listreceivedbyaddress
	params := []string{}
	pay := RpcPayload{id: 1, method: "listreceivedbyaddress", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcGetTransaction(txid string) map[string]interface{} {
	// gettransaction
	params := []string{txid}
	pay := RpcPayload{id: 1, method: "gettransaction", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcValidateAddress(address string) map[string]interface{} {
	// validateaddress
	params := []string{address}
	pay := RpcPayload{id: 1, method: "validateaddress", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcGetBlockCount() map[string]interface{} {
	// getblockcount
	params := []string{}
	pay := RpcPayload{id: 1, method: "getblockcount", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcGetDifficulty() map[string]interface{} {
	// getdifficulty
	params := []string{}
	pay := RpcPayload{id: 1, method: "getdifficulty", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func RpcSendToAddress(address string, amount string) map[string]interface{} {
	// WARNING: USE WITH CAUTION
	// sendtoaddress
	params := []string{address, amount}
	pay := RpcPayload{id: 1, method: "sendtoaddress", params: params, url: "http://127.0.0.1:9195"}
	map_res := callRpc(pay)
	return map_res
}

func CheckReceived(address string) RpcCheckResponse {
	// Check if address has Received funds from user
	// Address is generated by RpcGetNewAddress() (New empty address)
	// Triple check: listreceivedbyaddress, listtransactions, gettransaction
	amount := 0.0
	var user_address string
	status_received := false
	// Check if address has Received funds in recent transaction
	response := RpcListReceivedByAddress()
	for _, u := range response["result"].([]interface{}) {
		// Check info about casting in golang
		uu := u.(map[string]interface{})
		if uu["address"] == address {
			// Address found, get amount
			amount = uu["amount"].(float64)
			break
		}
	}
	// Parse listtransactions
	if amount > 0 {
		// After the check we will have all needed data (HouseAddress, UserAddress, bet_amount)
		response = RpcListTransactions("*", "80", "0")
		for _, u := range response["result"].([]interface{}) {
			vv := u.(map[string]interface{})
			if vv["address"] == address {
				txid := vv["txid"].(string)
				txdata := RpcGetTransaction(txid)
				// Let's do some magic!
				vout := txdata["result"].(map[string]interface{})["vout"]
				scriptPubKey := vout.([]interface{})[0].(map[string]interface{})["scriptPubKey"]
				user_address = scriptPubKey.(map[string]interface{})["addresses"].([]interface{})[0].(string)
				break
			}
		}
	}
	if amount > 0 && user_address != "" {
		status_received = true
	}
	data := RpcCheckResponse{amount, address, user_address, status_received}
	return data
}

func callRpc(payload RpcPayload) map[string]interface{} {
	// Prepare jsonrpc 2.0 request
	response, _ := jsonrpc.Request(payload.id, payload.method, payload.params)
	s := string(response)

	res, err := http.Post(payload.url, "application/json", strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	var byte_tab []byte
	byte_tab, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data_map map[string]interface{}
	err = json.Unmarshal(byte_tab, &data_map)
	if err != nil {
		//error handling goes here
		fmt.Printf("err %s", err)
	}
	return data_map
}
