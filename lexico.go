package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// arbol binario con valores enteros.
type Arbol struct {
	Izquierda *Arbol
	Valor     string
	Derecha   *Arbol
}

func insertar(t *Arbol, v string, z string) *Arbol {
	if t == nil {
		return &Arbol{nil, v, nil}
	}
	if t.Izquierda == nil {
		t.Izquierda = insertar(t.Izquierda, v, z)
		t.Derecha = insertar(t.Derecha, z, v)
		return t
	} else {
		t.Izquierda = insertar(t.Izquierda, v, z)
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

func OperarArbol(t *Arbol) int64 {
	if t == nil {
		return 0
	}
	if t.Valor == "+" {
		return OperarArbol(t.Izquierda) + OperarArbol(t.Derecha)
	} else if t.Valor == "-" {
		return OperarArbol(t.Izquierda) - OperarArbol(t.Derecha)
	} else if t.Valor == "*" {
		return OperarArbol(t.Izquierda) * OperarArbol(t.Derecha)
	} else if t.Valor == "/" {
		return OperarArbol(t.Izquierda) / OperarArbol(t.Derecha)
	}
	num, _ := strconv.ParseInt(t.Valor, 0, 64)
	return num
}

func ingresarDatos()  {
	fmt.Println("ingrese ecuacion cantidad de ecuaciones")
	cantidad_ecuaciones := ""

	fmt.Scanln(&cantidad_ecuaciones)

	valornumerico, _ := strconv.Atoi(cantidad_ecuaciones)
   revision(valornumerico)


}



//revision(valornumerico,ecuacion)



func revision(valor int)  {

	ecuacion :=make([]string, valor)

	for i:=0; i<valor;i++{
		fmt.Println("ingrese ecuacion")
		fmt.Scanln(&ecuacion[i])
		r, _ := regexp.Compile("([-=!*/^+0123456789abcdefghijklmnopqrstuvwxyz])")

		cadena_separada := strings.Split(ecuacion[i], "")

		for i := 0; i < len(cadena_separada); i++ {

			varlor_nescesario := r.MatchString(cadena_separada[i])

			compila, _ := regexp.Compile("^:=$")
			var verificacion bool
			if varlor_nescesario == true {


			}

			if varlor_nescesario == false {

				if cadena_separada[i] == ":" {
					concatenacion := cadena_separada[i]
					concatenacion += cadena_separada[i+1]
					verificacion = compila.MatchString(concatenacion)

				}

				if verificacion == false {
					fmt.Println("error: ",cadena_separada[i])


				}

			}

		}

}
ingresarEcuaciones(valor,ecuacion)
}
func ingresarEcuaciones(numeroperaciones int,opera []string) {


	resul := make([]int64, numeroperaciones)

	for i := 0; i < numeroperaciones; i++ {
		var t *Arbol
		var v string
		operaArray := strings.Split(opera[i],"")
		v = operaArray[len(operaArray)-1]
		t = insertar(t, operaArray[0], "")
		for j := 1; j < len(operaArray)/2; j++ {
			t = insertar(t, operaArray[j], operaArray[len(operaArray)-j-1])
		}
		resul[i] = OperarArbol(t)
		RecorrerInorden(t)
		fmt.Println("")
		fmt.Print(v)
		fmt.Print(" = ")
		fmt.Println(resul[i])
		for j := 0; j < numeroperaciones; j++ {
			operaArray2 := strings.Split(opera[j], " ")
			for k := 0; k < len(operaArray2); k++ {
				var inc string
				inc = operaArray2[k]
				if inc == v {
					operaArray2[k] = strconv.Itoa(int(resul[i]))
				}
			}
			opera[j] = ""
			for k := 0; k < len(operaArray2)-1; k++ {
				opera[j] += operaArray2[k]
				opera[j] += " "
			}
			opera[j] += operaArray2[len(operaArray2)-1]
		}
	}
}

func main() {
	ingresarDatos()
}
