/**
@author: Jason Pang
@desc: 常用加密算法
@date: 2022/12/29
**/
package various

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rc4"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

////////////////////////////////RC4
func RC4() {
	//加密
	var key []byte = []byte("fd6cde7c2f4913f22297c948dd530c84") //初始化用于加密的KEY，长度1byte~256byte
	rc4obj1, _ := rc4.NewCipher(key)                            //返回 Cipher
	rc4str1 := []byte("RC4待加密数据")                               //需要加密的字符串
	plaintext := make([]byte, len(rc4str1))
	rc4obj1.XORKeyStream(plaintext, rc4str1)   //加密
	stringinf1 := fmt.Sprintf("%x", plaintext) //转换字符串，base-16 编码的字符串，每个字节使用 2 个字符表示
	fmt.Println("RC4加密后:" + stringinf1)

	//解密
	dest2 := make([]byte, len(rc4str1))
	cipher2, _ := rc4.NewCipher(key) // 切记：这里不能重用cipher1，必须重新生成新的
	cipher2.XORKeyStream(dest2, plaintext)
	fmt.Printf("RC4解密后:%s \n\n\n\n", dest2)
}

////////////////////////////////AES
func AES() {
	origData := []byte("AES待加密数据")    // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥，只能128位、192位和256位
	fmt.Println("原文：", string(origData))

	fmt.Println("------------------ CBC模式 --------------------")
	encrypted := AesEncryptCBC(origData, key)
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCBC(encrypted, key)
	fmt.Println("解密结果：", string(decrypted))

	fmt.Println("------------------ ECB模式 --------------------")
	encrypted = AesEncryptECB(origData, key)
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptECB(encrypted, key)
	fmt.Println("解密结果：", string(decrypted))
	fmt.Println("\n\n")
}

// =================== CBC ======================
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式，key[:blockSize]是IV
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

////////////////////////////////RSA
func RSA() {
	RSAEncDec()
	RSASignVerify(crypto.SHA256)
}

func RSAEncDec() {
	origData := []byte("RSA待加密数据") // 待加密的数据，不能超过指定长度
	fmt.Println("原文：", string(origData))
	//生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	ShowRSAKeys(privateKey)
	//生成公钥
	publicKey := privateKey.PublicKey
	//根据公钥加密
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		origData, //需要加密的字符串
		nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文(bytes): ", encryptedBytes)
	fmt.Println("密文(hex)：", hex.EncodeToString(encryptedBytes))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encryptedBytes))
	//根据私钥解密
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	fmt.Println("decrypted message: ", string(decryptedBytes))
	fmt.Println("\n\n")
}

//PKCS1格式的key
func ShowRSAKeys(rsaPrivateKey *rsa.PrivateKey) {
	privateKey := string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rsaPrivateKey),
	}))

	derPkix, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
	if err != nil {
		return
	}

	publicKey := string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}))
	fmt.Printf("公钥：%v\n私钥：%v\n", publicKey, privateKey)
}

//签名和验签
func RSASignVerify(algorithmSign crypto.Hash) {
	origData := []byte("RSA待签名数据") // 待签名的数据，长度无影响
	fmt.Println("原文：", string(origData))
	//生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	ShowRSAKeys(privateKey)
	//生成公钥
	publicKey := privateKey.PublicKey

	//签名
	hash := algorithmSign.New()
	hash.Write(origData)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, algorithmSign, hash.Sum(nil))
	if err != nil {
		panic(err)
	}
	fmt.Println("签名(bytes): ", sign)
	fmt.Println("签名(hex)：", hex.EncodeToString(sign))

	//验签
	err = rsa.VerifyPKCS1v15(&publicKey, algorithmSign, hash.Sum(nil), sign)
	if err == nil {
		fmt.Println("验签成功")
	} else {
		fmt.Println("验签失败")
	}
	fmt.Println("\n\n")
}

////////////////////////////////Sha256
func SHA256() {
	src := "sha256待处理数据"
	fmt.Println("原文：", string(src))
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	fmt.Println("sha256摘要数据：", res) //长度256bit，64字节
	fmt.Println("\n\n")
}

func secretmain() {
	RC4()
	AES()
	RSA()
	SHA256()
}
