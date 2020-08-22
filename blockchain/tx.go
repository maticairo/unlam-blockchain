package blockchain

import (
	"bytes"

	"github.com/maticairo/unlam-blockchain/wallet"
)

//TxOutput define el output de una transaccion
type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

//TxInput define el input de una transaccion
type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

//NewTXOutput genera un nuevo output para un address
func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

//UsesKey verifica si un hash de clave publica es igual a la clave publica del input de una TX
func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

//Lock bloquea el output para un address específico
func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

//IsLockedWithKey verifica si un output fue bloqueado con una key
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}
