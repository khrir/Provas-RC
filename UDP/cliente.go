package main

import (
	"net"
	"fmt"
)

//função para verificar erros
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main (){

	//estabelece o endereço de comunicação com o servidor 
	sAddr,err := net.ResolveUDPAddr("udp",":8080")
    CheckError(err)
 
	//estabelece um endereço local para o cliente
	//Por se tratar do ":0", o endereço é aleatório
    lAddr, err := net.ResolveUDPAddr("udp", ":0" )
    CheckError(err)
	
	//disca uma conexão udp, com os endereços do local e do servidor
    conn, err := net.DialUDP("udp", lAddr, sAddr)
    CheckError(err)
	defer conn.Close()

	//cria uma mensagem que é convertida em bytes e a escreve na conexão direcionada ao servidor
	msg := fmt.Sprint("Opa")
	buf := []byte(msg)
    conn.Write(buf)

	//lê a resposta do servidor
	n, addr, err := conn.ReadFromUDP(buf)
	CheckError(err)
	fmt.Println("\nO servidor falou: ", string(buf[0:n]), "\tde: ", addr)
	

}