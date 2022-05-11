package calc

//Tipos exportaveis (inicio com letra maiuscula) podem ser registrados por servidores
type Args struct {
	A, B float64
}

/*
	O array memory eh privado do objeto Arith, portanto, o unico modo de acessa-lo eh
	atraves dos metodos exportaveis Store e Load
	Este array pode ser de qualquer tipo, inclusive tipos criados neste arquivo
*/
type Arith struct {
	memory [10]float64
}

/*
	Metodos devem:
	-Pertencer a um tipo exportavel (Arith neste caso) e ser exportaveis
	-Possuir dois parametros de entrada. O primeiro pode ser qualquer tipo
	exportavel ou tipo nativo de go. O segundo de ser obrigatoriamente
	um ponteiro. O segundo argumento eh usado para o retorno do metodo.
	-Retornar um erro. Se for retornado algo alem de nil o cliente recebera 
	apenas o erro, sem o ponteiro de reply
*/
func (a *Arith) Sum(args *Args, reply *float64) error {
	*reply = args.A + args.B
	return nil
}

func (a *Arith) Sub(args *Args, reply *float64) error {
	*reply = args.A - args.B
	return nil
}

func (a *Arith) Mult(args *Args, reply *float64) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Div(args *Args, reply *float64) error {
	*reply = args.A / args.B
	return nil
}

func (a *Arith) Store(args *Args, reply *float64) error {
	var pos int = int(args.A)
	a.memory[pos] = args.B
	return nil
}

func (a *Arith) Load(args *Args, reply *float64) error {
	var pos int = int(args.A)
	*reply = a.memory[pos]
	return nil
}