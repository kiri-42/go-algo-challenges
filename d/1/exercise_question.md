# 数字の合計

## 問題概要
この問題では、与えられた一連の数字の合計を計算します。数値はカンマ区切りで渡され、それらの合計を返す必要があります。

## 関数の仕様
- 関数名: `SumNumbers`
- 引数:
  - `numbers string`: カンマ区切りの数字の文字列。例: "1,2,3"
- 戻り値:
  - `sum int`: 数字の合計
  - `err error`: 数字の解析中にエラーが発生した場合に返す

## 入力例
- `numbers = "1,2,3"` -> `sum = 6`, `err = nil`
- `numbers = "4,5,-6"` -> `sum = 3`, `err = nil`
- `numbers = "a,2,3"` -> `sum = 0`, `err = error("invalid number: a")`

## 注意事項
- 数字の解析には `strconv.Atoi` を使用します。
- 数値以外の文字が含まれている場合、`invalid number: [含まれている不正な文字]` というエラーメッセージを持つ `error` を返します。
- エラーメッセージは英語で記述します。
- Goの標準パッケージのみをインポートしてください。

このDランクの問題は、文字列処理とエラーハンドリングの基本を学ぶのに適しています。
