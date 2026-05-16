package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular        clientes.Titular
	Agencia, Conta int
	saldo          float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64, operacao string) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		fmt.Println(operacao, "realizado com sucesso!\nSaldo atual: R$", c.saldo)
		return operacao + " realizado com sucesso!"
	} else {
		fmt.Println("Saldo insuficiente para", operacao, "\nO saldo se mantem em: R$", c.saldo)
		return "saldo insuficiente."
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		fmt.Println("Deposito realizado com sucesso!\nSaldo atual: R$", c.saldo)
	} else {
		fmt.Println("Valor de deposito invalido.\nO saldo se mantem em: R$", c.saldo)
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDeDestino *ContaCorrente) {
	if valorTransferencia > 0 && valorTransferencia <= c.saldo {
		c.saldo -= valorTransferencia
		contaDeDestino.Depositar(valorTransferencia)
		fmt.Println("Transferencia realizada com sucesso!\nSeu saldo atual é: R$", c.saldo)
	} else {
		fmt.Println("Valor de transferencia é invalido.\nPor isso o saldo se mantem em: R$", c.saldo)
	}
}

func (c *ContaCorrente) ConsultaSaldo() string {
	return fmt.Sprintf("--- Saldo final: R$ %v ---\n", c.saldo)
}
