package main

import (
	"context" //Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	"fmt"     //Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"net"     //Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.

	proto "github.com/karim12345-gif/grpc_codingtest/service"
	"google.golang.org/grpc"            //Package grpc implements an RPC system called gRPC.See grpc.io for more information about gRPC.
	"google.golang.org/grpc/reflection" //To register server reflection on a gRPC server:
)

//clean way to add the tcp and port
const (
	protocol = "tcp"
	port     = ":2001"
)

func ServerCalc() {
	//Println formats using the default formats for its operands and writes to
	//standard output. Spaces are always added between operands and a newline is
	//appended. It returns the number of bytes written and any write error
	//encountered.
	fmt.Println("gRPC Server is working!")
	// listen to server with port 4040
	//Specify the port we want to use to listen for clients
	listener, err := net.Listen(protocol, port)
	// check for error
	if err != nil {
		panic(err)
	}

	// NewServer creates a gRPC server which has no service registered and has not
	// started to accept requests yet
	s := grpc.NewServer()

	proto.RegisterCalculatorServer(s, &Server{})

	//func Register(s GRPCServer)
	//Register registers the server reflection service on the given gRPC server.
	reflection.Register(s)

	if e := s.Serve(listener); e != nil {
		panic(err)

	}
}

type Server struct {
	//must be embedded to have forward compatible implementations.
	proto.UnimplementedCalculatorServer
}

///addition
func (s *Server) Add(ctx context.Context, in *proto.AddRequest) (*proto.AddReply, error) {
	// Printf formats according to a format specifier and writes to standard
	// output. It returns the number of bytes written and any write error
	//encountered.
	// return &proto.AddReply{Number1: in.GetNumber1() + in.GetNumber2()}, nil
	//fmt.Printf("Addition function called with: %v\n", in, &result)
	n1, n2 := in.GetNumber1(), in.GetNumber2()
	result := n1 + n2
	fmt.Printf("Addition function called with: %v\n", in)
	return &proto.AddReply{Number1: result}, nil
}

//subtraction
func (s *Server) Subtract(ctx context.Context, in *proto.SubtractRequest) (*proto.SubtractReply, error) {
	// Printf formats according to a format specifier and writes to standard
	// output. It returns the number of bytes written and any write error
	//encountered.

	n1, n2 := in.GetNumber1(), in.GetNumber2()
	result := n1 - n2
	fmt.Printf("Subtraction function called with: %v\n", in)
	return &proto.SubtractReply{Number1: result}, nil

}

//multiplication
func (s *Server) Multiply(ctx context.Context, in *proto.MultiplyRequest) (*proto.MultiplyReply, error) {
	// Printf formats according to a format specifier and writes to standard
	// output. It returns the number of bytes written and any write error
	//encountered.
	//return &proto.MultiplyReply{Number1: in.GetNumber1() * in.GetNumber2()}, nil

	n1, n2 := in.GetNumber1(), in.GetNumber2()
	result := n1 * n2
	fmt.Printf("Multiplication function called with: %v\n", in)
	return &proto.MultiplyReply{Number1: result}, nil
}

//division
func (s *Server) Divide(ctx context.Context, in *proto.DivideRequest) (*proto.DivideReply, error) {
	// Printf formats according to a format specifier and writes to standard
	// output. It returns the number of bytes written and any write error
	//encountered.

	n1, n2 := in.GetNumber1(), in.GetNumber2()
	result := n1 / n2
	fmt.Printf("Division function called with: %v\n", in)
	return &proto.DivideReply{Number1: result}, nil

}

func main() {
	ServerCalc()
}
