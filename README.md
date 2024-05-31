# Vigenère cipher
encrypt (with key)
```
vigenere-encrypt <encipherment key> <plaintext filename> > ciphertext.txt 
```
decrypt (with key)
```
vigenere-decrypt <decipherment key> <ciphertext filename> > recovered plaintext.txt
```
get keylength
```
vigenere-keylength ciphertext.txt
```
decrypt (without key)
```
vigenere-cryptanalyze <ciphertext filename> <key length>
```
