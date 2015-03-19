package main

import (
    "crypto/sha1"
    "fmt"
    "os"
    "bufio"
)

func calc(filename string) ([]byte, error){
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    info, err := file.Stat()
    if err != nil {
        return nil, err
    }

    size := info.Size()
    // 如果size > 1G，则用1M的缓冲区
    bufferSize := 4096
    if size > 1024*1024*1024 {
        bufferSize = 1024 * 1024
    }

    reader := bufio.NewReaderSize(file, bufferSize)
    hash := sha1.New()
    _, err = reader.WriteTo(hash)
    if err != nil {
        return nil, err
    }

    return hash.Sum(nil), nil
}

func main() {
    if len(os.Args) <= 1 {
        fmt.Printf("using sha1sum filename\n")
        return
    }

    for _, filename := range(os.Args[1:]) {
        hash, err := calc(filename)
        if err  != nil {
            fmt.Printf("err: %s             %s\n", err.Error(), filename)
            continue
        }

        fmt.Printf("%x          %s\n", hash, filename)
    }

}
