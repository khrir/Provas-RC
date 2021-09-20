package main

import(
  "fmt"
  "net"
)

func checkError(err error){
  if err != nil{
    fmt.Println("Erro: ", err)
  }
}

func main() {

  //disca a conexão
  conn, err := net.Dial("tcp", ":8080")
  defer conn.Close()
  checkError(err)

  //Cria a mensagem, empacota e manda pro servidor
  msg := fmt.Sprint("Opa")
	buf := []byte(msg)
  conn.Write(buf)

  //Recebe a resposta do servidor e mostra
  n, err := conn.Read(buf)
  checkError(err)
  fmt.Println("\nO servidor respondeu: ", string(buf[0:n]))
  
}