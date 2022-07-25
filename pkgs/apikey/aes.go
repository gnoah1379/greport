package apikey

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"io"
	"strings"
	"sync"
)

var (
	secretKey []byte
	once      sync.Once
)

func getKey() []byte {
	once.Do(func() {
		secretKey = []byte(viper.GetString("secretKey"))
	})
	return secretKey
}

func init() {
	viper.SetDefault("secretKey", "asuperstrong32bitpasswordgohere!")
}

func GetCredential(apikey string) (*credentials.Credentials, error) {
	userPassword, err := decrypt(apikey)
	if err != nil {
		return nil, errors.New("invalid api key")
	}
	userPasswordSplit := strings.SplitN(userPassword, ":", 2)
	if len(userPasswordSplit) != 2 {
		return nil, errors.New("invalid api key")
	}
	return credentials.NewStaticV4(userPasswordSplit[0], userPasswordSplit[1], ""), nil
}

func GetApiKey(user, password string) (string, error) {
	var userPassword = user + ":" + password
	return encrypt(userPassword)
}

func encrypt(message string) (encoded string, err error) {
	//Create byte array from the input string
	plainText := []byte(message)

	//Create a new AES cipher using the key
	block, err := aes.NewCipher(getKey())

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	//iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	//Encrypt the data:
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//Return string encoded in base64
	return base64.RawStdEncoding.EncodeToString(cipherText), err
}

func decrypt(secure string) (decoded string, err error) {
	//Remove base64 encoding:
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)

	//IF DecodeString failed, exit:
	if err != nil {
		return
	}

	//Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(getKey())

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//IF the length of the cipherText is less than 16 Bytes:
	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	//Decrypt the message
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}
