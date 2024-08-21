# Go-Crypto

This project is for locally encrypt and decrypt.

## Usage
```shell
./go-crypto -f <read_file_path> -k <secret key> [-o <output_file_path>][-d]
```
Required : 
* -f : The file path you want to encrypt. 
* -k : The secret key used for encrypting and decrypting file.

Optional : 
* -o : The output file path you want to place.
* -d : Decryption mode. Default is encryption mode.