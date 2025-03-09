# Golang

## Lịch sử ra đời
_ Ý tưởng của ngôn ngữ Go bắt đầu vào năm 2007 tại Google. 3 kỹ sư tại Google: Robert Griesemer, Rob Pike, và Ken Thompson cố gắng tìm ra một ngôn ngữ tại thời điểm đó thỏa mãn 3 điều kiện sau:

- Biên dịch hiệu quả
- Thực thi hiệu quả
- Dễ lập trình

Tuy nhiên, tại thời điểm đó, không tồn tại ngôn ngữ nào thỏa mãn cả 3 điều kiện trên. Do đó, họ muốn tạo ra một ngôn ngữ lập trình mới giúp thỏa mãn các điều kiện trên, và Golang ra đời.

## So sánh Golang với các ngôn ngữ khác
### Python
| |Go | Python | So sánh|
|---|-----|----|--------|
|Kiểu dịch | **Biên dịch** | Thông dịch | Tốc độ dịch của Go nhanh hơn Python|
|Kiểu dữ liệu | **Tĩnh** | Động | Kiểu dữ liệu của Go cố định và an toàn hơn Python|
|Concurrency | **Goroutine** | Chậm do chỉ dùng 1 thread | Điểm mạnh nhất của Go và điểm yếu nhất của Python|
|Cú pháp | Đơn giản | **Ngắn gọn** | Cú pháp của Python ngắn gọn hơn nhiều so với Go và các ngôn ngữ khác|

### Java
| |Go | Java | So sánh|
|---|-----|----|--------|
|Kiểu dịch | **Biên dịch** | Nửa biên dịch (thông qua JVM) | Tốc độ dịch của Go nhanh hơn Java|
|Cú pháp | **Đơn giản** | Phức tạp | Do Java thuần hướng đối tượng nên cú pháp phức tạp hơn nhiều|
|Concurrency | **Goroutine** | Thread | Goroutine nhẹ hơn nhiều so với Thread|
|Độ phổ biến | Đang phát triển | **Có cộng đồng lớn** | Java có cộng đồng lớn hơn nên có nhiều thư viện hỗ trợ hơn từ cộng đồng |

### C++
| |Go | C++ | So sánh|
|---|-----|----|--------|
|Kiểu dịch | Biên dịch | **Biên dịch** | Tốc độ dịch của C++ nhanh hơn Go|
|Quản lý bộ nhớ | **Tự động thông qua Garbage Collector** | Thủ công | C++ đòi hỏi người dùng tự quản lý bộ nhớ mình khai báo|

## Go được sử dụng cho những ứng dụng nào?
### Dịch vụ mạng và đám mây
_ Là giải pháp cho sự đánh đổi giữa thời gian phát triển và hiệu năng của server.

_ Giải quyết nhiều thách thức mà các nhà phát triển phải đối mặt với đám mây hiện đại, cung cấp các API chuẩn và tích hợp sẵn tính năng concurrency để tận dụng bộ xử lý multicore.

_ Ví dụ điển hình là Docker và Kubernetes.

### Phát triển ứng dụng Web
_ Triển khai trên nhiều nền tảng với tốc độ cao.

_ Tận dụng hiệu suất sẵn có của Go để mở rộng quy mô dễ dàng, đặc biệt không cần thiết phải có một framework nào cho Go.

### Command Line Interface
_ Ứng dụng CLI ưa chuộng ngôn ngữ Go do tính di động, hiệu năng và khả năng dễ dàng sáng tạo.

_ Tận dụng thời gian biên dịch nhanh để xây dựng các chương trình khởi động nhanh và chạy trên bất kỳ hệ thống nào.

## Function trong Go
Function trong Go thông thường được khai báo với cú pháp sau:
```go
func f(input INPUT_TYPE) OUTPUT_TYPE{
   return ...
}
```

Trong đó:

+ f là tên hàm
+ input là tham số với kiểu dữ liệu INPUT_TYPE
+ OUTPUT_TYPE là kiểu dữ liệu trả về

### Multiple return output
Để function có thể trả về nhiều biến với nhiều kiểu dữ liệu khác nhau, ta dùng cú pháp sau:
```go
func f() (OUTPUT_TYPE_1, OUTPUT_TYPE_2) {
    return output1, output2
}
```

### Variadic function
Variadic function là một hàm cho phép truyền vào 0 hoặc nhiều tham số cùng kiểu dữ liệu. Ví dụ:
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// sum(1, 2, 3) = 10
// sum(10, 20, 30, 40) = 100
// sum() = 0
```

trong đó, **number ...int** cho phép truyền nhiều tham số có kiểu dữ liệu int vào hàm, và biến numbers có thể coi như kiểu dữ liệu []int.

### Receiver
Để giới hạn một hàm chỉ được truy cập bởi một object nào đó, ta truyền receiver vào hàm đó.
```go
type Rectangle struct {
	width float64
	height float64
}

func (rec Rectangle) getArea() {
	return rec.width * rec.height
}
```

### Error handling
Trong Go, error được coi như một kiểu dữ liệu, và được trả về như một giá trị của hàm. Nếu không có lỡi thì error có giá trị là nil.
```go
func division(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.new("Division by zero")
    }
    
    return a / b, nil
}
```

### Exception handling
Trong Go không có cơ chế bắt ngoại lệ như try/catch như các ngôn ngữ khác. Thay vào đó, Go sử dụng **panic** và  **recover**.
```go
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recover from panic:", r)
        }
    }

    panic("Something went wrong!")
}
```

## For-loop
### Vòng lặp cơ bản
Ví dụ, để duyệt tất cả các số từ 1 đến 10: 
```go
for i := 1; i <= 10; i++ {
    //
}
```

_ Lệnh i := 1 là lệnh khởi tạo được chạy 1 lần duy nhất khi bắt đầu vòng lâu.

_ Lệnh i <= 10 là lệnh điều kiện để kết thúc vòng lặp.

_ Lệnh i++ là lệnh được chạy sau mỗi lần lặp.

### Vòng lặp while
Nếu chỉ giữ lại lệnh điều kiện, vòng lặp sẽ trở thành vòng lặp while:
```go
i := 1
for i <= 10 {
    i++
}
```

### Vòng lặp vô hạn
Nếu không truyền gì vào vòng lặp, vòng lặp trở thành vòng lặp vô hạn:
```go
for {
    fmt.Println("Hello World!")
}
```
### Duyệt slice
Giả sử có slice là 1 mảng gồm các số nguyên:
```go
var arr = []int{1, 2, 3, 4, 5}
```

Để duyệt slice, ta dùng cú pháp sau:
```go
for i, x := range arr {
    fmt.Println(i, x)
}
```

Khi đó i là chỉ số, còn x là giá trị arr[i] của mảng.

Lưu ý, nếu duyệt như sau:
```go
for x := range arr {
    fmt.Println(x)
}
```
thì x sẽ là chỉ số mảng, chứ không phải giá trị mảng như kỳ vọng.

Để duyệt giá trị mảng mà không cần quan tâm chỉ số, ta dùng cú pháp:
```go
for _, x := range arr {
    fmt.Println(x)
}
```
Trong đó _ đại diện cho 1 giá trị khai báo mà ta không muốn sử dụng.

### Kết thúc vòng lặp
_ Lệnh continue dùng để kết thúc lần lặp hiện tại để sang lần lặp mới.

_ Lệnh break dùng để kết thúc toàn bộ vòng lặp.

```go
for i := 1; i <= 10; i++ {
    if i % 2 != 0 {
        continue
    }
    
    if i > 5 {
        break
    }
}
```

## Package, Import
Package trong Go là một cách để tổ chức và tái sử dụng lại code. Package giúp phân chia source code thành các phần nhỏ hơn. Mỗi file đều thuộc 1 package chính là tên thư mục chứa file đó.

Danh sách các public packages của Golang ở [link](https://pkg.go.dev/std).

Để sử dụng một package, ta dùng lệnh import. Ví dụ:
```go
import "fmt"
```
Ví dụ cho import nhiều packages:
```go
import (
    "fmt"
    "strings"
    "structs"
)
```

Sau khi import, ta có thể sử dụng mọi function mà package đấy có. Ví dụ như muốn in ra màn hình Hello World:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

Ta còn có thể import 1 package nội bộ do chính mình tạo ra. 

## Go module
Module là một tập hợp các packages được lưu trong cùng 1 folder với file go.mod là gốc định nghĩa module path cho module đó. Ví dụ:

```md
example
├── hello
│   └── hello.go
├── go.mod
└── main.go
```

_ Đầu tiên, khởi tạo module cho folder hiện tại.
```bash
go mod example
```

Sẽ có file go.mod được tạo ra với thông tin:
```txt
module example

go 1.24.1
```

_ Để file main.go sử dụng được các hàm của package hello, ta có thể import "example/hello":
```go
package main

import hl "example/hello"

func main() {
	hl.Hello()
}
```


## Naming convention
### Tên package
    + Sử dụng tên gồm chữ in thường, ngắn gọn và miêu tả rõ chức năng của package đó.
    + Tránh sử dụng dấu gạch dưới hay số nhiều cho tên.
VD:
```go
package auth
package authentication  // Quá dài
package auth_utils      // Xấu do vừa có dấu gạch dưới, vừa là số nhiều
```

### Tên biến và hàm
    + Đặt tên theo camelCase.
    + Bắt đầu bằng chữ in thường với những biến và hàm private.
    + Bắt đầu bằng chữ in hoa với những biến và hàm public.

VD:
```go
func getUserID() int { ... }  // hàm private
func GetUserID() int { ... }  // hàm public
```

### Hằng số
    + Sử dụng camelCase cho hằng số cục bộ.
    + Viết hoa toàn bộ kết hợp dấu gạch dưới nếu là hằng số có liên quan đến một nhóm hằng số khác.

VD:
```go
const defaultTimeout = 30
const MAX_CONNECTIONS = 100
```

### Tên file
Viết thường kết hợp dấu gạch dưới.

VD:
```go
user_service.go
http_handler.go
```

### Error
Bắt đầu với err hoặc Err.

VD:
```go
var ErrNotFound = errors.New("not found")
```

## Data Structures
### Array
Mảng là một tập hợp các phần tử có cùng kiểu dữ liệu. Mảng phải khai báo trước số lượng phần tử tối đa trong mảng có thể được sử dụng.
#### Khai báo mảng
```go
var studentAge [10] int	// Mảng 1 chiều
var matrix [10][10]int		// Mảng 2 chiều
```

#### Khai báo mảng chỉ định các phần tử cho trước
```go
var studentsAge = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

#### Truy cập vào phần tử trong mảng
```go
for i := 0; i < 10; i++ {
	fmt.Println(studentAge[i])
}

for _, x := range studentAge {
	fmt.Println(x)
}
```

#### Sửa giá trị trong mảng
```go
studentsAge[0] = 5
studentsAge[4] = 15
studentsAge[7] = 10
```

#### Lấy độ dài của mảng
```go
fmt.Println(len(studentAge))
```

### Slice
Slice giống hệt mảng, điểm khác biệt duy nhất là slice có thể thay đổi độ dài, trong khi độ dài của mảng là cố định.
#### Khai báo slice
```go
var sliceOfIntegers []int	//Mảng 1 chiều
var matrix [][]int				//Mảng 2 chiều
```

#### Khai báo slice với hàm make
```go
sliceOfIntegers := make([]int, 5)  // [0 0 0 0 0]
sliceOfBooleans := make([]bool, 3) // [false false false]
```

#### Thêm phần tử vào slice
```go
sliceOfIntegers := []int{1, 2, 3}

sliceOfIntegers = append(sliceOfIntegers, 4)	// Thêm 1 phần tử
sliceOfIntegers = append(sliceOfIntegers, 5, 6, 7)	//Thêm nhiều phần tử (Do đây là hàm variadic input)
```

### Struct
Struct là một kiểu dữ liệu được tạo ra bởi người dùng từ một tập các kiểu dữ liệu đã được định nghĩa khác.
#### Khởi tạo struct
```go
type Rectangle struct {
  length  float64
  breadth float64
}
```

#### Khởi tạo struct instance
```go
var myRectangle Rectangle
myRectangle := Rectangle{}
```

#### Khởi tạo struct instance với dữ liệu cho trước
```go
myRectangle := Rectangle{10, 5}
myRectangle := Rectangle{length: 10, breadth: 5}
myRectangle := Rectangle{breadth: 10}
```

#### Truy cập và thay đổi các thuộc tính của struct
```go
myRectangle := Rectangle{length: 10, breadth: 5}

fmt.Println(myRectangle.length)  // 10
fmt.Println(myRectangle.breadth) // 5

myRectangle.length = 20
myRectangle.breadth = 8
fmt.Println(myRectangle) // {20 8}
```

### Map
### Interface

## Multithreading
## Common packages