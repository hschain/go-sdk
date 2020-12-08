package main

import (
	"encoding/json"
	"fmt"
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

	mnemonic := "drink lunch camera exhibit green spirit fiber mammal maze unable toilet hobby broken crop program physical village increase vapor jungle skirt section seed way"
	log.Printf("mnemonic is %s", mnemonic)

	wallet, err := hsc.FromMnemonic("hsc", mnemonic, "m/44'/532'/0'/0/0")
	if err != nil {
		log.Panicf("create wallet faied %s", err)
	}
	log.Printf("wallet is %+v", wallet)

	tx := hsc.TransactionPayload{
		Message: []json.RawMessage{
			json.RawMessage([]byte(fmt.Sprintf(`{"type":"cosmos-sdk/MsgSend","value":{"from_address":"%s","to_address":"%s","amount":[{"denom":"%s","amount":"10000000"}]}}`, wallet.Address, "hsc12503y3tnra7wlw2rg5htna3s0wdrtvfuws3mjm", "uhst"))),
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
