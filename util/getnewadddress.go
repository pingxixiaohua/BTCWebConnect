package util

import (
	"BcAddressCode/base58"
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func Getnewaddress() string {
	//1.生成私钥
	curve:=elliptic.P256()
	pri,x,y,err:=elliptic.GenerateKey(curve,rand.Reader)
	if err!=nil {
		fmt.Println(err.Error())
		return "错误"
	}
	fmt.Println(pri)


	//x.Bytes()大整型转字节切片
	//拼接x，y生成公钥
	pubKey:=append(x.Bytes(),y.Bytes()...)
	pubKey=append([]byte{4},pubKey...)
	publickey:=elliptic.Marshal(curve,x,y)
	fmt.Println("非压缩格式公钥：",publickey)

	pub256Hash:=SHA256hash(pubKey)

	pubRpmd160:=Rpmd160hash(pub256Hash)

	//2.添加版本号
	versionPubRpmd16:=append([]byte{0x00},pubRpmd160...)

	//3.计算校验位
	hash1:=SHA256hash(versionPubRpmd16)
	hash2:=SHA256hash(hash1)
	check1:=hash2[:4]//字节切片的截取

	//4.拼接
	addBytes:=append(versionPubRpmd16,check1...)

	//5.对地址拼接的结果进行base58编码
	Base58Check:=base58.Encode(addBytes)
	fmt.Println("生成新的比特币地址："+Base58Check)
	//校验
	if CheckAddress(Base58Check)==true{
		return Base58Check
	}else {
		return "错误"
	}

	return "错误"

}

//用校验位进行校验检查
func CheckAddress(add string)bool  {

	//对地址base58解码
	deAddBytes:=base58.Decode(add)
	// 取其后四位作为校验位
	check:=deAddBytes[len(deAddBytes)-4:]

	//双Hash计算，地址反编码后去除后四字节
	deAddBytesExpelCheck:=deAddBytes[:len(deAddBytes)-4]
	deHash1:= SHA256hash(deAddBytesExpelCheck)
	deHash2:= SHA256hash(deHash1)

	//获取前四个字节
	deCheck:=deHash2[:4]

	//判断
	if string(check)==string(deCheck) {
		fmt.Println("地址有效^-^")
		return true
	}else {
		fmt.Println("地址无效v_v")
		return false
	}

	return bytes.Compare(check,deCheck)==0
}