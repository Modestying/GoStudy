package ChineseChar

import (
	"fmt"
	"strconv"
	"testing"
)


func TestOther(t *testing.T) {
	text := "a4浮"
	//golang中一个汉字占3个byte
	t.Log(len([]byte(text)))

	textRune := []rune(text)
	textLen := len(textRune)
	t.Log("len:" + strconv.FormatInt(int64(textLen), 10))
	for i := 0; i <= textLen-1; i++ {
		t.Log(fmt.Sprintf("word:%s", string(textRune[i:i+1])))
	}
}
