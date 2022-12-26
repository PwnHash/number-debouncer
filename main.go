package main

import(

"strings"
"bufio"
 "fmt"
 "os"

)

func main(){

var filetext string
var isp string
fmt.Print("Enter filename : ")
fmt.Scan(&filetext)
fmt.Print("Enter ISP: ")
fmt.Scan(&isp)



readFile, err := os.Open(filetext)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)


    for fileScanner.Scan() {
 if strings.Contains(fileScanner.Text(),isp){

     num := fileScanner.Text()[0:12]
f, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
if err != nil {
    panic(err)
}

defer f.Close()

if _, err = f.WriteString(num+"\n"); err != nil {
    panic(err)
}
}

}
}