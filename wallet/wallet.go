package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4          //bytes
	version        = byte(0x00) //0 - donde reside el address en nuestra blockchain
)

//Wallet es la estructura que se compone de una clave pública y privada
type Wallet struct {
	PrivateKey ecdsa.PrivateKey //Elliptical Curve Digital Sign Algorithm
	PublicKey  []byte
}

//Address genera el hash de la clave pública, anexa esos bytes con la version, genera el checksum, obtiene el hash completo en base al hash versionado y checksum, y finalmente lo encodea en base58.
func (w Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)

	versionedHash := append([]byte{version}, pubHash...)
	checksum := Checksum(versionedHash)

	fullHash := append(versionedHash, checksum...)
	address := Base58Encode(fullHash)

	return address
}

//NewKeyPair obtiene las claves públicas y privadas a través del método de curvas elípticas
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256() //256 bytes output

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

//MakeWallet genera un nuevo par de claves pública y privada e instancia una wallet
func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

//PublicKeyHash usa el algoritmo ripemd160 para generar el hash de la clave pública
func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

//Checksum es importante para verificar y firmar transacciones
func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}

/*ValidateAddress valida un address comparando los checksums del hash y el target*/
func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-checksumLength:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-checksumLength]
	targetChecksum := Checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}
