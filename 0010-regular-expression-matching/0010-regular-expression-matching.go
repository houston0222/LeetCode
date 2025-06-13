/*
the difficult part is the *, because it will match zero or more of the preceding element, 
so ".*" will match everything
because of that we can't perform index by index comparsion.

I will loop the p, and p length is <= s length, so I will have compensatedSi to get correct campersion.

*/ 

func isMatch(s string, p string) bool {

    isMatched := true

    si := 0
    compensatedSi := 0
    pi := 0
    
    
    for pi < len(p) && si + compensatedSi < len(s) {
        strsch := string(s[si + compensatedSi])
        strpch := string(p[pi])

        if strpch == "."{
            isMatched = true
            si ++ 
            pi ++
        } else if strpch == "*" {
            previousStrpch := string(p[pi - 1])

            sequentialMatched := true

            for sequentialMatched == true && si + compensatedSi < len(s){
                strsch := string(s[si + compensatedSi])

                if previousStrpch == "."{
                    compensatedSi ++  
                } else {
                    if (strsch == previousStrpch){
                        compensatedSi ++  
                    } else {
                        sequentialMatched = false
                    }
                }
            }

            if sequentialMatched == true{
                si ++ 
                pi ++
            } else {
                isMatched = false
                break
            }

        } else {
            if strsch != strpch{
                isMatched = false
                break
            } else {
                si ++ 
                pi ++
            }

        }

    }


    if si+compensatedSi-1 != len(s){

        return false
    }

    return isMatched
    
}