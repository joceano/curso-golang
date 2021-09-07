package main

import (
	"fmt"

	"github.com/cod3rcursos/goarea"
	"github.com/joceano/curso-golang/pacote/liblocal"
)

func main() {
	// lib importada do github
	fmt.Println(goarea.Circ(6.0))
	fmt.Println(goarea.Rect(5.0, 2.0))

	// fmt.Println(area._TrianguloEq(5.0, 2.0))

	//lib importada localmente
	fmt.Println(liblocal.Teste())
}
