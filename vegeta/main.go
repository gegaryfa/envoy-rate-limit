package main

import (
	"fmt"
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func closeEnough(a, b float64) bool {
	epsilon := 0.03
	return (a-b) < epsilon && (b-a) < epsilon
}

func runTest(desc string, okPct float64, tgts ...vegeta.Target) {
	rate := vegeta.Rate{Freq: 10, Per: time.Second}
	duration := 10 * time.Second

	targeter := vegeta.NewStaticTargeter(tgts...)
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	fmt.Println(desc)
	for res := range attacker.Attack(targeter, rate, duration, "test") {
		metrics.Add(res)
	}
	metrics.Close()

	if closeEnough(metrics.Success, okPct) {
		fmt.Printf("OK! Got %0.2f which was close enough to %0.2f\n", metrics.Success, okPct)
	} else {
		fmt.Printf("Error: Got %0.2f which was too far from %0.2f\n", metrics.Success, okPct)
	}

	for code := range metrics.StatusCodes {
		fmt.Printf("\t%s: %d\n", code, metrics.StatusCodes[code])
	}
}

func main() {
	// An authenticated path
	authedTargetA := vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8888/header",
		Header: http.Header{
			"Authorization": []string{"Bearer fooA"},
		},
	}

	// Same path as A, different user
	authedTargetB := vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8888/header",
		Header: http.Header{
			"Authorization": []string{"Bearer fooB"},
		},
	}

	// Same path as A, simulating different user
	authedTargetC := vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8888/header",
		Header: http.Header{
			"Authorization": []string{"Bearer barC"},
		},
	}

	runTest("single authed path header fooA, target 10rpm", 0.1, authedTargetA)
	runTest("single authed path header fooB, target 10rpm", 0.10, authedTargetB)
	runTest("single authed path header barC, target 10rpm", 0.10, authedTargetC)
}
