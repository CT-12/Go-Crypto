# Go-Crypto

This project is for locally encrypt and decrypt.

## Usage
```shell
./go-encrypted -f <read_file_path> -k <secret key> [-o <output_file_path>][-d]
```
* -f : The file path you want to encrypt.
* -o : The output file path you want to place.
* -k : The secret key used for encrypting and decrypting file.
* -d : Decryption mode. Default is encryption mode.