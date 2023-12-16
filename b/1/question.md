# 宇宙探査ロボットの経路計算

## 問題概要
宇宙探査ロボットが未知の惑星を探索しています。このロボットは正方形のグリッド上を動きます。グリッドの各セルには、地形に応じたコストが割り当てられています。ロボットがスタート地点から目的地点まで移動する際に、最小コストの経路を計算するプログラムを作成してください。

## 関数の仕様
- 関数名: CalculateOptimalPath
- 引数:
  - `grid [][]int`: 正方形のグリッド。各セルには移動コスト（正の整数）が割り当てられている。
  - `startX, startY int`: スタート地点のX座標とY座標（0-indexed）。
  - `endX, endY int`: 目的地点のX座標とY座標（0-indexed）。
- 戻り値:
  - `totalCost int`: 最小コストの経路の合計コスト。
  - `err error`: 引数が無効な場合に返されるエラー。

## 入出力例
- 正常ケース1:
  - 入力: grid = [[1, 3, 1], [1, 5, 1], [4, 2, 1]], startX = 0, startY = 0, endX = 2, endY = 2
  - 出力: totalCost = 7
- 正常ケース2:
  - 入力: grid = [[1, 2, 3], [4, 5, 6], [7, 8, 9]], startX = 0, startY = 0, endX = 2, endY = 2
  - 出力: totalCost = 12
- エラーケース1:
  - 入力: grid = [[-1, 2], [3, 4]], startX = 0, startY = 0, endX = 1, endY = 1
  - 出力: err = "invalid grid value"
- エラーケース2:
  - 入力: grid = [[1, 2], [3, 4]], startX = -1, startY = 0, endX = 1, endY = 1
  - 出力: err = "invalid start or end coordinates"

## 注意事項
- gridは必ず正方形でなければなりません。そうでない場合、`error`を返してください。
- gridの値が負の数の場合、`invalid grid value`というエラーメッセージを返してください。
- startX, startY, endX, endYは、gridの範囲内に収まっている必要があります。範囲外の場合は`invalid start or end coordinates`というエラーメッセージを返してください。
- Goの標準パッケージのみを使用してください。

この問題はダイクストラのアルゴリズムやA*アルゴリズムの理解に役立ちます。また、エラーハンドリングと条件チェックの練習にもなります。
