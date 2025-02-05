package learn_test

import (
	"fmt"
)

func ArrayLoop() {
	var arr [5]int // Define
	arr[0] = 10    // Add (assign values to indices)
	arr[1] = 20    // Add
	arr[1] = 30    // Edit

	fmt.Println("Iterating over array:")
	for i, v := range arr {
		fmt.Printf("Index %d, Value %d\n", i, v)
	}
	// Expected output
	// Output:
	// Index 0, Value 10
	// Index 1, Value 20
	// Index 2, Value 30
}

func SliceLoop() {
	index := 1
	slice := []int{10, 20, 30}                        // Define
	slice = append(slice, 40)                         // Add
	slice[index] = 50                                 // Edit
	slice = append(slice[:index], slice[index+1:]...) // Remove specific index

	for i, v := range slice {
		fmt.Printf("Index %d, Value %d\n", i, v)
	}
	// Expected output
	// Output:
	// Slice after adding: [10 20 30 40]
	// Slice after editing: [10 50 30 40]
	// Slice after removing: [10 30 40]
	// Iterating over slice:
	// Index 0, Value 10
	// Index 1, Value 30
	// Index 2, Value 40
}

func MapLoop() {
	m := map[string]int{"one": 1, "two": 2}
	m["three"] = 3   // Add
	m["two"] = 20    // Edit
	delete(m, "one") // Delete

	for key, value := range m {
		fmt.Printf("Key %s, Value %d\n", key, value)
	}
	// Key three, Value 3
	// Key two, Value 20
}
