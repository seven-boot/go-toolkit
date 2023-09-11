package encrypt

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := []byte("vgYb7C978Spm2e0N")
	str := []byte("hello world")
	encrypt, err := Encrypt(key, str)
	if err != nil {
		t.Error(err)
	}

	//valid := utf8.Valid(encrypt)
	//t.Log(valid)
	t.Log(isGBK(encrypt))

	//decrypt, err := Decrypt([]byte(key), encrypt)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("str is %s\n", string(decrypt))
}

func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			//编码小于等于127,只有一个字节的编码，兼容ASCII吗
			i++
			continue
		} else {
			//大于127的使用双字节编码
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}
