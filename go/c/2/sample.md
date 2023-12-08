```go
package main

import (
    "errors"
)

// FindMaxValue は数列の中の最大値を探索します。
func FindMaxValue(numbers []int) (int, error) {
    if len(numbers) == 0 {
        return 0, errors.New("empty slice")
    }

    maxValue := numbers[0]
    for _, num := range numbers {
        if num < -1000 || num > 1000 {
            return 0, errors.New("invalid value in slice")
        }
        if num > maxValue {
            maxValue = num
        }
    }

    return maxValue, nil
}

```
