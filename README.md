# GoHashLib

This project implements a hashmap in Go that allows for keys of non-comparable types, including structs, by utilizing
the `encoding/json` package's `Marshal` method. Additionally, it provides a Set implementation based on the same hashmap
implementation.

### Features

- **Support for Non-Comparable Keys**: This hashmap implementation allows keys of types that are not directly comparable
  in Go, such as structs.
- **Simple Set Implementation**: The project also includes a Set data structure built on top of the hashmap, providing a
  convenient way to work with unique elements.
- **Marshaling for Keys**: Keys are converted to strings using the `Marshal` method from the `encoding/json` package,
  ensuring consistent hashing and retrieval.

### Usage

To use the hashmap and Set implementations in your Go project:

1. Clone this repository.
2. Import the `hashmap` package into your Go code.
3. Instantiate a new Map or Set using the provided constructors.
4. Add, retrieve, or delete elements as needed.

### Example

```go
package main

import (
	"fmt"
	"github.com/pietroagazzi/gohashlib/hashmap"
)

func main() {
	// Create a new hashmap
	myMap := hashmap.NewMap[any, string](2, 3)

	// Add elements to the map
	myMap.Set(MyStruct{ID: 1}, "Value1")
	myMap.Set("NonComparableKey", "Value2")

	// Retrieve elements from the map
	value1, _ := myMap.Get(MyStruct{ID: 1})
	value2, _ := myMap.Get("NonComparableKey")

	fmt.Println("Value1:", value1)
	fmt.Println("Value2:", value2)
}

type MyStruct struct {
	ID int
}

```

### Limitations

- **Performance Overhead**: Marshaling keys to strings incurs a performance overhead compared to native Go maps.
- **Memory Overhead**: The conversion of keys to strings may increase memory usage, especially for large datasets.
- **Marshaling Constraints**: Keys must be able to be marshaled into a string representation using `encoding/json`,
  limiting the types that can be used as keys.

### Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please open an issue or
submit a pull request.

---

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

