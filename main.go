package main
import (
  "fmt"
  "os"
  "runtime"
  "errors"
)

// 常數
const programVer = "snapshot_20181015:1735(dev)"
const giturl = "<url>"

// process 為複製檔案的函式。
func process(src string, dsc string, rec bool, ver bool) {
  // 判斷 src 和 dsc
  if _, err := os.Stat(dsc); os.IsExists(err) {
    panic err_dscExists
  } else if _, err := os.Stat(src); os.IsNotExists(err) {
    panic err_srcNotExists
  }
  
  // 判斷引數
  if rec && os.Stat(src).IsDir {
    rec_folder(src, dsc, ver)
  } else if !os.Stat(src).IsDir {
    cp_file(src, dsc, ver)
  } else if !rec && os.Stat(src).IsDir {
    panic err_FolderNotRecursive
  } else {
    panic err_unknownErrorWhenProcess
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
      if args == "-r" { recursive = true } // 若使用者指定 -r 則開啟遞迴模式
      if args == "-v" { verbose = true } // 若使用者指定 -v 則開啟詳細模式
    }
    process(os.Args[1], os.Args[2], recursive, verbose)
  }
}
