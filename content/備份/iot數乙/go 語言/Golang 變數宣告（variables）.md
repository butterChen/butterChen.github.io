[[go語言]]
[[:=]]
# [Golang] 變數宣告（variables）

## Basic Go Type[​](https://pjchender.dev/golang/variables/#basic-go-type "Basic Go Type的直接連結")

Go 屬於強型別（Static Types）的語言，其中常見的基本型別包含 `bool`, `string`, `int`, `float64`, `map`。

> 變數宣告時**變數名稱放前面 變數型別放後面**

### 第一種宣告方式（最常用）：short declaration[​](https://pjchender.dev/golang/variables/#%E7%AC%AC%E4%B8%80%E7%A8%AE%E5%AE%A3%E5%91%8A%E6%96%B9%E5%BC%8F%E6%9C%80%E5%B8%B8%E7%94%A8short-declaration "第一種宣告方式（最常用）：short declaration的直接連結")

使用 `:=` 宣告，表示之前沒有進行宣告過。這是在 go 中最常使用的變數宣告的方式，因為它很簡潔。但因為在 package scope 的變數都是以 keyword 作為開頭，因此不能使用縮寫的方式定義變數（`foo := bar`），只能在 function 中使用，具有區域性（local variable）：

```
// 第一種宣告方式function main() {    foo := "Hello"    bar := 100  // 也可以簡寫成  foo, bar := "Hello", 100}// 等同於function main() {    var foo string    foo = "Hello"}
```

### 第二種宣告方式：variable declaration[​](https://pjchender.dev/golang/variables/#%E7%AC%AC%E4%BA%8C%E7%A8%AE%E5%AE%A3%E5%91%8A%E6%96%B9%E5%BC%8Fvariable-declaration "第二種宣告方式：variable declaration的直接連結")

使用時機主要是：

- 當你不知道變數的起始值
- 需要在 package scope 宣告變數
- 當為了程式的閱讀性，將變數組織在一起時

> ⚠️ 留意：在 package scope 宣告的變數會一直保存在記憶體中，直到程式結束才被釋放，因此應該減少在 package scopes 宣告變數

```
// 第二種宣告方式，在 main 外面宣告（全域變數），並在 main 內賦值var foo stringvar bar int// 可以簡寫成var (    foo string    bar int)function main() {  foo = "Hello"    bar = 100}
```

> ⚠️ **不建議把變數宣告在全域環境**

如果變數型別一樣的話，也可以簡寫成這樣：

```
func main() {    var c, python, java bool    fmt.Println(c, python, java)}
```

### 第三種宣告方式[​](https://pjchender.dev/golang/variables/#%E7%AC%AC%E4%B8%89%E7%A8%AE%E5%AE%A3%E5%91%8A%E6%96%B9%E5%BC%8F "第三種宣告方式的直接連結")

直接宣告並賦值：

```
// 第三種宣告方式，直接賦值var (  foo string = "Hello"    bar int = 100)
```

### 三種方式是一樣的[​](https://pjchender.dev/golang/variables/#%E4%B8%89%E7%A8%AE%E6%96%B9%E5%BC%8F%E6%98%AF%E4%B8%80%E6%A8%A3%E7%9A%84 "三種方式是一樣的的直接連結")

下面這兩種寫法是完全一樣的：

```
var <name> <type> = <value>var <name> := <value>
```

```
// var card string = "Ace of Spades"card := "Ace of Spades"
```

```
// var pi float = 3.14pi := 3.14
```

**只有在宣告變數的時候可以使用 `:=` 的寫法，如果要重新賦值的話只需要使用 `=`。**

## 注意事項[​](https://pjchender.dev/golang/variables/#%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A0%85 "注意事項的直接連結")

#### 錯誤：重複宣告變數[​](https://pjchender.dev/golang/variables/#%E9%8C%AF%E8%AA%A4%E9%87%8D%E8%A4%87%E5%AE%A3%E5%91%8A%E8%AE%8A%E6%95%B8 "錯誤：重複宣告變數的直接連結")

```
// 錯誤：重複宣告變數paperColor := "Green"paperColor := "Blue"
```

#### 正確：我們可以在 main 函式外宣告變數，但無法在 main 函式外賦值[​](https://pjchender.dev/golang/variables/#%E6%AD%A3%E7%A2%BA%E6%88%91%E5%80%91%E5%8F%AF%E4%BB%A5%E5%9C%A8-main-%E5%87%BD%E5%BC%8F%E5%A4%96%E5%AE%A3%E5%91%8A%E8%AE%8A%E6%95%B8%E4%BD%86%E7%84%A1%E6%B3%95%E5%9C%A8-main-%E5%87%BD%E5%BC%8F%E5%A4%96%E8%B3%A6%E5%80%BC "正確：我們可以在 main 函式外宣告變數，但無法在 main 函式外賦值的直接連結")

```
// 正確：我們可以在 main 函式外宣告變數，但無法在 main 函式外賦值package mainimport "fmt"var deckSize intfunc main() {  deckSize = 50  fmt.Println(deckSize)}
```

#### 錯誤：無法在 main 函式外賦值[​](https://pjchender.dev/golang/variables/#%E9%8C%AF%E8%AA%A4%E7%84%A1%E6%B3%95%E5%9C%A8-main-%E5%87%BD%E5%BC%8F%E5%A4%96%E8%B3%A6%E5%80%BC "錯誤：無法在 main 函式外賦值的直接連結")

```
// 錯誤：但無法在 main 函式外賦值package mainimport "fmt"// syntax error: non-declaration statement outside function bodydeckSize := 20func main() {  fmt.Println(deckSize)}
```

#### 錯誤：變數需要先宣告完才能使用[​](https://pjchender.dev/golang/variables/#%E9%8C%AF%E8%AA%A4%E8%AE%8A%E6%95%B8%E9%9C%80%E8%A6%81%E5%85%88%E5%AE%A3%E5%91%8A%E5%AE%8C%E6%89%8D%E8%83%BD%E4%BD%BF%E7%94%A8 "錯誤：變數需要先宣告完才能使用的直接連結")

```
// 錯誤：變數需要先宣告完才能使用package mainimport "fmt"func main() {  deckSize = 52            // undefined: deckSize  fmt.Println(deckSize)    // undefined: deckSize}
```

## 常數（constant）[​](https://pjchender.dev/golang/variables/#%E5%B8%B8%E6%95%B8constant "常數（constant）的直接連結")

##### keywords: `iota`[​](https://pjchender.dev/golang/variables/#keywords-iota "keywords-iota的直接連結")

使用 `:=` 或 `var` 所宣告的會是變數，若需要宣告常數，需要使用 `const`：

```
const (    Monday = 1    Tuesday = 2    Wednesday = 3  // ...)// 可以簡寫成// iota 預設從 0 開始，後面的變數自動加一const (    Monday = iota + 1    Tuesday    Wednesday  // ...)
```

## 參考[​](https://pjchender.dev/golang/variables/#%E5%8F%83%E8%80%83 "參考的直接連結")

- [Go 語言基礎實戰 (開發, 測試及部署)](https://www.udemy.com/course/golang-fight/) @ Udemy by [Bo-Yi Wu](https://www.udemy.com/user/bo-yi-wu-2/)
- [Go (Golang): The Complete Bootcamp](https://www.udemy.com/course/learn-go-the-complete-bootcamp-course-golang/) @ Udemy by [Jose Portilla](https://www.udemy.com/user/joseportilla/)
- [Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/) @ Udemy by [Stephen Grider](https://www.udemy.com/user/sgslo/)