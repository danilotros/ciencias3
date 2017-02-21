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



func llenarDatos(nodo Nodo,pila Stack ) Stack{


fmt.Println("Ingrese nombre")
nombre_paciente:=""
fmt.Scanln(&nombre_paciente)
nodo.nombre=nombre_paciente

fmt.Println("Ingrese cedula")
cedula_paciente:=""
fmt.Scanln(&cedula_paciente)
nodo.cedula,_=strconv.Atoi(cedula_paciente)

fmt.Println("Ingrese eps")
eps_paciente:=""
fmt.Scanln(&eps_paciente)
nodo.eps=eps_paciente

fmt.Println("Ingrese hora")
hora_paciente:=""
fmt.Scanln(&hora_paciente)
nodo.hora=hora_paciente
pila.Push(&nodo)


return pila


}


func main() {
  var nodo Nodo
  var pila Stack

for{
fmt.Println("Desea ingresar Pasiente escriba 1","\n","Desea imprimir lista de pasientes ingresados digite  2","Desea salir digite 3")
opcion := ""
fmt.Scanln(&opcion)
switch opcion{
case "1":
fmt.Println("ingresar numero de pasientes a ingresar")
pacientes := ""
fmt.Scanln(&pacientes)
numero_pacientes,_:= strconv.Atoi(pacientes)

for i:=0;i<numero_pacientes;i++{
pila=llenarDatos(nodo,pila)

}

case "2":
for j:=0;j<=pila.contador;j++{
fmt.Println(pila.Pop())
  }
case "3":
  break
  break
}

}


}
