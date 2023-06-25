package jwt

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// this will generate JWT token
func GenerateToken(header string, payload map[string]string, secret string) (string, error) {
	// create a new hash of type sha256. We pass the secret key to it
	// sha512 is a symmetric cryptographic algorithm
	h := hmac.New(sha512.New, []byte(secret))

	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	payloadstr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error in Genertating the Token")
		return string(payloadstr), err
	}

	payload64 := base64.StdEncoding.EncodeToString(payloadstr)

	// Now add the encoded string.
	message := header64 + "." + payload64

	unsignedStr := header + string(payloadstr)

	// we write this to the SHA512 to hash it. We can use this to generate the signature now
	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	//Finally we have the token
	tokenStr := message + "." + signature
	return tokenStr, nil
}

// Let's validate the Token
func ValidateToken(token string, secret string) (bool, error) {
	// we have sepearted our JWT in three parts with "."

	splitToken := strings.Split(token, ".")
	// if length is not 3, we know that the token is corrupt
	if len(splitToken) != 3 {
		return false, nil
	}

	// decode the header and payload back to strings
	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}
	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, err
	}

	//again create the signature
	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Println(signature)

	// if both the signature dont match, this means token is wrong
	if signature != splitToken[2] {
		return false, nil
	}
	// This means the token matches
	return true, nil
}
