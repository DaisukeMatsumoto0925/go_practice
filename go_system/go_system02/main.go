package main

import (
	"bufio"
	"os"
)

// func main() {
//     file, err := os.Create("test.txt")
//     if err != nil {
//         panic(err)
//     }
//     file.Write([]byte("os.File example\n"))
//     file.Close()
// }

// func main() {
//     os.Stdout.Write([]byte("os.Stdout example\n"))
// }

// func main() {
//     var buffer bytes.Buffer
//     buffer.Write([]byte("bytes.Buffer example\n"))
//     fmt.Println(buffer.String())
// }

// func main() {
//     conn, err := net.Dial("tcp", "ascii.jp:80")
//     if err != nil {
//     panic(err)
//     }
//     conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
//     io.Copy(os.Stdout, conn)
// }

func main() {
    buffer := bufio.NewWriter(os.Stdout)
    buffer.WriteString("bufio.Writer ")
    buffer.Flush()
    buffer.WriteString("example\n")
    buffer.Flush()
}
