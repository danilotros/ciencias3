package main

import (
	"fmt"
	"strings"
	"strconv"
)
// arbol binario con valores enteros.
type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
}

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
	var N int
	N = 3
	opera := make([]string, N)
	resul := make([]int64, N)
	opera[0] = "+ 9 3 x"
	opera[1] = "+ - 5 2 1 y"
	opera[2] = "/ x y z"
	for i := 0; i < N; i++ {
    operaArray := strings.Split(opera[i]," ")
		t = insertar(t, operaArray[0], "")
		for j := 1; j < len(operaArray)/2; j++ {
	    t = insertar(t, operaArray[j], operaArray[len(operaArray)-j-1])
	  }
		resul[i] = OperarArbol(t)
		for j := 0; j < N; j++ {
	    strings.Replace(opera[j],operaArray[len(operaArray)-1],strconv.FormatInt((resul[i]), 10),len(operaArray))
	  }
		fmt.Println("Resultado:")
		fmt.Println(resul[N-1])
  }
	/*opera = "+ - * / 6 2 3 2 5"
	operaArray := strings.Split(opera," ")
	t = insertar(t, operaArray[0], "")
	for i := 1; i <= len(operaArray)/2; i++ {
    t = insertar(t, operaArray[i], operaArray[len(operaArray)-i])
  }
	fmt.Println("Inorden:")
	RecorrerInorden(t)
	fmt.Println()
	fmt.Println("Resultado:")
	fmt.Println(OperarArbol(t))*/
}
