package main

import (
	"fmt"
	"log"
	"net/rpc"

	"./admin"
)

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

	//Variavel para receber os resultados
	var reply admin.Conta
	var reply3 int
	var reply2 bool
	var conta admin.Conta

	var assinatura int

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
		fmt.Scanln(&op)
		assinatura = 0

		switch op {
		case 1:
			err = c.Call("Conta.Gerar_assinatura", assinatura, &reply3)
			if err != nil {
				log.Fatal("Conta error: ", err)
			}
			conta.Assinatura = reply3
			if conta.Assinatura != 0 {
				err = c.Call("Conta.Abrir_conta", conta, &reply)
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
				err = c.Call("Conta.Consultar_saldo", conta, &reply)
				fmt.Printf("Conta: %d\n", reply.Numero)
				fmt.Printf("Saldo: %.2f\n", reply.Saldo)
				reply.Saldo = 0
				if err != nil {
					log.Fatal("Consultar_saldo error: ", err)
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

			err = c.Call("Conta.Gerar_assinatura", assinatura, &reply3)
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

			err = c.Call("Conta.Gerar_assinatura", assinatura, &reply3)
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
				err = c.Call("Conta.Fechar_conta", conta, &reply)
				if err != nil {
					log.Fatal("Depositar error: ", err)
				} else {
					fmt.Println("Conta fechada")
				}
			}
		case 7:
			err = c.Call("Conta.Gerar_assinatura", assinatura, &reply3)
			assinatura = reply3
			if err != nil {
				log.Fatal("Conta error: ", err)
			}
			conta.Assinatura = assinatura
			if assinatura != 0 {

				err = c.Call("Conta.Abrir_conta", conta, &reply)
				if err != nil {
					log.Fatal("Conta error: ", err)
				}
				fmt.Printf("Conta número: %d criada\n", reply.Numero)
				fmt.Printf("Saldo da conta: %.2f\n", reply.Saldo)
				reply.Saldo = 0

				err = c.Call("Conta.Abrir_conta", conta, &reply)
				if err != nil {
					log.Fatal("Conta error: ", err)
				}
				if conta.Assinatura != -1 {
					fmt.Printf("Conta: %d\n", reply.Numero)
					fmt.Printf("Saldo: %.2f\n", reply.Saldo)
					reply.Saldo = 0
				}
			}

		case 8:
			fmt.Println("Informe o número da conta:")

			fmt.Scanln(&conta.Numero)

			err = c.Call("Conta.Gerar_assinatura", assinatura, &reply3)
			assinatura = reply3
			if err != nil {
				log.Fatal("Conta error: ", err)
			}
			conta.Assinatura = assinatura
			if assinatura != 0 {
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
						if conta.Assinatura != -1 {
							fmt.Printf("Conta: %d\n", reply.Numero)
							fmt.Printf("Saldo: %.2f\n", reply.Saldo)
							reply.Saldo = 0
						}

					}
				}

				//segunda vez
				fmt.Println("segunda tentativa:")
				err = c.Call("Conta.Autenticar", conta, &reply2)
				if err != nil {
					log.Fatal("Autenticar error: ", err)
				}
				if reply2 == true {
					err = c.Call("Conta.Sacar", conta, &reply)
					if err != nil {
						log.Fatal("Sacar error: ", err)
					} else {
						if conta.Assinatura != -1 {
							fmt.Printf("Conta: %d\n", reply.Numero)
							fmt.Printf("Saldo: %.2f\n", reply.Saldo)
							reply.Saldo = 0
						}
					}
				}
			}

		case 9:
			fmt.Println("Informe o número da conta:")

			fmt.Scanln(&conta.Numero)

			err = c.Call("Conta.Gerar_assinatura", assinatura, &reply3)
			assinatura = reply3
			if err != nil {
				log.Fatal("Conta error: ", err)
			}
			conta.Assinatura = assinatura
			if assinatura != 0 {
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
						if conta.Assinatura != -1 {
							fmt.Printf("Conta: %d\n", reply.Numero)
							fmt.Printf("Saldo: %.2f\n", reply.Saldo)
							reply.Saldo = 0
						}
					}
				}

				//segunda tentativa
				fmt.Println("Segunda tentativa")
				err = c.Call("Conta.Autenticar", conta, &reply2)
				if err != nil {
					log.Fatal("Autenticar error: ", err)
				}
				if reply2 == true {
					err = c.Call("Conta.Depositar", conta, &reply)
					if err != nil {
						log.Fatal("Depositar error: ", err)
					} else {
						if conta.Assinatura != -1 {
							fmt.Printf("Conta: %d\n", reply.Numero)
							fmt.Printf("Saldo: %.2f\n", reply.Saldo)
							reply.Saldo = 0
						}
					}
				}
			}

		case 0:
			fmt.Println("Tchau!!")
		default:
			fmt.Println("Comando inválido!")
		}
	}

}
