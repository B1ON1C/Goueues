package fifo

// Fifo holds info about the object
type Fifo struct{
	FifoList []interface{}
}

// New creates new Fifo pointer and returns it
func New() *Fifo{
	return &Fifo{}
}

// Add add to the *Fifo list new interface
func(f *Fifo) Add(i interface{}){
	f.FifoList = append(f.FifoList, i)
}

// Pop returns the first item of the list *Fifo
func(f *Fifo) Pop() (i interface{}) {
	if len(f.FifoList)>0{
		i = f.FifoList[0]
		f.FifoList = f.FifoList[1:]

		return
	}
	return nil
}

// Len returns int about the actual size of the *Fifo list
func(f *Fifo) Len() int{
	return len(f.FifoList)
}
