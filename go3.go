package main

import(
    "fmt"
    "os"
    "time"
	"strings"
)
func main() {
s, sep := "", ""
start := time.Now()
for _, arg := range os.Args[1:] {
s += sep + arg
sep = " "
}
fmt.Println(s)
fmt.Printf("%.5fs elapsed\n", time.Since(start).Seconds())
start2 := time.Now()
fmt.Println(strings.Join(os.Args[1:]," "))
fmt.Printf("%.5fs elapsed\n",time.Since(start2).Seconds())

}