package main;

// method push and pop

type Stack struct {
	items []int
	// length int
}


func (s *Stack) Push(value int) {
	s.items = append(s.items, value); 
	// s.length++
}

func (s *Stack) Pop() (value int) {	
	
	if len(s.items) == 0 {
		return -1;
	}

	item, items := s.items[len(s.items)-1], s.items[0:len(s.items)-1]
	s.items = items;
	return item;
}


func main(){
	stack := &Stack{}

	stack.Push(0)
	stack.Push(1)
	stack.Push(2) 
	stack.Push(3)

	println(stack.Pop())
	println(stack.Pop())
	println(stack.Pop())
	println(stack.Pop())
}