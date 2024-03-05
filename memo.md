# 3章データ型
## No.17 8進数リテラルで混乱を招く

```golang
sum := 100 + 010
fmt.Println(sum) // 108
```

Goでは、0から始まる整数リテラルは8進数とみなされる。

## 8進数の使いどころ
ファイルの権限など。
0644はLinuxの特定の権限を表している。
```golang
os.OpenFile("memo.md", os.O_CREATE, 0644)
```

## No.18 整数のオーバーフローを無視する
カウンターなどの計算で

```go
var counter int32 = math.MaxInt32
counter++
fmt.Println(counter) //-2147483648
```

## No.19 浮動小数点を理解していない
floatは整数の小数点を表現できない問題を解決するための方法。
