package main

import "time"
import "fmt"

var t time.Time = time.Now()

func main() {
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t, t.Location(), t)
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t.UTC(), t.UTC().Location(), t)
	t_by_unix := time.Unix(1487780010, 0)
	fmt.Println("'time.Unix': ", t_by_unix)
}
