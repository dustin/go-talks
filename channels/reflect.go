package main

import (
	"reflect"
)

func dynamicSendToAny(nodes []node, task thingToSend) {
	// Dynamically build a select{}
	cases := []reflect.SelectCase{}
	for _, node := range nodes {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(node.channel), // HL
			Send: reflect.ValueOf(thing),
		})
	}

	// Execute the select{} (ignore the results since we're not receiving)
	to, _, _ := reflect.Select(cases) // HL

	// Exactly one node has received our task.
	fmt.Printf("Sent to %v\n", nodes[to])
}
