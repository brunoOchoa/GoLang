package main

import "fmt"

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *contaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *contaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor errado!", c.saldo
	}
}

func (c *contaCorrente) Transferir(valorDaTransferencia float64, contaDestino *contaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {

	contaDaSu := contaCorrente{titular: "Suélen", saldo: 100}
	contaDoEloy := contaCorrente{titular: "Eloy", saldo: 100}

	status := contaDaSu.Transferir(200, &contaDoEloy)

	fmt.Println(status)
	fmt.Println(contaDaSu)
	fmt.Println(contaDoEloy)

	// contaDoBruno := contaCorrente{
	// 	titular:       "Bruno",
	// 	numeroAgencia: 7,
	// 	numeroConta:   123456,
	// 	saldo:         125.25,
	// }

	// fmt.Println(contaDoBruno)

	// contaDaSu := contaCorrente{
	// 	"Suélen",
	// 	8,
	// 	123457,
	// 	5660.00,
	// }

	// fmt.Println(contaDaSu)

	// var contaDoEloy *contaCorrente
	// contaDoEloy = new(contaCorrente)
	// contaDoEloy.titular = "Eloy"
	// contaDoEloy.saldo = 789

	// fmt.Println(*contaDoEloy)
