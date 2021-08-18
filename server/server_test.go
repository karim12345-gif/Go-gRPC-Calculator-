package main

import (
	"context" //Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	"testing" //Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the "go test" command, which automates execution of any function of the form

	proto "github.com/karim12345-gif/grpc_codingtest/service"
)

//*** Addition ***
type addTest struct {
	arg1, arg2, expected int32
}

var addTests = []addTest{
	{2, 3, 5},
	{10, 10, 20},
	//{10, 10, 1}, // will result in fail if you uncomment it
}

// *** subtraction ***
type subtractTest struct {
	arg1, arg2, expected int32
}

var subtractTests = []subtractTest{
	{10, 5, 5},
	{10, 10, 0},
	//{50, 40, 1}, // will return a fail
}

// *** Multiplication ***
type multiplicationTest struct {
	arg1, arg2, expected int32
}

var multiplicationTests = []multiplicationTest{
	{10, 5, 50},
	{2, 10, 20},
	//{10,10,100000}, //will return a fail
}

// *** Division ***
type divideTest struct {
	arg1, arg2, expected int32
}

var divideTests = []divideTest{
	{50, 5, 10},
	{10, 2, 5},
	//{100, 20, 25}, // will return a fail
}

// **addition**
func TestAddition(t *testing.T) {

	for _, test := range addTests {
		s := &Server{} //reference to the functions
		outputadd := proto.AddRequest{Number1: test.arg1, Number2: test.arg2}
		//Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline.
		// It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
		res, _ := s.Add(context.Background(), &outputadd)
		expectadd := test.expected

		// if the given number is not equal to the expected number, the test will fail
		if res.Number1 != expectadd {
			//Errorf is equivalent to Logf followed by Fail.
			t.Errorf("got %q, wanted %q", &outputadd, expectadd)
		}
	}

}

// **subtraction**
func TestSubtraction(t *testing.T) {
	for _, test := range subtractTests {
		s := &Server{} //reference to the functions
		outputsub := proto.SubtractRequest{Number1: test.arg1, Number2: test.arg2}
		//Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline.
		// It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
		res, _ := s.Subtract(context.Background(), &outputsub)
		// if the given number is not equal to the expected number, the test will fail
		expectsub := test.expected

		if res.Number1 != expectsub {
			//Errorf is equivalent to Logf followed by Fail.
			t.Errorf("got %q, wanted %q", &outputsub, expectsub)
		}

	}
}

// **Multiplication**
func TestMultiply(t *testing.T) {
	for _, test := range multiplicationTests {
		s := &Server{} //reference to the functions
		outputMultiply := proto.MultiplyRequest{Number1: test.arg1, Number2: test.arg2}
		//Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline.
		// It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
		res, _ := s.Multiply(context.Background(), &outputMultiply)
		// if the given number is not equal to the expected number, the test will fail
		expectmulti := test.expected

		if res.Number1 != expectmulti {
			//Errorf is equivalent to Logf followed by Fail.
			t.Errorf("got %q, wanted %q", &outputMultiply, expectmulti)
		}

	}
}

// **Division**
func TestDivide(t *testing.T) {
	for _, test := range divideTests {
		s := &Server{} //reference to the functions
		outputDivide := proto.DivideRequest{Number1: test.arg1, Number2: test.arg2}
		//Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline.
		// It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
		res, _ := s.Divide(context.Background(), &outputDivide)
		// if the given number is not equal to the expected number, the test will fail
		expectdiv := test.expected

		if res.Number1 != expectdiv {
			//Errorf is equivalent to Logf followed by Fail.
			t.Errorf("got %q, wanted %q", &outputDivide, expectdiv)
		}

	}
}
