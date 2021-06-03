package main

import "fmt"

func main() {
    name := "太郎"
    fmt.Printf("name :%v\n", name)

    namePoint := &name

    // namePointは、&nameが格納されているだけなので、stringへのポインタである*string型の値が格納されている。
    fmt.Printf("namePoint :%v\n", namePoint)

    // namePointが指している変数は、"*namePoint"という感じで、"*"をつけて表す。
    fmt.Printf("namePoint :%v\n", *namePoint)

    *namePoint = "二郎"

    // *namePointに値を代入することもできる。
    fmt.Printf("*namePointに二郎を代入後の*namePoint :%v\n", *namePoint)

    // 再代入したところで、namePointに格納されている*string型の値(アドレス)自体は、変わらない
    fmt.Printf("*namePointに二郎を代入後のnamePoint :%v\n", namePoint)

    // stringへのポインタである*string型の値(nameに格納されている値)を書き換えたので、nameの値も変更される。
    fmt.Printf("*namePointに二郎を代入後のname :%v\n", name)
}
