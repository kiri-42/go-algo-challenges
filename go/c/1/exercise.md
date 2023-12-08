# 文字列の回文判定

## 問題概要
与えられた文字列が回文（前から読んでも後ろから読んでも同じになる文字列）であるかを判定する関数を作成します。

## 関数の仕様
- 関数名: `IsPalindrome`
- 引数:
  - `s string`: 調べる文字列。英数字のみを含むとします。
- 戻り値:
  - `isPal bool`: 文字列が回文の場合はtrue、そうでない場合はfalseを返します。
  - `err error`: 引数が英数字以外を含む場合にエラーを返します。

## 入力例
- `IsPalindrome("racecar")` -> `true, nil`
- `IsPalindrome("hello")` -> `false, nil`
- `IsPalindrome("あいうえお")` -> `false, エラー: 英数字以外の文字が含まれています`

## 注意事項
- 標準パッケージのみを使用します。
- 文字列に英数字以外が含まれる場合はエラーとして処理します。
- 文字列は空であっても構いませんが、空の場合はfalseを返します。

この練習問題は基本的な文字列操作とエラーハンドリングを学ぶのに役立ちます。また、Go言語における簡単な関数の作成方法と、引数の検証に関する知識も習得できます。