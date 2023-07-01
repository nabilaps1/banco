package main

// Aplicacao para gerenciar contas correntes de um banco

import (
	"fmt"

	"github.com/nabilaps1/banco/clientes"
	"github.com/nabilaps1/banco/contas"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

func main() {
	cliente1 := clientes.Titular{Nome: "Nabila", CPF: "123.132.123.22", Profissao: "Desenvolvedora"}
	contaCliente1 := contas.ContaCorrente{Titular: cliente1}

	contaCliente1.Depositar(500)
	PagarBoleto(&contaCliente1, 200)

	fmt.Println(contaCliente1.ObterSaldo())

}
