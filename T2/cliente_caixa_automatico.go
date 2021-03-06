package main

import (
	"fmt"
	"log"
	"net/rpc"

	"./admin"
)

func main() {

	var reply admin.Conta
	var reply2 bool
	var reply3 int
	var conta admin.Conta
	var assinatura int

	c, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Error dialing: ", err)
	}

	var op int = -1

	for op != 0 {
		fmt.Println("Digite a operação desejada:")
		fmt.Println("1 - Depositar")
		fmt.Println("2 - Sacar")
		fmt.Println("3 - Consultar Saldo")
		fmt.Println("0 - Sair")
		fmt.Scanln(&op)

		switch op {
		case 1:
			fmt.Println("Informe o número da conta:")
			fmt.Scanln(&conta.Numero)

			err = c.Call("Conta.GerarAssinatura", assinatura, &reply3)
			if err != nil {
				log.Fatal("Conta error: ", err)
			}
			conta.Assinatura = reply3
			if conta.Assinatura != 0 {
				err = c.Call("Conta.Autenticar", conta, &reply2)
				if err != nil {
					log.Fatal("Autenticar error: ", err)
				}
				if reply2 == true {
					fmt.Println("Valor do deposito:")
					fmt.Scanln(&conta.Movimentacao)
					err = c.Call("Conta.Depositar", conta, &reply)
					if err != nil {
						log.Fatal("Depositar error: ", err)
					} else {
						fmt.Printf("Conta: %d\n", reply.Numero)
						fmt.Printf("Saldo: %.2f\n", reply.Saldo)
						reply.Saldo = 0
					}
				}
			}

		case 2:
			fmt.Println("Informe o número da conta:")

			fmt.Scanln(&conta.Numero)

			err = c.Call("Conta.GerarAssinatura", assinatura, &reply3)
			if err != nil {
				log.Fatal("Conta error: ", err)
			}
			conta.Assinatura = reply3
			if conta.Assinatura != 0 {
				fmt.Printf("assiantura: %d \n", assinatura)
				err = c.Call("Conta.Autenticar", conta, &reply2)
				if err != nil {
					log.Fatal("Autenticar error: ", err)
				}
				if reply2 == true {
					fmt.Println("Valor do saque:")
					fmt.Scanln(&conta.Movimentacao)
					err = c.Call("Conta.Sacar", conta, &reply)
					if err != nil {
						log.Fatal("Sacar error: ", err)
					} else {
						fmt.Printf("Conta: %d\n", reply.Numero)
						fmt.Printf("Saldo: %.2f\n", reply.Saldo)
						reply.Saldo = 0
					}
				}
			}

		case 3:
			fmt.Println("Informe o número da conta:")

			fmt.Scanln(&conta.Numero)

			err = c.Call("Conta.Autenticar", conta, &reply2)
			if err != nil {
				log.Fatal("Autenticar error: ", err)
			}
			if reply2 == true {
				err = c.Call("Conta.ConsultarSaldo", conta, &reply)
				fmt.Printf("Conta: %d\n", reply.Numero)
				fmt.Printf("Saldo: %.2f\n", reply.Saldo)
				reply.Saldo = 0
				if err != nil {
					log.Fatal("ConsultarSaldo error: ", err)
				}
			} else {
				fmt.Println("Conta inexistente")
			}

		case 0:
			fmt.Println("Encerrando")
		default:
			fmt.Println("Comando inválido!")
		}
	}
}
