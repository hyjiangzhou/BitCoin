package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {

	//1. 创建私钥
	//2. 创建公钥
	//3. 使用私钥签名
	//4. 校验

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKeyRaw := privateKey.PublicKey

	data := "helloworld"
	dataHash := sha256.Sum256([]byte(data))

	// func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error) {
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, dataHash[:])

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("privateKey : %x\n", privateKey)
	fmt.Printf("pubKey : %x\n", pubKeyRaw)

	fmt.Printf("r : %x, len : %d\n", r.Bytes(), len(r.Bytes()))
	fmt.Printf("s : %x, len : %d\n", s.Bytes(), len(s.Bytes()))

	//将r，s拼接成字节流
	signature := append(r.Bytes(), s.Bytes()...)

	//将公钥拼接成字节流传递，在对端拆分
	//type PublicKey struct {
	//	elliptic.Curve
	//	X, Y *big.Int
	//}

	//真正传递的形式
	pubKey := append(pubKeyRaw.X.Bytes(), pubKeyRaw.Y.Bytes()...)


	//....

	//根据signature 切出来r1, s1, 一分为二
	r1 := big.Int{}
	s1 := big.Int{}

	r1Data := signature[:len(signature)/2]
	s1Data := signature[len(signature)/2:]

	r1.SetBytes(r1Data)
	s1.SetBytes(s1Data)

	//切pubkey字节流
	x1 := big.Int{}
	y1 := big.Int{}

	x1Data := pubKey[:len(pubKey)/2]
	y1Data := pubKey[len(pubKey)/2:]

	x1.SetBytes(x1Data)
	y1.SetBytes(y1Data)

	curve := elliptic.P256()
	pubKeyOrigin := ecdsa.PublicKey{curve, &x1, &y1}



	// func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool {
	//res := ecdsa.Verify(&pubKey, dataHash[:], r, s)
	res := ecdsa.Verify(&pubKeyOrigin, dataHash[:], &r1, &s1)

	fmt.Printf("res : %v\n", res)
}
