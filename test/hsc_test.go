package main

import (
	"fmt"
	"log"

	"github.com/hschain/go-sdk/hsc"
)

func main() {
	hschain := hsc.NewHsc("https://testnet.hschain.io/api/lcd")

	mnemonic := "drink lunch camera exhibit green spirit fiber mammal maze unable toilet hobby broken crop program physical village increase vapor jungle skirt section seed way"
	wallet, err := hsc.FromMnemonic("hsc", mnemonic, "m/44'/532'/0'/0/0", hschain)
	if err != nil {
		log.Panicf("create wallet faied %s", err)
	}
	log.Printf("wallet is %+v", wallet)

	// txHash, err := wallet.TransferAccounts("hsc13xg0642wy23hhe4u4gvqt7lyuy7lj5pyaj58eu", "hst", 1000)
	// log.Printf("out is %s", txHash)

	txHash, err := wallet.TransferDestory(1000, "hst", "")

	//tx, err := wallet.Hsc.GetTxHashTxInfo(txHash)

	if err != nil {
		log.Printf("tx hash is error : %+v", err)
	}

	fmt.Printf("%+v", txHash)
}
