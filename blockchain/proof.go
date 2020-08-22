package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

/*
	Procedimiento:
	 - Tomar la data del bloque,
	 - Crear un contador (nonce) que empiece en 0,
	 - Crear hash de los datos más el contador,
	 - Chequear si el hash cumple con los requerimientos

	Requerimientos:
	 - Los primeros bytes deben contener 0s
*/

//Difficulty setea la dificultad de la PoW
const Difficulty = 12

//ProofOfWork es la estructura que define a la PoW
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

//NewProof instancia una nueva PoW
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

//InitData inicializa los datos a partir del nonce y cada bloque
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

//Run ejecuta la PoW (por cada nonce, compara con el target de la PoW hasta cumplir con el requerimiento)
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return nonce, hash[:]
}

//Validate valida que la PoW sea correcta
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

//ToHex es una funcion auxiliar para convertir un int a bytes
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}
