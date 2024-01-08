package main

import (
        "fmt"
        "log"
        "net"
        "net/http"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
        tcpConnActive = prometheus.NewGauge(prometheus.GaugeOpts{
                Name:      "tcp_conn_active",
                Help:      "The number of active tcp connections",
        })

        NbrOfConn = 0
)

type Server struct {
        listenAddr      string
        ln                      net.Listener
        quitch          chan struct{}
}

func NewServer(listenAddr string) *Server  {
        return &Server{
                listenAddr: listenAddr,
                quitch: make(chan struct{}),
        }
}

func (s *Server) Start() error {
        ln, err := net.Listen("tcp", s.listenAddr)
        if err != nil {
                return err
        }
        defer ln.Close()
        s.ln = ln
        go s.acceptLoop()
        <-s.quitch
        return nil
}

func (s *Server) acceptLoop() {
        for {
                conn, err := s.ln.Accept()
                if err != nil {
                        fmt.Printf("Accept Error: %v\n", err)
                        continue
                }
                NbrOfConn++
                fmt.Println("New connection to the server from: ", conn.RemoteAddr())
                fmt.Printf("The number of connections is: %v\n", NbrOfConn)
                tcpConnActive.Add(1)
                go s.readLoop(conn)
        }
}

func (s *Server) readLoop(conn net.Conn) {
        defer conn.Close()
        buf := make([]byte, 2048)
        for {
                n, err := conn.Read(buf)
                if err != nil {
                        fmt.Printf("Read Error: %v\n", err)
                        tcpConnActive.Dec()
                        return
                }
                msg := buf[:n]
                fmt.Println(string(msg))
        }
}

// Start prometheus HTTP endpoint
func StartMetrics() {
        // Register tcpConnActive gauge metric
        prometheus.MustRegister(tcpConnActive)

        // Exposing PRometheus endpoint
        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2222", nil)
}

func main()  {
        // Start Prometheus metrics endpoint
        go StartMetrics()

        // Exposing TCP Server
        server := NewServer(":3333")
        log.Fatal(server.Start())
}
