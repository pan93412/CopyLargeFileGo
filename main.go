/*
 * Copy Large File (CLF) 主檔案
 * 版本：v0.7.2-beta
 *
 * 編譯方式：  go build -o clf.exe *.go
 * (Windows) go build -o clf.exe *.go
 * 如果是 Linux，則直接輸入 bash build.sh 即可。
 * 若要在 Linux 環境編譯全平台，請輸入 bash build.sh --all
 *
 * 語言：zh_TW。若要翻譯字串至您的語言，
 * 請翻譯 string.go 中字串。
 */

package main

import (
  "fmt" // 輸出文字
  "os"  // 檢查檔案
)

// 程式版本
const ProgramVer = "v0.7.2-beta"

// 貢獻者
const Contributors = `pan93412 <http://www.github.com/pan93412> (軟體主作者), 2018.` + "\n"

// process 為複製檔案的函式。
func process(src string, dsc string, rec bool, ver bool) {
  // 判斷 src 和 dsc
  if _, err := os.Stat(dsc); err == nil { // 若 dsc 存在
    ErrorHandler(Err_dscExists)
  } else if _, err := os.Stat(src); os.IsNotExist(err) { // 若 src 不存在
    ErrorHandler(Err_srcNotExists)
  }

  // 判斷引數
  stat, _ := os.Stat(src)
  switch {
  case rec && stat.IsDir(): // 若有加上 -r 引數，且是個資料夾
    err := CopyDirectory(src, dsc, ver)
    if err != nil {
      PanicHandler(err)
    }
    break
  case !stat.IsDir(): // 若是個檔案
    err := CopyFile(src, dsc, ver)
    if err != nil {
      PanicHandler(err)
    }
    break
  case !rec && stat.IsDir(): // 若是個資料夾，卻沒加上 -r 引數
    ErrorHandler(Err_FolderNotRecursive)
    break
  default: // 此處未包括的狀況
    ErrorHandler(Err_unknownErrorWhenProcess)
  }
}

// usage 函式
func usage() {
  fmt.Printf(Usage, ProgramVer, os.Args[0], Contributors)
  os.Exit(1)
}

// 主函式
func main() {
  // 初始化變數
  var recursive bool = false
  var verbose bool = false

  // 檢查語言檔版本
  if StrVer != ProgramVer {
    ErrorHandler(fmt.Sprintf(Err_LanguageFileVer, ProgramVer, StrVer))
  } else if LibVer != ProgramVer {
    ErrorHandler(fmt.Sprintf(Err_LibFileVer, ProgramVer, LibVer))
  }

  if len(os.Args) < 2 || len(os.Args) > 5 {
    usage() // 顯示用法文字後退出程式
  } else {
    for _, args := range os.Args {
      switch args {
      case "-r":
        recursive = true // 若使用者指定 -r 則開啟遞迴模式
      case "-v":
        verbose = true // 若使用者指定 -v 則開啟詳細模式
      case "--check-devel":
        Updater(ProgramVer, "devel")
        os.Exit(0)
      case "--check-stable":
        Updater(ProgramVer, "stable")
        os.Exit(0)
      }
    }
    process(os.Args[1], os.Args[2], recursive, verbose)
  }
}
