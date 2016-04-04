// getdir project main.go
/*用于phpvod
遍历扫描指定目录下的电影文件，然后添加到phpvod的MYSQL数据库中
使用：main /www/vod
*/
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func getFilelist(path string) {
	var (
		db     *sql.DB
		result sql.Result
		lastid int64
		err    error
	)
	db, err = sql.Open("mysql", "xiongyi:cexoyq1020@tcp(61.183.118.129:3308)/phpvod?charset=utf8")
	defer func() {
		db.Close()
	}()
	checkErr(err)

	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println("path:", path, "\tfile:", f.Name())
		if i := ts(f.Name(), "."); i == 1 {
			cid := lx(path) //电影的类型
			result, err = db.Exec("INSERT pv_video SET cid=?,nid=1,author='xiongyi',authorid=1,subject=?", cid, f.Name())
			checkErr(err)
			lastid, err = result.LastInsertId()
			checkErr(err)
			result, err = db.Exec("INSERT pv_urls SET vid=?,pid=5,url=?", lastid, path)
			checkErr(err)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ts(s string, sp string) int { //判断文件类型是否为电影
	vs := strings.Split(s, sp)
	fmt.Println("split:", vs, "\t", vs[len(vs)-1])
	switch kzm := vs[len(vs)-1]; kzm {
	case "mp4", "MP4", "RMVB", "rmvb", "rm", "RM", "dat", "DAT", "avi", "AVI", "MPG", "mpg":
		fmt.Println("file is movie!")
		return 1
	}
	return 0
}

func lx(s string) int { //通过文件路径，判断影片类型是日语英语等
	var slx = map[int]string{
		35: "日语",
		34: "俄语",
		33: "韩语",
		32: "德语",
		31: "法语",
		30: "国语",
		43: "英语",
	}
	for k, v := range slx {
		i := strings.Count(s, v) //有多少次匹配到的字符
		if i > 0 {
			return k
		}
	}
	return 30
}

func main() {

	flag.Parse()
	root := flag.Arg(0)
	fmt.Println(root)
	getFilelist(root)
	ts("天天是.rmvb", ".")

}
