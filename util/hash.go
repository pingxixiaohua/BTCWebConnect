package util

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

func SHA256hash(data []byte)[]byte  {
	hash:=sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func Rpmd160hash(data []byte)[]byte  {
	rpmd160:=ripemd160.New()
	rpmd160.Write(data)
	return rpmd160.Sum(nil)
}
