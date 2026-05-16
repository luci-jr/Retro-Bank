package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64, operacao string) string {
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

func (c *ContaPoupanca) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		fmt.Println("Deposito realizado com sucesso!\nSaldo atual: R$", c.saldo)
	} else {
		fmt.Println("Valor de deposito invalido.\nO saldo se mantem em: R$", c.saldo)
	}
}

func (c *ContaPoupanca) ConsultaSaldo() string {
	return fmt.Sprintf("--- Consulta final: R$ %v ---\n", c.saldo)
}