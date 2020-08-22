package main

import (
	"os"

	"github.com/maticairo/unlam-blockchain/cli"
)

/*
	Inicialización del programa, todo comienza ejecutando la interfaz de linea de comandos.
*/
func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
