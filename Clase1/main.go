package main

import "fmt"

//interfaces
func main(){
	var e EjemploInterface = &Estructura{} // Ahora e es una interfaz
	fmt.Println(e.miFuncion())  
	fmt.Println(e.Calculo(2, 10)) 	
}

type EjemploInterface interface{
	miFuncion() string
	Calculo(n1 int, n2 int) int

}

type Estructura struct{}

func (*Estructura) Calculo(n1 int, n2 int) int{
	resultado := n1 + n2
	return resultado
}

func (*Estructura) miFuncion() string{
	return "texto de la funcion"
}