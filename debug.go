package logprint //包名最好小写字母，并且最好和目录名一致
import (
	"fmt"
	"time"
)

//大写表示可以跨包导入，小写表示不能跨包导入

func Debug(msg interface{}) {
	t := time.Now()
	fmt.Printf("[debug]%s:%s\n",t.Format("2006-01-02 15:04:05.000"),msg)
}