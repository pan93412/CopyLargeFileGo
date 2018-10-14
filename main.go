package main
import (
  "fmt"
  "os"
  "runtime"
)

// 常數
const programVer = "snapshot_20181015:0053(dev)"
const giturl = "<url>"

// process 為複製檔案的函式。
func process(src string, dsc string, rec bool, ver bool) {
  fmt.Printf(Err_NotSupport, runtime.GOOS, runtime.GOARCH, giturl)
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
      if args == "-r" { recursive = true } // 若使用者指定 -r 則開啟遞迴模式
      if args == "-v" { verbose = true } // 若使用者指定 -v 則開啟詳細模式
    }
    process(os.Args[1], os.Args[2], recursive, verbose)
  }
}
