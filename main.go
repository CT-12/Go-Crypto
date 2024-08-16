package main

import (
	"flag"
	"fmt"
	"go-encrypted/utils/crypto"
	"go-encrypted/utils/file_manager"
	"os"
)

var(
	read_path, key string
	output_path string = "./output.txt"
)

func main(){	
	decrypt_flag := flag.Bool("d", false, "For decryption")
	read_path_flag := flag.String("f", "", "Read file path")
	key_flag := flag.String("k", "", "The key for the encryption and decryption")
	output_path_flag := flag.String("o", "output.txt", "Output file path for encrypted file")
	flag.Parse()

	if *read_path_flag == "" || *key_flag == ""{
		fmt.Println("Missing read file path or secret key")
		flag.Usage()
		os.Exit(1)
	}
	read_path = *read_path_flag
	key = *key_flag
	output_path = *output_path_flag

	text, err := file_manager.ReadFile(read_path)
	checkErr(err)

	var result string
	if *decrypt_flag{
		// 解密		
		result, err = crypto.DecryptAES(text, key)
	}else{
		// 加密
		result, err = crypto.EncryptAES(text, key)
	}
	checkErr(err)

	err = file_manager.WriteFile(output_path, result)
	checkErr(err)
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}