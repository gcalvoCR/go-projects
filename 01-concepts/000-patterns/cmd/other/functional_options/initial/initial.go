package initial

import "fmt"

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type Server struct {
	Opts Opts
}

func newServer(opts Opts) *Server {
	return &Server{
		Opts: opts,
	}
}

func MainServer() {
	// Issue --> What if I pass an empty Opts struct?
	// We should be able to apply a default config
	s := newServer(Opts{})
	fmt.Println(s)
}
