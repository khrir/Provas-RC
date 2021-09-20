Opa!

###  Ecoador Tabajara com socket TCP 

Execute primeiro o servidor:

	go run server.go
	
Logo após, execute o cliente:

	go run client.go

-------------------------------------------------------------------------
O que acontece se a ordem for trocada?
 
	Bem, vai dar erro de conexão recusada, visto que o cliente estará
	interagindo com um servidor que não está conectado.

-------------------------------------------------------------------------
Qual a proposta desse projeto?

	O servidor vai ouvir a mensagem do cliente e vai retornar,
	ao cliente, o que foi recebido.
