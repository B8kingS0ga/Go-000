
## 作业
我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码


## 答案
结论: 可以考虑, 但是不一定必须
根据课上讲的几点要求
    1. 如果是基础库, 直接返回根原错误不要使用`wrap`, 这点`sql`库满足要求, 使用了 `sentinel error` ErrNoRows
    2. 如果是其他库进行写作, 考虑使用 `errors.Wrap` 或者 `errors.Wrapf` 保存堆栈信息, 这点 `dao` 层满足需求, 属于跟`sql`层协作, 需要根据需求包装``error
    
根据课上讲的几点要求, 是可以考虑使用 `errors.Wrap` 的, 但是结合业务场景, 上层并不需要知道这个错误,因为没有寻找到数据的情况非常常见, 
可以直接通过判断返回数组长度(多条数据) 或者返回数据是否为空(单条数据)来判断, 并不需要抛给上层知道错误发生. 但是如果当理论上一定会获取数据时, 
业务上来讲没有查询到数据就会是一个错误, 这种时候就需要包装error给上层, 让上层判断并执行操作, 并打出堆栈日志.


```go
package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

//********************* dao 层处理 *********************
type Data struct {
	A int
}

type FooDao struct{}

var ErrNoFound = errors.New("啥也没找着")

//大部分情况
//如果可能会获取多条数据, 个人觉得可以不用返回 errorNowRows,
//内部消化就好, 因为空数组在外部看来是可能的结果, 比如翻太多页等等
//service可以通判断数组是否为空解决
//但其他错误又可以被正常捕获
func (fd *FooDao) GetRaws(id, limit, offset int) (data []Data, err error) {
	//dao 操作...
	dbErr := sql.ErrNoRows
	if dbErr != nil {
		//一些降级处理...
		_ = dbErr //ignore this error, 直接返回空数组
	}
	return
}

//如果有单条数据时, 或者理论上一定会获取一条数据的情况, 可以考虑 wrap错误
func (fd *FooDao) GetRow(id int) (data Data, err error) {
	//dao 操作...
	dbErr := sql.ErrNoRows
	if errors.Is(dbErr, sql.ErrNoRows) {
		//wrap error
		return data, errors.Wrapf(ErrNoFound, "GetRow: id[%d] err:", id)
	}
	//一些逻辑处理
	data.A = 1
	return
}

//********************* service 层处理 *********************
type FooService struct{}

func (fs *FooService) Do() error {
	dao := FooDao{}
	data, err := dao.GetRow(1)

	//一些处理...
	_ = data //模拟处理数据
	return err
}

//********************* api 层处理 *********************
func api(w http.ResponseWriter, req *http.Request) {
	service := FooService{}

	err := service.Do()
	//处理ErrorNoFound
	if err != nil {
		fmt.Printf("发生了错误: %+v", err)

		w.WriteHeader(500)
	}

	//fmt.Fprintf(w, "Hello\n")
}

func main() {

	http.HandleFunc("/api", api)
	http.ListenAndServe(":8090", nil)
}


```
    

## 个人理解
在我们现在的业务场景中, 基本在使用课程里讲过的以下模式
```go
package test 
import "errors"
import "logs"
func foo() error {
    return errors.New("check")
}
func main() {
    err := foo()
    if err != nil {
        logs.Error("error")
        return 
    }
    
}
```
这种模式经常产出日志噪音, 而且很不好处理, 很多地方需要寻找错误原因, 很烦, 这次学习到了 用wrap的方式包裹错误并且还可以记录堆栈信息, 可以立马在项目中使用


    