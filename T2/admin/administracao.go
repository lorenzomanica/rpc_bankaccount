package admin

import "fmt"

/*
	Tipo exportável e também tipo usado para realizar operações na administração
*/

type Conta struct {
	Numero       int
	Saldo        float64
	Movimentacao float64
	Ativa        bool
	Assinatura   int
}

var gerador int = 0
var assinatura int = 0
var database []Conta

/*
	Metodos devem:
	-Pertencer a um tipo exportavel (Arith neste caso) e ser exportaveis
	-Possuir dois parametros de entrada. O primeiro pode ser qualquer tipo
	exportavel ou tipo nativo de go. O segundo de ser obrigatoriamente
	um ponteiro. O segundo argumento eh usado para o retorno do metodo.
	-Retornar um erro. Se for retornado algo alem de nil o cliente recebera
	apenas o erro, sem o ponteiro de reply
*/
func (c *Conta) Abrir_conta(conta Conta, reply *Conta) error {
	if conta.Assinatura == assinatura {
		gerador = gerador + 1
		conta.Numero = gerador
		conta.Saldo = 0
		conta.Movimentacao = 0
		conta.Ativa = true

		database = append(database, conta)

		// Retorna a conta criada agora (última da slice)
		assinatura = assinatura + 1
		*reply = database[len(database)-1]
		return nil
	}

	return fmt.Errorf("Abrir_conta falhou!")
}

func (c *Conta) Consultar_saldo(conta Conta, reply *Conta) error {

	for i := 0; i < gerador; i++ {
		if database[i].Numero == conta.Numero {
			*reply = database[i]
			return nil
		}
	}
	return fmt.Errorf("Consultar_saldo falhou!")
}

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
	return fmt.Errorf("Autenticar falhou!")
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
	return fmt.Errorf("Sacar falhou!")
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
	return fmt.Errorf("Depositar falhou!")
}

func (c *Conta) Fechar_conta(conta Conta, reply *Conta) error {

	for i := 0; i < gerador; i++ {
		if database[i].Numero == conta.Numero {
			database[i].Ativa = false
			*reply = database[i]
			return nil
		}
	}
	return fmt.Errorf("Fechar_conta falhou!")
}

func (c *Conta) Gerar_assinatura(tmp int, reply *int) error {
	assinatura = assinatura + 1
	*reply = assinatura
	return nil
}
