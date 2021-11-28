package main

import "fmt"


func main(){
	contaHenry := ContaCorrente{
		titular:       "Henry",
		numeroAgencia: 12,
		numeroConta:   1212,
		saldo:         100000,
	}

	contaTeste := ContaCorrente{
		titular:       "Teste",
		numeroAgencia: 1221,
		numeroConta:   1212,
		saldo:         1212,
	}
	fmt.Println(contaTeste, contaHenry)
}

