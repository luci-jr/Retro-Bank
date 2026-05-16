package main

import (
	"banco/contas"
	"fmt"
	"time"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto, "Pagamento de boleto")
}

type verificarConta interface {
	Sacar(valor float64, operacao string) string
}

func main() {
	fmt.Println("========================================")
	fmt.Println(" [ SYSTEM BOOT ] ... OK")
	time.Sleep(1 * time.Second)
	fmt.Println(" [ CONEXÃO DIAL-UP ] ... ESTABELECIDA")
	time.Sleep(1 * time.Second)
	fmt.Println(" ::: BEM-VINDO AO MAINFRAME DO RETRO BANK :::")
	fmt.Println("========================================")
	fmt.Println()
	time.Sleep(1 * time.Second)

	fmt.Println(">>> ACESSANDO DADOS DO CLIENTE 01 (Cliente01)...")
	contaCliente01 := contas.ContaCorrente{}
	contaCliente01.Depositar(1000)
	PagarBoleto(&contaCliente01, 500)

	fmt.Println(contaCliente01.ConsultaSaldo())

	time.Sleep(1 * time.Second)
	fmt.Println(">>> ACESSANDO DADOS DA CLIENTE 02 (Cliente02)...")
	contaCliente02 := contas.ContaPoupanca{}
	contaCliente02.Depositar(10)
	PagarBoleto(&contaCliente02, 300)

	fmt.Println(contaCliente02.ConsultaSaldo())

	time.Sleep(1 * time.Second)
	fmt.Println(">>> DESCONECTANDO... TERMINAL ENCERRADO.")
}