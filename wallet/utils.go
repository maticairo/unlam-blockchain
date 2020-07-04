package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

/*derivado del base64, es introducido por blockchain*/

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}

// 0 O l I + / <- se sacaron estos caracteres de base64 porque se confunden facilmente.
