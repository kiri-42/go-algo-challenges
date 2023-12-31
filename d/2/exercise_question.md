# 数字の合計

## 問題概要
この問題では、与えられた整数の配列の中で、特定の範囲（インデックスの開始と終了）にある数値の合計を計算します。

## 関数の仕様
- 関数名: `sumRange`
- 引数:
  - `numbers`: `[]int`型の整数配列。要素数は1以上100以下。
  - `start`: `int`型で、合計を開始するインデックス。有効な値は0以上`len(numbers)-1`以下。
  - `end`: `int`型で、合計を終了するインデックス。有効な値は`start`以上`len(numbers)-1`以下。
- 戻り値:
  - `sum int`: 指定された範囲の合計値
  - `err error`: エラーがある場合に返される

## 入力例
- 正常なケース
  - `numbers`: `[1, 2, 3, 4, 5]`, `start`: `1`, `end`: `3` -> `sum`: `9`
  - `numbers`: `[10, 20, 30]`, `start`: `0`, `end`: `2` -> `sum`: `60`
- エラーケース
  - `numbers`: `[1, 2, 3]`, `start`: `2`, `end`: `1` -> `err`: `error`
  - `numbers`: `[1, 2, 3, 4, 5]`, `start`: `-1`, `end`: `4` -> `err`: `error`
  - `numbers`: `[1, 2, 3, 4, 5]`, `start`: `1`, `end`: `5` -> `err`: `error`

## 注意事項
- `start`が`end`より大きい場合、または`start`または`end`が無効な範囲にある場合、`err`として`error`を返します。
- Goの標準パッケージのみを使用します。

この問題は基本的な配列操作とエラーハンドリングの練習に適しており、プログラミングの基本を学ぶのに役立ちます。
