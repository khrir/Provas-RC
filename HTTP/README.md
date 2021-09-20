Opa!

###  Ecoador Tabajara com socket HTTP 

Execute primeiro o servidor:

	  go run main.go
  
Acesse o endereço informado no navegador:
  
    localhost:8080 ou
    127.0.0.1:8080

-------------------------------------------------------------------------
O que acontece se a ordem for trocada?
 
	  Bem, vai dar erro de conexão recusada, visto que o cliente estará
	  interagindo com um servidor que não está conectado.

-------------------------------------------------------------------------
Qual a proposta desse projeto?

	  O servidor vai ouvir a mensagem do cliente, através de um formulário,
    e vai retornar, ao cliente, o que foi recebido na página.
