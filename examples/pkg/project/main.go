// Package project is amazing. It fires up all our servers for us on different goroutines
package project

import (
	"errors"
	"fmt"
	"time"
)

// Server - our interface for a server
type Server interface {
	GetName() string
	Start() error
}

type baseServer struct {
	Name string
	Port string
}

// WebServer - wow, it's our magic web server! It's amazing
type WebServer struct {
	baseServer // Added by composition
	SiteTitle  string
}

// APIServer - this is just as amazing as our web server
type APIServer struct {
	baseServer
	BasePath string
}

func (bs baseServer) GetName() string {
	return bs.Name
}

// Start - add this method allows WebServer to be an instance of Server
func (ws WebServer) Start() error {
	fmt.Println("Starting web server")
	time.Sleep(time.Second * 10)
	return errors.New("Web server has died")
}

// Start - add this method allows ApiServer to be an instance of Server
// Simply run this bad boy with:
//   as.Start()
func (as APIServer) Start() error {
	fmt.Println("Starting API server")
	time.Sleep(time.Second * 15)
	return errors.New("Api server has died")
}

// Run - runs our project
func Run() {
	ws := WebServer{baseServer: baseServer{Name: "My web server"}, SiteTitle: "Site title"}
	as := APIServer{baseServer: baseServer{Name: "My API server"}, BasePath: "/base/path"}
	servers := []Server{}
	servers = append(servers, ws, as)
	StartServers(servers)
}

// StartServers - Starts all servers in slice of Servers
func StartServers(servers []Server) {
	c := make(chan string)

	for _, s := range servers {
		go func(s Server) {
			fmt.Printf("%v - %T\n", s.GetName(), s)
			err := s.Start()
			c <- err.Error()
		}(s)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-c)
	}
}
