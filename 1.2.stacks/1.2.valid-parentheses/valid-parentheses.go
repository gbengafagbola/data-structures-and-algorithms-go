package main

func isValid(s string) bool {

	brackets := map[rune]rune{
	   '(': ')',
	   '{': '}',
	   '[': ']',
   }
   
   stack := []rune{}
   for _, digit := range s {
	   if _, ok := brackets[digit]; ok {
		   stack = append(stack, digit)
	   } else {
		   if len(stack) == 0 {
			   return false
		   }
		   lastOpen := stack[len(stack)-1]
		   if brackets[lastOpen] != digit {
		   return false
	   }
		   stack = stack[:len (stack) -1]
	   }
	
   }
   
   if len(stack) == 0 {
	return true
   } else {
	   return false 
   }
   }