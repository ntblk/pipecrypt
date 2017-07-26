package main

import (
	"io"
	"log"
	"net"
	"time"
)


type SNIProxy struct {
	// DialTimeout sets the timeout for establishing the outbound
	// connection.
	// TODO: Do we need to set this?
	DialTimeout time.Duration

	// Lookup returns a target host for the given server name.
	// The proxy will panic if this value is nil.
	//Lookup func(host string) *route.Target
}

func (p *SNIProxy) ServeTCP(in net.Conn) error {
	defer in.Close()

	// capture client hello
	data := make([]byte, 1024)
	n, err := in.Read(data)
	if err != nil {
		return err
	}
	data = data[:n]

	host, ok := readServerName(data)
	if !ok {
		log.Print("[DEBUG] tcp+sni: TLS handshake failed")
		return nil
	}

	if host == "" {
		log.Print("[DEBUG] tcp+sni: server_name missing")
		return nil
	}

/*
	t := p.Lookup(host)
	if t == nil {
		return nil
	}
	addr := t.URL.Host
*/
//addr := "178.62.180.138:443"
addr := host + ":443"


	out, err := net.DialTimeout("tcp", addr, p.DialTimeout)
	if err != nil {
		log.Print("[WARN] tcp+sni: cannot connect to upstream ", addr)
		return err
	}
	defer out.Close()

	// copy client hello
	n, err = out.Write(data)
	if err != nil {
		log.Print("[WARN] tcp+sni: copy client hello failed. ", err)
		return err
	}

	errc := make(chan error, 2)
	cp := func(dst io.Writer, src io.Reader) {
		errc <- copyBuffer(dst, src)
		//buf := make([]byte, 32*1024)
		//errc <- io.CopyBuffer(dst, src, buf)
	}

	go cp(in, out)
	go cp(out, in)

	err = <-errc
	if err != nil && err != io.EOF {
		log.Print("[WARN]: tcp+sni:  ", err)
		return err
	}
	return nil
}
