package lifo

// Lifo holds info about the object
type Lifo struct{
	LifoList []interface{}
}

// New creates new Lifo pointer and returns it
func New() *Lifo{
	return &Lifo{}
}

// Add add to the *Lifo list new interface
func(f *Lifo) Add(i interface{}){
	f.LifoList = append(f.LifoList, i)
}

// Pop returns the first item of the list *Lifo
func(f *Lifo) Pop() (i interface{}) {
	if len(f.LifoList)>0{
		i = f.LifoList[(f.Len()-1)]
		f.LifoList = f.LifoList[:(f.Len()-1)]

		return
	}
	return nil
}

// Len returns int about the actual size of the *Lifo list
func(f *Lifo) Len() int{
	return len(f.LifoList)
}
