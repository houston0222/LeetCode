func romanToInt(s string) int {
    single := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    total := 0
    n := len(s)

    for i := 0; i < n; i++ {
        // If this is not the last character and the next value is larger, it's a subtractive case
        if i+1 < n && single[s[i]] < single[s[i+1]] {
            total += single[s[i+1]] - single[s[i]]
            i++ // skip the next character
        } else {
            total += single[s[i]]
        }
    }

    return total
}
