package main
import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(item string){
	s.items = append(s.items, item)	
}


func (s *Stack) Pop() string {
	sizeOfStack := len(s.items)
	
	if sizeOfStack == 0 {
		return "NOTHING LEFT TO POP OFF"
	}

	item, items := s.items[sizeOfStack-1], s.items[:sizeOfStack-1]
	s.items = items
	return item
}
func main(){
 s := Stack{}

 s.Push("a")
 s.Push("e")
 s.Push("i")
 s.Push("o")
 s.Push("u")


 fmt.Println(s.items)
 fmt.Println("pop it all")

 fmt.Println(s.Pop())
 fmt.Println(s.Pop())
 fmt.Println(s.Pop())
 fmt.Println(s.Pop())
 fmt.Println(s.Pop())
}