func expandAroundCenter(s string, left int, right int)(int, int, int){
    for left >= 0 && right <= len(s) - 1 && s[left] == s[right]{
        left --
        right ++
    }
    return right - left - 1, left + 1, right -1 
}

func longestPalindrome(s string) string {
    // expand around center
    start, end := 0, 0
    maxLength := 0

    for i, _ := range s {
        // for odd palindromic, "aba"
        oddPalindrome, oddLeft, oddRight := expandAroundCenter(s, i, i)
        // for even palindromc, "abba"
        evenPalindrome, evenLeft, evenRight := expandAroundCenter(s, i, i+1)
        m := max(oddPalindrome, evenPalindrome)

        if m > maxLength{
            maxLength = m
            if oddPalindrome > evenPalindrome{
                start = oddLeft
                end = oddRight
            } else {
                start = evenLeft
                end = evenRight
                
            }

        }


    }

    return s[start:end+1]

}