func isValid(s string) bool {
    if len(s) % 2 == 1{
        return false
    }

    stk := []rune{}
    m := map[rune]rune{
        '(': ')',
        '[': ']',
        '{': '}',
    }

    for _, c := range s{
        _, ok := m[c]
        if ok {
            stk = append(stk, c)
        } else {
            if len(stk) == 0 {
                return false
            }
            if c != m[stk[len(stk) - 1]] {
                return false
            } else {
                stk = stk[:len(stk)-1]
            }
        }
    }

    if len(stk) == 0 {
        return true
    } else {
        return false
    }
}