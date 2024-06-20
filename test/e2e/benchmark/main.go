package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	logger := log.New(os.Stdout, "test-e2e-benchmark", log.LstdFlags)

	tests := []Test{
		{"TwoNodeSimple", TwoNodeSimple},
	}

	// check the test name passed as an argument and run it
	specificTestFound := false
	for _, arg := range os.Args[1:] {
		for _, test := range tests {
			if test.Name == arg {
				runTest(logger, test)
				specificTestFound = true
				break
			}
		}
	}

	if !specificTestFound {
		logger.Println("No particular test specified. Running all tests.")
		logger.Println("go run ./test/e2e/benchmark <test_name> to run a specific test")
		logger.Printf("Valid tests are: %s\n\n", getTestNames(tests))
		// if no specific test is passed, run all tests
		for _, test := range tests {
			runTest(logger, test)
		}
	}
}

type TestFunc func(*log.Logger) error

type Test struct {
	Name string
	Func TestFunc
}

func runTest(logger *log.Logger, test Test) {
	logger.Printf("=== RUN %s", test.Name)
	err := test.Func(logger)
	if err != nil {
		logger.Fatalf("--- ERROR %s: %v", test.Name, err)
	}
	logger.Printf("--- ✅ PASS: %s \n\n", test.Name)
}

func getTestNames(tests []Test) string {
	testNames := make([]string, 0, len(tests))
	for _, test := range tests {
		testNames = append(testNames, test.Name)
	}
	return strings.Join(testNames, ", ")
}
