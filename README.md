# UNLaM Criptografía - Blockchain

## Objetivo

Este proyecto fue desarrollado por alumnos de la Universidad Nacional de La Matanza, en el marco de la cátedra de Criptografía, y pretende demostrar de manera simple la implementación de una Blockchain para validar y securizar transacciones entre Wallets.

La herramienta provee una interfaz de linea de comandos a través de la cual se puede simular y ver la generación de una blockchain.

### Comandos

| Integrantes | Parametros |
| ------ | ------ |
| createwallet | - |
| createblockchain | -ADDRESS |
| listaddresses | - |
| send | -FROM -TO -AMOUNT |
| printchain | - |
| getbalance | -ADDRESS |

### Ejemplos

#### createwallet

Input:

```sh
$ go run main.go createwallet
```

Output:
```sh
New address is: 1C6yMEC88e5UXPn59upTHQsmJCYAZE3SLy
```

#### createblockchain

Input:

```sh
$ go run main.go createblockchain -address 1C6yMEC88e5UXPn59upTHQsmJCYAZE3SLy
```

Output:
```sh
000337396060da20d4d5ae4584c64415df6779db937a450fb613cd749e29b79a
Genesis created
Finished!
```

#### printchain

Input:

```sh
$ go run main.go printchain
```

Output:
```sh
Prev. hash:
Hash: 000337396060da20d4d5ae4584c64415df6779db937a450fb613cd749e29b79a
PoW: true
--- Transaction 1a8a316b7e74eeb859b143aaa4f42b52206e44cab9f0b8e396c8f7246c73361b:
     Input 0:
       TXID:
       Out:       -1
       Signature:
       PubKey:    4669727374205472616e73616374696f6e2066726f6d2047656e65736973
     Output 0:
       Value:  100
       Script: 79ca80415dac51ee07cfae8fdc243bd5f0df38a4
```

#### send

Input:

```sh
$ go run main.go send -from 1C6yMEC88e5UXPn59upTHQsmJCYAZE3SLy -to 1MFkBG8sJjJE8hBww5U5HNDV6BW8YB9bXC -amount 30
```

Output:
```sh
000fd826f4c05d0490d1ccc43a28248bae12554febadb0b0ca67ec200aa4ecda
Success!
```

#### getbalance

Input:

```sh
$ go run main.go getbalance -address 1C6yMEC88e5UXPn59upTHQsmJCYAZE3SLy
```

Output:
```sh
Balance of 1C6yMEC88e5UXPn59upTHQsmJCYAZE3SLy: 70
```

#### listaddresses

Input:

```sh
$ go run main.go listaddresses
```

Output:
```sh
1MFkBG8sJjJE8hBww5U5HNDV6BW8YB9bXC
1C6yMEC88e5UXPn59upTHQsmJCYAZE3SLy
```

### Equipo

Dillinger is currently extended with the following plugins. Instructions on how to use them in your own application are linked below.

| Integrantes |
| ------ |
| Cairo, Matías |
| De Rito, Micaela |
| Foglia, Julieta |
| Reynoso, Thomas |
| Perez, Lautaro |
