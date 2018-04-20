# Goueues
Work with simple queues like FIFO(First In First Out) and LIFO(Last In First Out).

## Funcs
Funcs than you can call when you create a Gonfiguration object:

```Go
- func New() *Fifo
- func(f *Fifo) Add(i interface{})
- func(f *Fifo) Pop() (i interface{})
- func(f *Fifo) Len() int
```

## Usage
You only need to create a variable and call the "New" method, it will return a pointer with the initialized fifo/lifo object. Now you can use all the functions of the object. An example of use:

```Go
	fifoExample := New() // Create new Object
	fifoExample.Add("Simple") // Add element to the queue
	fifoExample.Add("Fifo") // Add element to the queue
	fifoExample.Add("Example") // Add element to the queue
	fmt.Print(fifoExample.Pop()) // Simple
	fmt.Print(fifoExample.Pop()) // Fifo
	fmt.Println(fifoExample.Pop()) // Example

	fmt.Print(fifoExample.Pop()) // <Nil> because there aren`t more elements in the queue
```

### Note
This package is not thread-safe.
