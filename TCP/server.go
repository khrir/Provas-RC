
package main

import(
  "fmt"
  "net"
)

func checkError(err error){
  if err != nil{
    panic(err)
  }
}

func main() {

  //Prepara o servidor para ouvir requisições
  ln, err := net.Listen("tcp", ":8080")
  checkError(err)

  for{
    //Aceita as conexões que o servidor receber
    conn, err := ln.Accept()
    checkError(err)

    go respond(conn)
  }
}

func respond(conn net.Conn){
  //Lê as requisições do cliente 
  resp := make([]byte, 1024)
  n, err := conn.Read(resp)
  checkError(err)

  //Mostra o que foi recebido e reenvia de volta ao cliente
  fmt.Println("\nO servidor recebeu: ", string(resp[0:n]))
  conn.Write(resp)
}