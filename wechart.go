
// http://blog.csdn.net/ygrx/article/details/11773151
package main

import (
	"fmt"
	"net"
	"os"
)

//error check
func checkError(err error, info string) (res bool)  {
	if (err != nil) {
		fmt.Println(info+" "+err.Error())
		return false
	}
	return true
}


//服务器接收数据线程

func Handler(conn net.Conn, messages chan string)  {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

	buf := make([]byte, 1024)

	for{
		length, err := conn.Read(buf)
		if(checkError(err, "connection")==false){
			conn.Close()
			break
		}
		if length > 0 {
			buf[length] = 0
		}


		reciveStr := string(buf[0:length])
		messages <- reciveStr

	}
}


//服务器发送数据线程
func echoHandler(conns *map[string]net.Conn, messages chan string)  {
	for  {
		msg := <- messages
		fmt.Println(msg)

		for key,value := range *conns {
			fmt.Println("connection is connected from ...", key)
			_,err := value.Write([]byte(msg))
			if (err != nil) {
				fmt.Println(err.Error())
				delete(*conns,key)
			}
		}
		
	}

}

//启动服务器

func StartServer(port string)  {
	service := ":"+port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "resolveTCPAddr")
	l, err:= net.ListenTCP("tcp", tcpAddr)

	checkError(err, "ListenTCP")
	conns :=make(map[string]net.Conn)
	messages := make(chan string, 10)

	go echoHandler(&conns, messages)

	for {
		fmt.Println("Listening ...")
		conn, err := l.Accept()
		checkError(err, "Accept")
		fmt.Println("Accepting ...")
		conns[conn.RemoteAddr().String()] = conn
		//启动线程
		go Handler(conn, messages)
	}
}


// 客户端发送线程

func chatSend(conn net.Conn)  {
	var input string
	username := conn.LocalAddr().String()
	for {
		fmt.Scanln(&input)

		if input == "/quit" {
			fmt.Println("ByeBye..")
			conn.Close()
			os.Exit(0)
		}

		lens,err := conn.Write([]byte(username+"Say :::" + input))
		fmt.Println(lens)
		if(err != nil){
			fmt.Println(err.Error())
			conn.Close()
			break
		}
	}


}


//客户端启动函数

func StartClient(tcpaddr string)  {
	tcpAddr,err := net.ResolveTCPAddr("tcp4", tcpaddr)
	checkError(err, "ResolveTCPAddr")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "DialTCP")
	//启动客户端发送线程
	go chatSend(conn)

	//客户端轮循
	buf := make([]byte, 1024)
	for{
		length, err := conn.Read(buf)
		if (checkError(err, "connection") == false) {
			conn.Close()
			fmt.Println("Server is dead ... ByeBye")
			os.Exit(0)
		}
		fmt.Println(string(buf[0:length]))
	}
}

func main()  {
	if len(os.Args)!=3 {
		fmt.Println("Wrong pare")
		os.Exit(0)
	}
	if os.Args[1] == "server" && len(os.Args) == 3 {
		StartServer(os.Args[2])
	}

	if os.Args[1] == "client" && len(os.Args) == 3 {
		StartClient(os.Args[2])
	}
}




