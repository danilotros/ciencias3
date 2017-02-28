package main

import (
	"fmt"
	"strings"
	"strconv"
	//"math/rand"
)
// arbol binario con valores enteros.
type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
}
/*
func New(n, k int) *Arbol {
	var t *Arbol
	for _, v := range rand.Perm(n) {
		t = insertar(t, (1+v)*k)
	}
	return t
}
*/
func insertar(t *Arbol, v string, z string) *Arbol {
	if t == nil {
		return &Arbol{nil, v, nil}
	}
	if t.Izquierda == nil {
		t.Izquierda = insertar(t.Izquierda, v , z)
		t.Derecha = insertar(t.Derecha, z, v)
		return t
	}else{
		t.Izquierda = insertar(t.Izquierda, v , z)
		return t
	}
	return t
}

func RecorrerInorden(t *Arbol) {
	if t == nil {
		return
	}
	RecorrerInorden(t.Izquierda)
	fmt.Print(t.Valor)
	fmt.Print(" ")
	RecorrerInorden(t.Derecha)
}

func OperarArbol(t *Arbol) int64{
	if t == nil {
		return 0
	}
	if t.Valor == "+"{
		return OperarArbol(t.Izquierda) + OperarArbol(t.Derecha)
	}else if t.Valor == "-"{
		return OperarArbol(t.Izquierda) - OperarArbol(t.Derecha)
	}else if t.Valor == "*"{
		return OperarArbol(t.Izquierda) * OperarArbol(t.Derecha)
	}else if t.Valor == "/"{
		return OperarArbol(t.Izquierda) / OperarArbol(t.Derecha)
	}
	num, _ := strconv.ParseInt(t.Valor, 0, 64)
	return num
}

func main() {
  var t *Arbol
	var opera string
	opera = "+ - * / 6 2 3 2 5"
	operaArray := strings.Split(opera," ")
	t = insertar(t, operaArray[0], "")
	for i := 1; i <= len(operaArray)/2; i++ {
    t = insertar(t, operaArray[i], operaArray[len(operaArray)-i])
  }
	fmt.Println("Inorden:")
	RecorrerInorden(t)
	fmt.Println()
	fmt.Println("Resultado:")
	fmt.Println(OperarArbol(t))
}
