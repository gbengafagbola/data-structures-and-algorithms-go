package main

func isPowerOfTwo(n int) bool {
    if n == 0 {
        return false
    }
    for n%2 == 0 {
        n /= 2
    }
   return n == 1
}

// power of 3 
// same logic apply

func isPowerOfThree(n int) bool {
    if n == 0 {
        return false
    }

    for n%3 == 0 {
        n /= 3
    }
    
   return n == 1
}