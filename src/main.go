package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall"
)

func main() {

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	password := strings.TrimSpace(string(bytePassword))

	h := sha1.New()
	io.WriteString(h, password)

	hash := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	suffix := hash[5:len(hash)]
	prefix := hash[0:5]

	resp, err := http.Get(fmt.Sprintf("https://api.pwnedpasswords.com/range/%v", prefix))
	if err != nil {
		fmt.Printf("\nOn No! An Error Occurred! \n")
		fmt.Printf("%v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bs := string(body)
	s := strings.SplitN(bs, "\r\n", 600)

	m := make(map[string]string)
	for _, pair := range s {
		z := strings.Split(pair, ":")
		m[z[0]] = z[1]
	}
	value := m[suffix]
	if value == "" {
		fmt.Printf("\nPassword not found.")
	} else {
		fmt.Printf("\nFound %v times!", value)
	}
}
