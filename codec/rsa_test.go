package codec

import (
	"encoding/base64"
	"fmt"
	"testing"
)

// func TestRSA(T *testing.T) {
// 	data, _ := hex.DecodeString("4e53554356318000ad402145b5f6c34ff6d4d112bd2910ccbc68673a38456acd22f7cf6a523efd70e1f40579ed14cd4175d99d4af727c89926e84dcb020a1d96119cb2848033f29b95a150d8fca178bb7b2c6b307e51a1fc8a6162d5730ff7b5b1723f34b0253eae9c521f9b68eb70f45d3d1e257ef636f3bc3cb2c63c371fb75e561e77bcf961ed")
// 	data = data[8:]
// 	data2, err := RsaDecrypt(data[:128], )
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	println(string(data2))
// 	enc, _ := RsaEncrypt([]byte("q342342352352352354"))
// 	fmt.Println(len(enc))
// }

func TestEncrypt(T *testing.T) {
	a := []byte("但是无论他怎样努力，终究事与愿违。他的父亲被弟弟钉死在茅坑，他的姐姐殚精竭虑想杀死弟弟，弟弟远走他乡，三个孩子都死于非命。他本来就不该在权力的中心，在政治，家族，爱情和梦想中，所有人都逼迫他作出选择，但是他从来都不是那个游戏者，时隔多年，他还是惦念着15岁白衣白甲的骑士少年。")
	b, _ := RsaEncrypt(a)
	fmt.Println(base64.StdEncoding.EncodeToString(b))
}
