package server

import "net"

func ListenRPC() (net.Listener, error) {
	return net.Listen("tcp", ":8000")
}

func main() {

}
