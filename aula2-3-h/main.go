package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println(contaDaSilvia.Credito(300))
	fmt.Println(contaDaSilvia.saldo)
}

func (c *ContaCorrente) Sacar(valor float64) string {
	saldo := c.saldo
	mensagem, verificaSaldo := c.VerificaSaldo(valor, saldo)

	if verificaSaldo == true {
		c.saldo -= valor
		return mensagem
	}
	return mensagem
}

func (c *ContaCorrente) VerificaSaldo(valorDoSaque float64, saldoCliente float64) (string, bool) {
	if saldoCliente > valorDoSaque && valorDoSaque > 0 {
		return "pode sacar", true
	}
	return "Saldo insuficiente", false
}

func (c *ContaCorrente) Credito(valor float64) string {
	mensagem, valorCredito := VerificaValorCreditado(valor)

	if valorCredito == true {
		c.saldo += valor
		return mensagem
	}
	return mensagem
}

func VerificaValorCreditado(valor float64) (string, bool) {
	verificaValor := valor

	if verificaValor < 0 {
		return "o valor depositado não pode ser menor que zero", false
	}
	return "o valor vai ser debitado", true
}
