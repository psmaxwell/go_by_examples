package main

import (
     "fmt"
     "time"
)

func main() {
    p := fmt.Println

    t := time.Now()

    p(t.Format(time.RFC3339))

    t1, e := time.Parse(
	time.RFC3339,"2023-02-22T14:09:41+00:00")
    p(t1)

    p(t.Format("2:10PM"))
    p(t.Format("Mon Feb _22 14:04:05 2023"))
    p(t.Format("2023-02-22T14:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)

    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",t.Year(), t.Month(), t.Day(),t.Hour(), t.Minute(), t.Second())

    ansic := "Wed Feb _22 14:04:05 2023"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}
