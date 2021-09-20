package main

import (
	"net"
	"fmt"
)

//função para verificar erros
func CheckError(err error){
	if(err != nil){
		fmt.Println("Erro: ", err)
	}
}

func main(){

	//estabelecendo um endereço para a conexão udp
	sAddr, err := net.ResolveUDPAddr("udp",":8080")
	CheckError(err)

	//iniciando o servidor que vai ouvir as requisições
	sConn, err := net.ListenUDP("udp", sAddr)
	CheckError(err)
	defer sConn.Close()

	//criando um slicebyte para receber as mensagens do cliente
	buf := make([]byte, 1024)
    
	//lê as mensagens do cliente em bytes e converte para string
	//além de armazenar o endereço do cliente
	n, addr, err := sConn.ReadFromUDP(buf)
	CheckError(err)
    fmt.Println("\nO servidor recebeu: ",string(buf[0:n]), "\tde: ",addr)

	//envia o que foi recebido de volta ao cliente
	n, err = sConn.WriteTo(buf[0:n], addr)
	CheckError(err)
		
	
    

}