package main

import (
	"fmt"

	"github.com/okx/go-wallet-sdk/coins/polkadot"
)

func main() {
	// address()
	tx()
}

func address() {
	pk := []byte{91, 218, 158, 183, 240, 200, 28, 9, 5, 118, 36, 164, 104, 123, 187, 188, 151, 243, 200, 129, 229, 157, 9, 201, 251, 21, 39, 222, 114, 180, 51, 88}
	address, _ := polkadot.PubKeyToAddress(pk, polkadot.SubstratePrefix)
	fmt.Println(address)
}

func tx() {
	tx := polkadot.TxStruct{
		From:         "5E99A91wzW6rBio1eE6vxtqn2j2pCoGazGwxM1FTEbCEMQRX",
		To:           "5DoW9HHuqSfpf55Ux5pLdJbHFWvbngeg8Ynhub9DrdtxmZeV",
		Amount:       500000000000,
		Nonce:        0,
		Tip:          1000000000,
		BlockHeight:  8215485,
		BlockHash:    "0615071566c93df347d71e05b75ed360c8d99ef8a28334ebd50d6d2f337ff87c",
		GenesisHash:  "6408de7737c59c238890533af25896a2c20608d8b380bb01029acb392781063e",
		SpecVersion:  104000,
		TxVersion:    24,
		ModuleMethod: "a40403",
		Version:      "01",
	}

	signed, _ := polkadot.SignTx(tx, polkadot.Transfer, "e259ae61f2d82713167be5a367147aa0e608569fe0ae1cb04daea39a90b8c8cd")

	fmt.Print("signed payload = ")
	fmt.Println(signed)

	//			  rust transfer   (acala): a80a00 00 dc621b10081b4b51335553ef8df227feb0327649d00beab6e09c10a1dce973590b00407a10f3 5a24 010000dc07000001000000fc41b9bd8ef8fe53d58c7ea67c794c7ec9a73daf05e6d54b14ff6342c99ba64c5cfeb3e46c080274613bdb80809a3e84fe782ac31ea91e2c778de996f738e620
	// 			  rust transfer (westend): a40400 00 4ce05abd387b560855a3d486eba6237b9a08c6e9dfe351302a5ceda90be801fe070088526a74 dfbb 0002286bee40960100180000006408de7737c59c238890533af25896a2c20608d8b380bb01029acb392781063e0615071566c93df347d71e05b75ed360c8d99ef8a28334ebd50d6d2f337ff87c
	// rust transfer keep alive (westend): a40403 00 4ce05abd387b560855a3d486eba6237b9a08c6e9dfe351302a5ceda90be801fe070088526a74 dfbb 0002286bee40960100180000006408de7737c59c238890533af25896a2c20608d8b380bb01029acb392781063e0615071566c93df347d71e05b75ed360c8d99ef8a28334ebd50d6d2f337ff87c
	//   					           go: 0500   00 4ce05abd387b560855a3d486eba6237b9a08c6e9dfe351302a5ceda90be801fe070088526a74 d503 0002286bee40960100180000006408de7737c59c238890533af25896a2c20608d8b380bb01029acb392781063e0615071566c93df347d71e05b75ed360c8d99ef8a28334ebd50d6d2f337ff87c
}
