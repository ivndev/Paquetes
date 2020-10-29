package operaciones

import "fmt"

//Suma func
func Suma() {
	var a, b int
	fmt.Println("---SUMA----")
	fmt.Println("ingrese primer numero ")
	fmt.Scanf("%d", &a)
	fmt.Println("ingrese segundo numero ")
	fmt.Scanf("%d", &b)
	fmt.Println("el resulñtado es:", a+b)
}

//Resta func
func Resta() {
	var i, l int
	fmt.Println("----RESTA----")
	fmt.Println("ingrese primer numero ")
	fmt.Scanf("%d\n", &i)
	fmt.Println("ingrese segundo numero ")
	fmt.Scanf("%d\n", &l)
	fmt.Println("el resulñtado es:", i-l)
}

//Division func
func Division() {
	var m, p int
	fmt.Println("----DIVISION-----")
	fmt.Println("\ningrese primer numero ")
	fmt.Scanf("%d\n", &m)
	fmt.Println("ingrese segundo numero ")
	fmt.Scanf("%d\n", &p)
	fmt.Println("el resultado es:", m/p)
}
