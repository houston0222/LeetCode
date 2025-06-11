/*
    Possible solution:
    1. convert it to string and reverse loop it to make new string and convert to new int
        but string is immutable and require memory allocation
    2. use math is better because it has no memory allocation and safe overflow check

*/
func reverse(x int) int {

    reverseInteger := 0

    for x != 0 {
        remainder := x % 10

        reverseInteger = reverseInteger*10 + remainder

        if reverseInteger > 1<<31 -1 || reverseInteger < -(1<<31 -1){
            return 0
        }

        x /= 10
        
    }

    return reverseInteger
    
}