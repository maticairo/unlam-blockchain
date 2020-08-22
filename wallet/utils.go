package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

/*
 Funciones derivadas de Base64 que fueron introducidas por blockchain.
 La principal diferencia radica en que los caracteres (0 O l I +) fueron quitados porque se confunden facilmente entre sí.
*/

//Base58Encode encodea en base58
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

//Base58Decode decodea en base58
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}
