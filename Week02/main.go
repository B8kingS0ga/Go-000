//这里是模拟项目运行方法
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
