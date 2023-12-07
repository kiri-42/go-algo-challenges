```go
package main

import (
	"errors"
	"unicode"
)

// IsPalindrome は与えられた文字列が回文かどうかを判定します。
func IsPalindrome(s string) (bool, error) {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false, errors.New("エラー: 英数字以外の文字が含まれています")
		}
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false, nil
		}
	}
	return true, nil
}
```
