/*
this use recursion + memo to solve, because * need backtrack to try multiple way

dp(si, pi) means check if s[si:] match p[pi:]

we store result in map so we don't repeat same check

if p[pi+1] is '*', we can:
1. skip the char and * (means 0 match)
2. if current s[si] match p[pi], try match next s (keep p same)

if no *, just check if current char match, and go next

base case: if pattern finished, s must also finished
*/

func isMatch(s string, p string) bool {
    memo := make(map[string]bool)

    var dp func(si, pi int) bool

    dp = func(si, pi int) bool {
        key := fmt.Sprintf("%d,%d", si, pi)
        if val, ok := memo[key]; ok {
            return val
        }

        if pi == len(p) {
            return si == len(s)
        }

        match := false

        currentMatch := si < len(s) && (s[si] == p[pi] || p[pi] == '.')

        if pi+1 < len(p) && p[pi+1] == '*' {
            match = dp(si, pi+2) || (currentMatch && dp(si+1, pi))
        } else {
            if currentMatch {
                match = dp(si+1, pi+1)
            }
        }

        memo[key] = match
        return match
    }

    return dp(0, 0)
}
