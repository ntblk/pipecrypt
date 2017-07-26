package main

import (
    "fmt"
    "net"
    "os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
)

/*
func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer l.Close()
	return s.Serve(l)
}
*/
func main() {
    // Listen for incoming connections.
    l, err := net.Listen("tcp", ":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

    prox := new(SNIProxy);

    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            // FIXME
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go prox.ServeTCP(conn);
    }
}
