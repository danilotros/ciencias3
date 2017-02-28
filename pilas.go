  package main

import (
  "fmt"
  "strconv"
)

type Nodo struct{
nombre string
cedula int
eps string
hora string


}






func (n *Nodo) String() string {
	return fmt.Sprint(n.cedula, "->" ,n.nombre,"->",n.eps,"->",n.hora)
}

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodos []*Nodo
	 contador int
}

// Push adds a Nodo to the stack.
func (s *Stack) Push(n *Nodo) {
	s.nodos = append(s.nodos[:s.contador], n)
	s.contador++
}

func (s *Stack) Pop() *Nodo {
	if s.contador == 0 {
		return nil
	}
	s.contador--
	return s.nodos[s.contador]
}



func llenarDatos(numero_pacientes int)[50]Nodo{
var nodo [50]Nodo
for i := 0;  i < numero_pacientes; i++ {


fmt.Println("Ingrese nombre")
nombre_paciente:=""
fmt.Scanln(&nombre_paciente)
nodo[i].nombre=nombre_paciente

fmt.Println("Ingrese cedula")
cedula_paciente:=""
fmt.Scanln(&cedula_paciente)
nodo[i].cedula,_=strconv.Atoi(cedula_paciente)

fmt.Println("Ingrese eps")
eps_paciente:=""
fmt.Scanln(&eps_paciente)
nodo[i].eps=eps_paciente

fmt.Println("Ingrese hora")
hora_paciente:=""
fmt.Scanln(&hora_paciente)
nodo[i].hora=hora_paciente



}
return nodo

}


func main() {
fmt.Println("ingresar numero de pasientes")
pacientes := ""
fmt.Scanln(&pacientes)
numero_pacientes,_:= strconv.Atoi(pacientes)
var nodo [50]Nodo
var pila Stack
//pilar:=NewStack()

nodo=llenarDatos(numero_pacientes)
for i:= 0;  i< numero_pacientes; i++ {
  pila.Push(&nodo[i])

}
for j := 0; j <= pila.contador; j++ {
  fmt.Println(pila.Pop())
}




}
