/*
    we use hashmap to store the roman pair then prevent nested loop
    because 1 <= num <= 3999,
    i use digit 1000, and /10 every loop until the digit == 0.
    if 9 or 4, then *10 or *5 to find the roman pair, minus the *1
    if not just loop num/digit
*/

func intToRoman(num int) string {

    romanHashMap := map[int]string{
        1000:"M",
        500:"D",
        100:"C",
        50:"L",
        10:"X",
        5:"V",
        1:"I",
    }

    roman := ""

    digit := 1000

    for digit != 0 {

        n := num/digit
        loopN := n
        fmt.Println(num, digit, n)

        if n == 9 {
            roman += romanHashMap[digit]
            roman += romanHashMap[digit*10]
            loopN = 0

        } else if n == 4{
            roman += romanHashMap[digit]
            roman += romanHashMap[digit*5]
            loopN = 0
        } else if n >= 5 {
            roman += romanHashMap[digit*5]
            loopN = loopN - 5
        } 

        for i:=0; i < loopN; i++{
            roman += romanHashMap[digit]
        }

        num = num - n*digit
        digit /= 10

    }

    return roman
    
}