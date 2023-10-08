package main

import (
	"github.com/vinicius-n4/alura-golang-course/alura-bank/accounts"
	"github.com/vinicius-n4/alura-golang-course/alura-bank/clients"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteVinicius := clients.Titular{"Vinicius", "123.123.123.12", "Desenvolvedor Go"}
	contaDoVinicius := accounts.ContaCorrente{Titular: clienteVinicius, NumeroAgencia: 123, NumeroConta: 1234}
	poupancaDoVinicius := accounts.ContaPoupanca{Titular: clienteVinicius, NumeroAgencia: 123, NumeroConta: 112233, Operacao: 1}

	fmt.Println(contaDoVinicius)
	fmt.Println(poupancaDoVinicius)

	contaDoVinicius.Depositar(200)
	poupancaDoVinicius.Depositar(100)

	fmt.Println(contaDoVinicius.ObterSaldo())
	fmt.Println(poupancaDoVinicius.ObterSaldo())

	PagarBoleto(&contaDoVinicius, 50)
	PagarBoleto(&poupancaDoVinicius, 50)

	fmt.Println(contaDoVinicius.ObterSaldo())
	fmt.Println(poupancaDoVinicius.ObterSaldo())
}
