package main

import (
	//"bufio"§
	"context" //Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	"fmt"     //Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"log"     //Package log implements a simple logging package.
	"time"    //Package time provides functionality for measuring and displaying time.

	proto "github.com/karim12345-gif/grpc_codingtest/service"
	"google.golang.org/grpc" //Package grpc implements an RPC system called gRPC.See grpc.io for more information about gRPC.
)

//Package constant implements Values representing untyped Go constants and their corresponding operations.
//adding the port number Package constant implements Values representing untyped Go constants and their corresponding operations.
const (
	address = "localhost:2001"
)

var n1, n2 int32
var operator string //Package strings implements simple functions to manipulate UTF-8 encoded strings.

func ClientCalc() {

	//Println formats using the default formats for its operands and writes to
	//standard output. Spaces are always added between operands and a newline is
	// appended. It returns the number of bytes written and any write error
	// encountered.
	fmt.Println("This is the gRPC Client, welcome to the calculator....choose whatever operator you want!")

	//Print formats using the default formats for its operands and writes to
	// standard output. Spaces are added between operands when neither is a string.
	//It returns the number of bytes written and any write error encountered.
	fmt.Print("Enter the operator you want (+) , (-) , (*) , (/): ")

	// Scanln is similar to Scan, but stops scanning at a newline and after the
	//final item there must be a newline or EOF.
	fmt.Scanln(&operator) //  get the operator sign , & acts as a reference

	//Print formats using the default formats for its operands and writes to
	// standard output. Spaces are added between operands when neither is a string.
	//It returns the number of bytes written and any write error encountered.
	fmt.Print("Enter value of number1: ")

	//Scan scans text read from standard input, storing successive space-separated
	//values into successive arguments. Newlines count as space. It returns the
	//number of items successfully scanned. If that is less than the number of
	// arguments, err will report why.
	fmt.Scan(&n1) // get the n1 , & acts as a reference

	//Print formats using the default formats for its operands and writes to
	// standard output. Spaces are added between operands when neither is a string.
	//It returns the number of bytes written and any write error encountered.
	fmt.Print("Enter value of number2: ")

	//Scan scans text read from standard input, storing successive space-separated
	//values into successive arguments. Newlines count as space. It returns the
	//number of items successfully scanned. If that is less than the number of
	// arguments, err will report why.
	fmt.Scan(&n2) // get the n2 , & acts as a reference

	// Dial creates a client connection to the given target.
	//WithInsecure returns a DialOption which disables transport security for this
	//ClientConn. Note that transport security is required unless WithInsecure is set.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		//Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
		log.Fatalf("Failed to Dial!!: %s", err) // incase the dial to port fails
	}
	//close connection
	//Close tears down the ClientConn and all underlying connections.
	defer conn.Close()

	client := proto.NewCalculatorClient(conn)
	//We also pass a context.Context object which lets us change our RPC’s behavior if necessary,
	//such as time-out/cancel an RPC in flight.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(150)*time.Millisecond)
	defer cancel()

	//choose the operator
	switch operator {
	case "+":
		//Calling the  RPC
		//RPC call waits for the server to respond, and will either return a response or an error.
		//we call the method on the stub we got earlier.
		//In our method parameters we create and populate a request protocol buffer object (in our case AddRequest)
		addingResult, err := client.Add(ctx, &proto.AddRequest{Number1: n1, Number2: n2})

		// if there is an error
		if err != nil {
			// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
			log.Fatalf("There is an error with the addition: %s", err)
		}
		// Printf formats according to a format specifier and writes to standard
		// output. It returns the number of bytes written and any write error
		// encountered.
		fmt.Printf("%.0d + %.0d = %.0d", n1, n2, addingResult.Number1)
		// Println formats using the default formats for its operands and writes to
		// standard output. Spaces are always added between operands and a newline is
		// appended. It returns the number of bytes written and any write error
		// encountered.
		fmt.Println()

	case "-":
		//Calling the  RPC
		//RPC call waits for the server to respond, and will either return a response or an error.
		//we call the method on the stub we got earlier.
		//In our method parameters we create and populate a request protocol buffer object (in our case SubtractRequest)
		subtractResult, err := client.Subtract(ctx, &proto.SubtractRequest{Number1: n1, Number2: n2})

		if err != nil {
			// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
			log.Fatalf("There is an error with the subtraction: %s", err)
		}
		// Printf formats according to a format specifier and writes to standard
		// output. It returns the number of bytes written and any write error
		// encountered.
		fmt.Printf("%.0d - %.0d = %.0d", n1, n2, subtractResult.Number1)
		// Println formats using the default formats for its operands and writes to
		// standard output. Spaces are always added between operands and a newline is
		// appended. It returns the number of bytes written and any write error
		// encountered.
		fmt.Println()

	case "*":
		//Calling the  RPC
		//RPC call waits for the server to respond, and will either return a response or an error.
		//we call the method on the stub we got earlier.
		//In our method parameters we create and populate a request protocol buffer object (in our case MultiplyRequest)
		multiplyResult, err := client.Multiply(ctx, &proto.MultiplyRequest{Number1: n1, Number2: n2})

		if err != nil {
			// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
			log.Fatalf("There is an error with the multiplication: %s", err)
		}
		// Printf formats according to a format specifier and writes to standard
		// output. It returns the number of bytes written and any write error
		// encountered.
		fmt.Printf("%.0d * %.0d = %.0d", n1, n2, multiplyResult.Number1)
		// Println formats using the default formats for its operands and writes to
		// standard output. Spaces are always added between operands and a newline is
		// appended. It returns the number of bytes written and any write error
		// encountered.
		fmt.Println()

	case "/":
		//Calling the  RPC
		//RPC call waits for the server to respond, and will either return a response or an error.
		//we call the method on the stub we got earlier.
		//In our method parameters we create and populate a request protocol buffer object (in our case DivideRequest)
		divideResult, err := client.Divide(ctx, &proto.DivideRequest{Number1: n1, Number2: n2})

		if err != nil {
			// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
			log.Fatalf("There is an error with the division: %s", err)
		}
		// Printf formats according to a format specifier and writes to standard
		// output. It returns the number of bytes written and any write error
		// encountered.
		fmt.Printf("%.0d / %.0d = %.0d", n1, n2, divideResult.Number1)
		// Println formats using the default formats for its operands and writes to
		// standard output. Spaces are always added between operands and a newline is
		// appended. It returns the number of bytes written and any write error
		// encountered.
		fmt.Println()

	default:
		// Println formats using the default formats for its operands and writes to
		//standard output. Spaces are always added between operands and a newline is
		//appended. It returns the number of bytes written and any write error encountered.
		fmt.Println("Invalid operator!")
	}

}

func main() {
	ClientCalc()

}
