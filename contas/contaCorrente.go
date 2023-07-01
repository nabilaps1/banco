package contas

import (
	"fmt"

	"github.com/nabilaps1/banco/clientes"
)

// A primeira letras dos campos define a visibilidade dos atributos. Se minuscula, privado. Se maiuscula, publico
type ContaCorrente struct {
	Titular     clientes.Titular
	NumAgencia  int
	NumeroConta int
	saldo       float64
}

// c *ContaCorrente  = this em java (ou self em python). siginifica que o metodo sacar smp vai apontar para a conta q esta chamando
func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor deposito menor que 0.", c.saldo
	}
}

// contaDesino *ContaCorrente aponta para a conta que quero fazer a alteracao
func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDesino *ContaCorrente) bool {
	fmt.Println("Status transferencia: ")
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDesino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

// maneiras de se utilizar de uma struct

// var conta3 *ContaCorrente = new(ContaCorrente)
// conta3.titular = "Cris"
// contaExemplo := contas.ContaCorrente{}
// fmt.Println(contaExemplo) // cria uma nova estrutura com inicializacao zero
// fmt.Println(conta1)
// fmt.Println(conta2)
// fmt.Println(conta3)  // traz o endereco (&) e mostra o conteudo
// fmt.Println(*conta3) //  mostra o conteudo
