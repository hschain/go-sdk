package hsc

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func (w Wallet) TransferAccounts(toAddress, asset, memo string, amount float64) (string, error) {

	strAmount := strconv.FormatInt(int64(amount*1000000), 10)
	tx := TransactionPayload{
		Message: []json.RawMessage{
			json.RawMessage([]byte(fmt.Sprintf(`{"type":"cosmos-sdk/MsgSend","value":{"from_address":"%s","to_address":"%s","amount":[{"denom":"u%s","amount":"%s"}]}}`, w.Address, toAddress, asset, strAmount))),
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

func (w Wallet) TransferDestory(asset, memo string, amount float64) (string, error) {

	strAmount := strconv.FormatInt(int64(amount*1000000), 10)
	tx := TransactionPayload{
		Message: []json.RawMessage{
			json.RawMessage([]byte(fmt.Sprintf(`{"type":"cosmos-sdk/MsgDestory","value":{"from_address":"%s","amount":[{"denom":"u%s","amount":"%s"}]}}`, w.Address, asset, strAmount))),
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
