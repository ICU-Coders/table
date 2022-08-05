
 <p align="center" >
   <img src="https://raw.githubusercontent.com/ICU-Coders/IconLib/master/icon.jpg" alt="ICU-Coders" title="ICU-Coders">
 </p>

# Terminal Table

![example](./example.png)

### How To Use

```shell
go get github.com/ICU-Coders/table
```
```go
import "github.com/ICU-Coders/table"

table.Show([]string{"Module", "Type", "Path", "Author"}, [][]string{
    {"1", "2", "3", "4"},
    {"1", "2", "3", "4"},
    {"1", "2", "3", "4"},
    {"1", "2", "3", "4"},
})
```


### Config
```go
var MaxCellWidth = 40
var LineEndTag = "+"
var LineBody = "-"
var LineDivider = "|"
```

Example
```go
func TestShow(t *testing.T) {
	LineEndTag = "*"
	LineBody = "="
	LineDivider = "/"
	Show([]string{"Module", "Type", "Path", "Author"}, [][]string{
		{"1111", "2", "3", "4"},
		{"1", "2", "3", "4"},
		{"1", "2", "3", "4"},
		{"1", "2", "3", "4"},
	})
}
```
Display
```
*========*======*======*========*
/ Module / Type / Path / Author /
*========*======*======*========*
/ 1111   / 2    / 3    / 4      /
/ 1      / 2    / 3    / 4      /
/ 1      / 2    / 3    / 4      /
/ 1      / 2    / 3    / 4      /
*========*======*======*========*
```

### Auto warp

```go
MaxCellWidth = 20 // default 40
Show([]string{"Module", "Type", "Path", "Author"}, [][]string{
    {"11111111111111111111111111111111111111", "2", "3", "4"},
    {"1", "2", "3", "4"},
    {"1", "2", "3", "4"},
    {"1", "2", "3", "4"},
})


```

```go
+----------------------+------+------+--------+
| Module               | Type | Path | Author |
+----------------------+------+------+--------+
| 11111111111111111111 | 2    | 3    | 4      |
| 111111111111111111   |      |      |        |
| 1                    | 2    | 3    | 4      |
| 1                    | 2    | 3    | 4      |
| 1                    | 2    | 3    | 4      |
+----------------------+------+------+--------+
```

## MIT License

Copyright (c) 2022 ICU-Coders

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
