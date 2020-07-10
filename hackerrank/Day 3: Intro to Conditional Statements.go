package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)
    if  N % 2 != 0   { 
        fmt.Println("Weird")
    } else if N%2==0 && (N>1 && N<6){
        fmt.Println("Not Weird")
    }else if N%2==0 && (N>5 && N<21){
        fmt.Println("Weird")
    }else if N%2==0 && N>20{
        fmt.Println("Not Weird")
}
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
