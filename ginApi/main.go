package main

import (
	"fmt"
	"ginApi/route"
	"ginApi/vue"
	"github.com/gin-gonic/gin"
	"net"
	"os"
)

func main() {
	r := gin.Default()
	r.Use(vue.Cors())
	route.User(r)
	r.Run(":7000")

	//service := flag.String("host", "127.0.0.1:8888", "an ip address")
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stdout, "Usage of %s:\n", "mock http request")
	//	flag.PrintDefaults()
	//}
	//flag.Parse()
	//
	//tcp_addr, err := net.ResolveTCPAddr("tcp4", *service)
	//checkError(err)
	//// 发起连接请求
	//conn, err := net.DialTCP("tcp", nil, tcp_addr)
	//checkError(err)
	//
	//// 读写数据
	//_, err = conn.Write([]byte("timest amp"))
	//checkError(err)
	//
	//re(conn)
	//
	//// 关闭连接
	//conn.Close()
	//os.Exit(0)

}

func re(conn *net.TCPConn) {
	buffer := make([]byte, 256)
	_, err := conn.Read(buffer)
	checkError(err)
	fmt.Println("[client] receive from:", conn.RemoteAddr().String())
	fmt.Println(string(buffer))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
