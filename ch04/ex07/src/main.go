package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) []byte {
	if len(b) == 0 {
		return b
	}

	result := make([]byte, len(b))
	i := len(b)
	for {
		if i <= 0 {
			break
		}
		_, size := utf8.DecodeLastRune(b)
		for _, j := range b[len(b)-size : len(b)] {
			result = append(result, j)
		}
		b = b[0 : len(b)-size]
		i -= size
	}

	return result
}

func main() {
	str := `
。い多が者用利に者理管ムテスシやーマラグロプに特、めたるいてし化特にのるす集編をルイァフ定設ムテスシ
やドーコムラグロプ、てせ併と能機の他のmiV。るす在存が家好愛miVのく多数
、めたるきでがとこる得を度速集編トスキテのどほいならなに較比はとどな帳モメばえましてれ慣旦一、らがなしかし
。るなと要必がれ慣でまるなにうよるきでが業作集編トスキテのり通一
、めたるな異でるまが法方作操はとタィデエの他のどな）どな帳モメ（タィデエ系swodniW
	`
	fmt.Println("before")
	fmt.Println(str)

	fmt.Println("after")
	res := reverse([]byte(str))
	fmt.Println(string(res))
}
