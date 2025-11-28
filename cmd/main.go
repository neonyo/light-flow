package main

import (
	"fmt"

	"github.com/neonyo/light-flow/flow"
)

func First(step flow.Step) (any, error) {
	if input, exists := step.Get("input"); exists {
		fmt.Printf("[Step: %s] got input: %v\n", step.Name(), input)
	}
	step.Set("key", "value")
	return "result", nil
}

func Second(step flow.Step) (any, error) {
	if value, exists := step.Get("key"); exists {
		fmt.Printf("[Step: %s] got key: %v\n", step.Name(), value)
	}
	if result, exists := step.Result(step.Dependents()[0]); exists {
		fmt.Printf("[Step: %s] got result: %v\n", step.Name(), result)
	}
	return nil, nil
}

func ErrorStep(step flow.Step) (any, error) {
	if value, exists := step.Get("key"); exists {
		fmt.Printf("[Step: %s] got key: %v\n", step.Name(), value)
	} else {
		fmt.Printf("[Step: %s] cannot get key \n", step.Name())
	}
	return nil, fmt.Errorf("execution failed")
}

func ErrorHandler(step flow.Step) (bool, error) {
	if step.Has(flow.Failed) {
		fmt.Printf("[Step: %s] has failed\n", step.Name())
	} else {
		fmt.Printf("[Step: %s] succeeded\n", step.Name())
	}
	return true, nil
}

func init() {
	process := flow.FlowWithProcess("Example")
	process.Follow(First, Second)
	process.Follow(ErrorStep)
	process.AfterStep(true, ErrorHandler)
}

func main() {
	flow.DoneFlow("Example", map[string]any{"input": "Hello world"})
}
