package main

import "fmt"

// Pre-Generics Approaches : tyring to handle multiple map types using `any`
//Problems:
//- Boilerplate: each map type requires duplicatign the for loop
//- Type safety only happens at runtime - compile-time errors are gone
//- Returns type is []any -caller must type-assert each key
//- Passing a wrong type compiles but fails at runtime
func getKeys(m any) ([]any, error) {
	switch t := m.(type) {
	default:
		//Unknown type discovered too late = runtime error
		return nil, fmt.Errorf("unknown type: %T", t)
	case map[string]int:
		var keys []any
		for k := range t {
			//loses type info - becomes `any`
			keys = append(keys, k)
		}
		return keys, nil
	case map[int]string:
		//Entire loop duplicated just for a different map type
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	}
}

//BAD: Using generics incorectly - adds no real benefit
//T- is only used to call a method that already exists on io.Writer
//Just use io.Writer directly instead of defining a generic type constraint
func writeData[T interface{ Write([]byte) (int, error) }](w T, data []byte) {
	_, _ = w.Write(data)
}

func main() {
	m1 := map[string]int{"one": 1, "two": 2}

	keys1, err := getKeys(m1)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Caller must type-assert each key-unsafe
	for _, k := range keys1 {
		fmt.Println(k.(string))
	}

	m2 := map[int]string{1: "one", 2: "two"}

	keys2, err := getKeys(m2)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range keys2 {
		fmt.Println(k.(int))
	}

	//Passing a completely wrong type - compiles fine, fails at runtime
	_, err = getKeys("oops")
	fmt.Println(err)
}

//Compile time - Errors or checks that the compiler can catch before the program runs
//Happens when the code is being translated into machine code(compile)

//Run time - Errors or checks that occur when the program is running
//Often happen due to unexpected input, invalid operations or misuse of interfaces
