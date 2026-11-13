package ch01

import (
	"bytes"
	"errors"
	"fmt"
)

var ErrInvalidInput = errors.New("the input is not valid padded bytes")

func padding(src []byte, blockSize int) []byte {
	srcLen := len(src)
	padLen := blockSize - srcLen%blockSize
	paddedBytes := bytes.Repeat([]byte{byte(padLen)}, padLen)
	result := append(src, paddedBytes...)
	return result
}

func unpadding(padded []byte, blockSize int) ([]byte, error) {
	size := len(padded)
	if size%blockSize != 0 || size < blockSize {
		return nil, ErrInvalidInput
	}
	padLen := int(padded[len(padded)-1])
	result := padded[:size-padLen]
	return result, nil
}

func PKCS7Run() {
	src1 := []byte{1, 2, 3, 4}
	result1 := padding(src1, 8)
	unpadded1, err := unpadding(result1, 8)
	fmt.Printf("src1: %v\n", src1)
	fmt.Printf("result1: %v\n", result1)
	if err != nil {
		fmt.Printf("failed to unpad data: %s\n", err)
	}
	fmt.Printf("result1: %v\n", unpadded1)

	src2 := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	result2 := padding(src2, 8)
	unpadded2, err := unpadding(result2, 8)
	fmt.Printf("src2: %v\n", src2)
	fmt.Printf("result2: %v\n", result2)
	if err != nil {
		fmt.Printf("failed to unpad data: %s\n", err)
	}
	fmt.Printf("result2: %v\n", unpadded2)
}
