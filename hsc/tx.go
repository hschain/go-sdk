package hsc

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func (w Wallet) TransferAccounts(toAddress string, amount float64, denom string, memo string) (string, error) {

	strAmount := strconv.FormatInt(int64(amount*1000000), 10)
	tx := TransactionPayload{
		Message: []json.RawMessage{
			json.RawMessage([]byte(fmt.Sprintf(`{"type":"cosmos-sdk/MsgSend","value":{"from_address":"%s","to_address":"%s","amount":[{"denom":"u%s","amount":"%s"}]}}`, w.Address, toAddress, denom, strAmount))),
		},
		Fee: Fee{
			Amount: []Coin{},
			Gas:    "200000",
		},
		Memo: memo,
	}

	txHash, err := w.SignAndBroadcast(tx, "async")
	if err != nil {
		log.Printf("SignAndBroadcast is err:%+v", err)
		return "", err
	}
	return txHash, nil
}

func (w Wallet) TransferDestory(amount float64, denom string, memo string) (string, error) {

	strAmount := strconv.FormatInt(int64(amount*1000000), 10)
	tx := TransactionPayload{
		Message: []json.RawMessage{
			json.RawMessage([]byte(fmt.Sprintf(`{"type":"cosmos-sdk/MsgDestory","value":{"from_address":"%s","amount":[{"denom":"u%s","amount":"%s"}]}}`, w.Address, denom, strAmount))),
		},
		Fee: Fee{
			Amount: []Coin{},
			Gas:    "200000",
		},
		Memo: memo,
	}

	txHash, err := w.SignAndBroadcast(tx, "async")
	if err != nil {
		log.Printf("SignAndBroadcast is err:%+v", err)
		return "", err
	}
	return txHash, nil
}
