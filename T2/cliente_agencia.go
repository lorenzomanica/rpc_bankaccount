package main

import (
	"fmt"
	"log"
	"net/rpc"

	"./admin"
)

var conta admin.Conta
var assinatura int

var reply admin.Conta
var reply2 bool
var reply3 int

func main() {
	/*
	   Inicializa o cliente na porta 4040 do localhost
	   utilizando o protocolo tcp. Se o servidor estiver
	   em outra maquina deve ser utilizado IP:porta no
	   segundo parametro.
	*/
	c, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Error dialing: ", err)
	}

	/*
	   Call chama um metodo atrves da conexao estabelecida
	   anteriormente. Os parametros devem ser:
	   -Metodo a ser chamado no servidor no formato 'Tipo.Nome'.
	   Este parametro deve ser uma string
	   -Primeiro argumento do metodo
	   -Segundo argumento do metodo(ponteiro para receber a resposta)
	*/

	var op int = -1

	for op != 0 {
		printMenu()
		fmt.Scanln(&op)
		assinatura = 0

		switch op {
		case 1:
			err = c.Call("Conta.GerarAssinatura", assinatura, &reply3)
			if err != nil {
				log.Fatal("Conta error: ", err)
			}
			conta.Assinatura = reply3
			if conta.Assinatura != 0 {
				err = c.Call("Conta.AbrirConta", conta, &reply)
				if err != nil {
					log.Fatal("Conta error: ", err)
				}
				fmt.Printf("Conta número: %d criada\n", reply.Numero)
				fmt.Printf("Saldo da conta: %.2f\n", reply.Saldo)
				fmt.Printf("Status da conta: %t\n", reply.Ativa)
				reply.Saldo = 0
			}

		case 2:
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

		case 3:
			fmt.Println("Informe o número da conta:")

			fmt.Scanln(&conta.Numero)
			err = c.Call("Conta.Autenticar", conta, &reply2)
			if err != nil {
				log.Fatal("Autenticar error: ", err)
			}
			if reply2 == true {
				fmt.Println("Conta existente")
			} else {
				fmt.Println("Conta inexistente")
			}

		case 4:
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

		case 5:
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

		case 6:
			fmt.Println("Informe o número da conta:")

			fmt.Scanln(&conta.Numero)
			err = c.Call("Conta.Autenticar", conta, &reply2)
			if err != nil {
				log.Fatal("Autenticar error: ", err)
			}
			if reply2 == true {
				err = c.Call("Conta.FecharConta", conta, &reply)
				if err != nil {
					log.Fatal("Depositar error: ", err)
				} else {
					fmt.Println("Conta fechada")
				}
			}
		case 7:
			simulaErroAbrirConta(c, err)
		case 8:
			simulaErroSacar(c, err)
		case 9:
			simulaErroDepositar(c, err)
		case 0:
			log.Fatal("Fechando")
		default:
			fmt.Println("Comando inválido!")
		}

		op = -1
		fmt.Println("\n\n")
	}

}

func printMenu() {
	fmt.Println("Digite a operação desejada:")
	fmt.Println("1 - Abrir Conta")
	fmt.Println("2 - Consultar Saldo")
	fmt.Println("3 - Autenticar Conta")
	fmt.Println("4 - Sacar")
	fmt.Println("5 - Depositar")
	fmt.Println("6 - Fechar Conta")
	fmt.Println("7 - Abrir conta com erro (teste)")
	fmt.Println("8 - Sacar com erro (teste)")
	fmt.Println("9 - Depositar com erro (teste)")
	fmt.Println("0 - Sair")
	fmt.Println("")
}

func simulaErroAbrirConta(c *rpc.Client, err error) {
	err = c.Call("Conta.GerarAssinatura", assinatura, &reply3)
	if err != nil {
		log.Fatal("Conta error: ", err)
	}
	conta.Assinatura = reply3
	if conta.Assinatura != 0 {
		err = c.Call("Conta.AbrirContaTimeout", conta, &reply)
		if err != nil {
			//time.Sleep(10 * time.Second)
			log.Println("Erro na operação Abrir conta: ", err)
			log.Println("Executando segunda tentativa")
			err = c.Call("Conta.AbrirConta", conta, &reply)
			if err != nil {
				log.Println("Erro na operação Abrir conta: ", err)
			} else {
				fmt.Printf("Conta número: %d criada\n", reply.Numero)
				fmt.Printf("Saldo da conta: %.2f\n", reply.Saldo)
				fmt.Printf("Status da conta: %t\n", reply.Ativa)
				reply.Saldo = 0
			}
		} else {
			fmt.Printf("Conta número: %d criada\n", reply.Numero)
			fmt.Printf("Saldo da conta: %.2f\n", reply.Saldo)
			fmt.Printf("Status da conta: %t\n", reply.Ativa)
			reply.Saldo = 0
		}
	}
}

func simulaErroSacar(c *rpc.Client, err error) {
	fmt.Println("Informe o número da conta:")
	fmt.Scanln(&conta.Numero)
	err = c.Call("Conta.GerarAssinatura", assinatura, &reply3)
	assinatura = reply3
	if err != nil {
		log.Fatal("Erro na operação Gerar assinatura: ", err)
	}
	conta.Assinatura = assinatura
	if assinatura != 0 {
		err = c.Call("Conta.Autenticar", conta, &reply2)
		if err != nil {
			log.Fatal("Erro na operação Autenticar: ", err)
		}
		if reply2 == true {
			fmt.Println("Valor do saque:")
			fmt.Scanln(&conta.Movimentacao)
			err = c.Call("Conta.SacarErro", conta, &reply)
			if err != nil {
				log.Println("Erro na operação Sacar: ", err)
				log.Println("Executando segunda tentativa")
				err = c.Call("Conta.Autenticar", conta, &reply2)
				if err != nil {
					log.Fatal("Erro na operação Autenticar: ", err)
				}
				if reply2 == true {
					err = c.Call("Conta.Sacar", conta, &reply)
					if err != nil {
						log.Println("Erro na operação Sacar: ", err)
					} else {
						if conta.Assinatura != -1 {
							fmt.Printf("Conta: %d\nSaldo: %.2f\n", reply.Numero, reply.Saldo)
							reply.Saldo = 0
						}
					}
				}
			} else {
				if conta.Assinatura != -1 {
					fmt.Printf("Conta: %d\nSaldo: %.2f\n", reply.Numero, reply.Saldo)
					reply.Saldo = 0
				}
			}
		}
	}
}

func simulaErroDepositar(c *rpc.Client, err error) {
	fmt.Println("Informe o número da conta:")
	fmt.Scanln(&conta.Numero)
	err = c.Call("Conta.GerarAssinatura", assinatura, &reply3)
	assinatura = reply3
	if err != nil {
		log.Fatal("Erro na operação Gerar assinatura: ", err)
	}
	conta.Assinatura = assinatura
	if assinatura != 0 {
		err = c.Call("Conta.Autenticar", conta, &reply2)
		if err != nil {
			log.Fatal("Erro na operação Autenticar: ", err)
		}
		if reply2 == true {
			fmt.Println("Valor do deposito:")
			fmt.Scanln(&conta.Movimentacao)
			err = c.Call("Conta.DepositarErro", conta, &reply)
			if err != nil {
				log.Println("Erro na operação Depositar: ", err)
				log.Println("Executando segunda tentativa")
				err = c.Call("Conta.Autenticar", conta, &reply2)
				if err != nil {
					log.Fatal("Erro na operação Autenticar: ", err)
				}
				if reply2 == true {
					err = c.Call("Conta.Depositar", conta, &reply)
					if err != nil {
						log.Println("Erro na operação Depositar: ", err)
					} else {
						if conta.Assinatura != -1 {
							fmt.Printf("Conta: %d\nSaldo: %.2f\n", reply.Numero, reply.Saldo)
							reply.Saldo = 0
						}
					}
				}
			} else {
				if conta.Assinatura != -1 {
					fmt.Printf("Conta: %d\nSaldo: %.2f\n", reply.Numero, reply.Saldo)
					reply.Saldo = 0
				}
			}
		}
	}
}
