package main

import (
	"encoding/json"
	"log"

	"github.com/hschain/go-sdk/hsc"
)

func main() {
	/*
		mnemonic, err := sacco.GenerateMnemonic()
		if err != nil {
			log.Panicf("create mnmonic faied %s", err)
		}
	*/

	mnemonic := "hurt embark exclude harvest silly oval jar metal obscure they renew junk often artwork link situate fat town uncover bamboo monster federal pink debris"
	log.Printf("mnemonic is %s", mnemonic)

	wallet, err := hsc.FromMnemonic("hsc", mnemonic, "m/44'/532'/0'/0/0")
	if err != nil {
		log.Panicf("create wallet faied %s", err)
	}
	log.Printf("wallet is %+v", wallet)

	tx := hsc.TransactionPayload{
		Message: []json.RawMessage{
			json.RawMessage([]byte(`{"type":"cosmos-sdk/MsgSend","value":{"from_address":"hsc1tqkux4ck3uawakl452sy7rzjq7x3ersurlzp6z","to_address":"hsc13q8fvtemy0tt0skhj4td9haqjkad6ffhz28qdg","amount":[{"denom":"uhst","amount":"10"}]}}`)),
		},
		Fee: hsc.Fee{
			Amount: []hsc.Coin{},
			Gas:    "200000",
		},
	}

	stx, err := wallet.Sign(tx, "test", "0", "0")
	if err != nil {
		log.Panicf("sign faied %s", err)
	}

	log.Printf("stx is %+v", stx)

	out, err := wallet.SignAndBroadcast(tx, "http://119.28.60.149:1317", "async")
	if err != nil {
		log.Panicf("broadcast faied %s", err)
	}

	log.Printf("out is %s", out)
}
