package admin

import (
	"fmt"
	"log"
)

type Conta struct {
	Numero       int
	Saldo        float64
	Movimentacao float64
	Ativa        bool
	Assinatura   int
}

var gerador int = 0
var database []Conta
var assinatura int = 0

func (c *Conta) Autenticar(conta Conta, reply *bool) error {
	*reply = false
	for i := 0; i < gerador; i++ {
		if database[i].Numero == conta.Numero {
			if database[i].Ativa {
				*reply = true
				return nil
			} else {
				*reply = false
				return nil
			}
		}
	}
	return fmt.Errorf("Conta inválida")
}

func (c *Conta) GerarAssinatura(tmp int, reply *int) error {
	assinatura = assinatura + 1
	*reply = assinatura
	return nil
}

func (c *Conta) AbrirConta(conta Conta, reply *Conta) error {
	if conta.Assinatura == assinatura {
		gerador = gerador + 1
		conta.Numero = gerador
		conta.Saldo = 0
		conta.Movimentacao = 0
		conta.Ativa = true
		database = append(database, conta)
		assinatura = assinatura + 1
		*reply = database[len(database)-1]
		return nil
	}
	log.Printf("Erro ao Abrir conta: Assinatura inválida %d\n", assinatura)
	return fmt.Errorf("Assinatura inválida")
}

func (c *Conta) AbrirContaErro(conta Conta, reply *Conta) error {
	log.Println("Executando AbrirConta")
	log.Println("Verificando Assinatura")
	if conta.Assinatura == assinatura {
		log.Println("Assinatura Verificada")
		log.Println("Executando Transação")
		gerador = gerador + 1
		conta.Numero = gerador
		conta.Saldo = 0
		conta.Movimentacao = 0
		conta.Ativa = true
		database = append(database, conta)
		assinatura = assinatura + 1
		*reply = database[len(database)-1]
		log.Println("Operação realizada com sucesso")
		return fmt.Errorf("Erro após operação concluída")
	}
	log.Printf("Erro ao Abrir conta: Assinatura inválida %d\n", assinatura)
	return fmt.Errorf("Assinatura inválida")
}

func (c *Conta) FecharConta(conta Conta, reply *Conta) error {
	for i := 0; i < gerador; i++ {
		if database[i].Numero == conta.Numero {
			database[i].Ativa = false
			*reply = database[i]
			return nil
		}
	}
	return fmt.Errorf("Conta inválida")
}

func (c *Conta) ConsultarSaldo(conta Conta, reply *Conta) error {
	for i := 0; i < gerador; i++ {
		if database[i].Numero == conta.Numero {
			*reply = database[i]
			return nil
		}
	}
	return fmt.Errorf("Conta inválida")
}

func (c *Conta) Sacar(conta Conta, reply *Conta) error {
	if conta.Assinatura == assinatura {
		for i := 0; i < gerador; i++ {
			if database[i].Numero == conta.Numero {
				database[i].Saldo -= conta.Movimentacao
				assinatura = assinatura + 1
				*reply = database[i]
				return nil
			}
		}
	}
	log.Printf("Erro ao sacar: Assinatura inválida %d\n", assinatura)
	return fmt.Errorf("Assinatura inválida")
}

func (c *Conta) SacarErro(conta Conta, reply *Conta) error {
	log.Println("Executando Sacar")
	log.Println("Verificando Assinatura")
	if conta.Assinatura == assinatura {
		log.Println("Assinatura Verificada")
		log.Println("Executando Transação")
		for i := 0; i < gerador; i++ {
			if database[i].Numero == conta.Numero {
				database[i].Saldo -= conta.Movimentacao
				assinatura = assinatura + 1
				*reply = database[i]
				log.Println("Operação realizada com sucesso")
				return fmt.Errorf("Erro após operação concluída")
			}
		}
	}
	log.Printf("Erro ao sacar: Assinatura inválida %d\n", assinatura)
	return fmt.Errorf("Assinatura inválida")
}

func (c *Conta) Depositar(conta Conta, reply *Conta) error {
	if conta.Assinatura == assinatura {
		for i := 0; i < gerador; i++ {
			if database[i].Numero == conta.Numero {
				database[i].Saldo += conta.Movimentacao
				assinatura = assinatura + 1
				*reply = database[i]
				return nil
			}
		}
	}
	log.Printf("Erro ao depositar: Assinatura inválida %d\n", assinatura)
	return fmt.Errorf("Assinatura inválida")
}

func (c *Conta) DepositarErro(conta Conta, reply *Conta) error {
	log.Println("Executando Depositar")
	log.Println("Verificando Assinatura")
	if conta.Assinatura == assinatura {
		log.Println("Assinatura Verificada")
		log.Println("Executando Transação")
		for i := 0; i < gerador; i++ {
			if database[i].Numero == conta.Numero {
				database[i].Saldo += conta.Movimentacao
				assinatura = assinatura + 1
				*reply = database[i]
				log.Println("Operação realizada com sucesso")
				return fmt.Errorf("Erro após operação concluída")
			}
		}
	}
	log.Printf("Erro ao depositar: Assinatura inválida %d\n", assinatura)
	return fmt.Errorf("Assinatura inválida")
}
