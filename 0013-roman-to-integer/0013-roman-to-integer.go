func romanToInt(s string) int {

    hashMap := map[string]int {
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,

        "IV":4,
        "IX":9,

        "XL":40,
        "XC":90,

        "CD":400,
        "CM":900,

    }

    i := 0
    num := 0

    for i < len(s){
        currentCh := string(s[i])

        if i + 1 < len(s){

            nextCh := string(s[i+1])

            if v, found := hashMap[currentCh+nextCh]; found {
                num += v
                i ++ 
            } else {
                num += hashMap[currentCh]
            }
            
        } else {
            num += hashMap[currentCh]
        }


        i++
    }

    return num
}