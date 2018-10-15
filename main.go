package main
import (
  "fmt"
  "os"
  _ "runtime"
)

// 常數
const programVer = "snapshot_20181015:1826(dev)"
const giturl = "<url>"
                   
// process 為複製檔案的函式。
func process(src string, dsc string, rec bool, ver bool) {
  // 判斷 src 和 dsc
  if _, err := os.Stat(dsc); err == nil { // 若 dsc 存在
    panic(Err_dscExists)
  } else if _, err := os.Stat(src); os.IsNotExist(err) { // 若 src 不存在
    panic(Err_srcNotExists)
  }
  
  // 判斷引數
  stat, _ := os.Stat(src)
  switch {
    case rec && stat.IsDir(): // 若有加上 -r 引數，且是個資料夾
      fmt.Println("case 1 detected.")
      break
    case !stat.IsDir(): // 若是個檔案
      fmt.Println("case 2 detected.")
      break
    case !rec && stat.IsDir(): // 若是個資料夾，卻沒加上 -r 引數
      panic(Err_FolderNotRecursive)
      break
    default: // 此處未包括的狀況
      panic(Err_unknownErrorWhenProcess)
  }
}

// usage 函式
func usage() {
  fmt.Printf(Usage, programVer, os.Args[0])
  os.Exit(1)
}

// 主函式
func main() {
  // 初始化變數
  var recursive bool = false
  var verbose bool = false

  // 檢查語言檔版本
  if Version != programVer {
    panic(fmt.Sprintf(Err_LanguageFileVer, programVer, Version))
  }
  
  if len(os.Args) < 3 || len(os.Args) > 5 {
    usage() // 顯示用法文字後退出程式
  } else {
    for _, args := range os.Args {
      switch args {
        case "-r": recursive = true // 若使用者指定 -r 則開啟遞迴模式
        case "-v": verbose = true // 若使用者指定 -v 則開啟詳細模式
      }
    }
    process(os.Args[1], os.Args[2], recursive, verbose)
  }
}
