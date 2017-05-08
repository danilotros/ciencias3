package main

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
	"bufio"
	"os"
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

func insertarExp(t *Arbol, v string, z string) *Arbol {
	if t == nil {
		return &Arbol{nil, TablaSimbolos(v), nil}
	}
	if t.Izquierda == nil {
		t.Izquierda = insertar(t.Izquierda, TablaSimbolos(v) , TablaSimbolos(z))
		t.Derecha = insertar(t.Derecha, TablaSimbolos(z), TablaSimbolos(v))
		return t
	}else{
		t.Izquierda = insertar(t.Izquierda, TablaSimbolos(v) , TablaSimbolos(z))
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

func InordenValores(t *Arbol, te *Arbol) {
	if t == nil {
		return
	}
	InordenValores(t.Izquierda, te.Izquierda)
	if te.Valor == "Valor" {
		fmt.Print("Valor : ")
		fmt.Println(t.Valor)
	}
	if te.Valor == "Variable" {
		fmt.Print("Variable : ")
		fmt.Println(t.Valor)
	}
	if te.Valor == "Expresión" {
		fmt.Print("Expresión : ")
		RecorrerInorden(t)
		fmt.Println("")
	}
	InordenValores(t.Derecha, te.Derecha)
}

func TablaSimbolos(raiz string) string {
	valo, _ := regexp.MatchString("[0-9,true,false]",raiz)
	if valo == true {
		return "Valor"
	}
	vari, _ := regexp.MatchString("[a-z]",raiz)
	if vari == true {
		return "Variable"
	}
	expr, _ := regexp.MatchString("[<,>,=,+,-,*,/,&,|,!]",raiz)
	if expr == true {
		return "Expresión"
	}
	return "-"
}

func Remplazar (va string, t *Arbol, tp *Arbol, v *Arbol, vp *Arbol) {
	if t == nil {
		return
	}
	if t.Izquierda == nil && t.Derecha == nil {
		return
	}
	if t.Izquierda.Valor == va {
		t.Izquierda = tp
		v.Izquierda = vp
		Remplazar(va,t.Derecha,tp,v.Derecha,vp)
	}else if t.Derecha.Valor == va {
		t.Derecha = tp
		v.Derecha = vp
		Remplazar(va,t.Izquierda,tp,v.Izquierda,vp)
	}else{
		Remplazar(va,t.Derecha,tp,v.Derecha,vp)
		Remplazar(va,t.Izquierda,tp,v.Izquierda,vp)
	}
}

func main() {
	var Nec string
	var N int
	fmt.Println("Ingrese # de ecuaciones:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Nec = scanner.Text()
	Nue, _ := strconv.ParseInt(Nec, 0, 64)
	N = int(Nue)
	opera := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Println("Ingrese ecuacion:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		opera[i] = scanner.Text()
	}
	fmt.Println("Arboles:")
	arbol := make([]*Arbol, N)
	arbole := make([]*Arbol, N)
	for i := 0; i < N; i++ {
		var t *Arbol
		var te *Arbol
    operaArray := strings.Split(opera[i]," ")
		t = insertar(t, operaArray[0], "")
		te = insertarExp(te, operaArray[0], "")
		for i := 1; i <= len(operaArray)/2; i++ {
	    t = insertar(t, operaArray[i], operaArray[len(operaArray)-i])
			te = insertarExp(te, operaArray[i], operaArray[len(operaArray)-i])
	  }
		arbol[i] = t
		arbole[i] = te
		/*fmt.Print((i + 1))
		fmt.Println(":")
		RecorrerInorden(resul[i])
		fmt.Println("")
		RecorrerInorden(te)
		fmt.Println("")
		InordenValores(t,te)*/
  }
	for i := 0; i < N-1 ; i++ {
		var v string
		operaArray := strings.Split(opera[i]," ")
		v = operaArray[len(operaArray)-1]
		for j := i + 1; j < N; j++ {
			Remplazar(v,arbol[j],arbol[i].Izquierda,arbole[j],arbole[i].Izquierda)
		}
	}
  InordenValores(arbol[N-1],arbole[N-1])
}
