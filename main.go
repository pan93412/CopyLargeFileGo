/*
 * Copy Large File (CLF) 主檔案
 * 版本：v0.5_beta
 *
 * 編譯方式：  go build -o clf.out main.go string.go libs.go
 * (Windows) go build -o clf.exe main.go string.go libs.go
 * 如果是 Linux，則直接輸入 bash build.sh 即可。
 *
 * 語言：zh_TW。若要變更語系，請翻譯 string.go 中字串。
 */

package main
import (
  "fmt"
  "os"
  _ "runtime"
)

// 常數
const ProgramVer = "v0.5_beta"
                   
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
      err := CopyDirectory(src, dsc, ver)
      if err != nil {
        panic(err)
      }
      break
    case !stat.IsDir(): // 若是個檔案
      err := CopyFile(src, dsc, ver)
      if err != nil {
        panic(fmt.Sprintf("%v%s", err, Err_GitHubIT))
      }
      break
    case !rec && stat.IsDir(): // 若是個資料夾，卻沒加上 -r 引數
      panic(Err_FolderNotRecursive + Err_GitHubIT)
      break
    default: // 此處未包括的狀況
      panic(Err_unknownErrorWhenProcess + Err_GitHubIT)
  }
}

// usage 函式
func usage() {
  fmt.Printf(Usage, ProgramVer, os.Args[0])
  os.Exit(1)
}

// 主函式
func main() {
  // 初始化變數
  var recursive bool = false
  var verbose bool = false

  // 檢查語言檔版本
  if StrVer != ProgramVer {
    panic(fmt.Sprintf(Err_LanguageFileVer, ProgramVer, StrVer) + Err_GitHubIT)
  } else if LibVer != ProgramVer {
    panic(fmt.Sprintf(Err_LibFileVer, ProgramVer, LibVer) + Err_GitHubIT)
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
