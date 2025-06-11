/*
    handle whitespace 
*/
import "strconv"
func myAtoi(s string) int {

    newInt := 0

    whiteSpaceCheck := false
    allowSign := true
    sign := 1
    digit := 10
    
    numberHashmap := map[string]bool {
        "0":true,
        "1":true,
        "2":true,
        "3":true,
        "4":true,
        "5":true,
        "6":true,
        "7":true,
        "8":true,
        "9":true,
    }
    signHashmap := map[string]int {
        "-":-1,
        "+":1,
    }

    for _, s := range s {

        stringS := string(s)

        if whiteSpaceCheck == false {
            if stringS != " "{
                whiteSpaceCheck = true
            }
        }


        if whiteSpaceCheck == true {
            
            if allowSign{

                if si, foundSign := signHashmap[stringS]; foundSign {


                    allowSign = false
                    sign = si
                } else if _, found := numberHashmap[stringS]; found {
                    allowSign = false

                    num, _ := strconv.Atoi(stringS)
                    newInt = newInt*(digit) + num

                    roundupInt, isOverflow := checkOverflow(newInt*sign)

                    if isOverflow {
                        return roundupInt
                    }


                } else {
                    break
                }
            } else {

                if _, found := numberHashmap[stringS]; found {

                    num, _ := strconv.Atoi(stringS)
                    
                    newInt = newInt*(digit) + num

                    roundupInt, isOverflow := checkOverflow(newInt*sign)

                    if isOverflow {

                        return roundupInt
                    }


                } else {
                    break
                }
            }


        }
    }

    return newInt*sign


}


func checkOverflow(s int) (int, bool) {
    if s > 1<<31 -1 {
        return  1<<31 -1, true
    } else if s < -(1<<31){
        return -(1<<31), true
    }
    return s, false
}