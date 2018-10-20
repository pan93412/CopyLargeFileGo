/*
 * Copy Large File (CLF) 函式庫
 * 版本：v0.7.1-beta
 */

package main
import (
  "log"
  "io"
  "os"
  "io/ioutil"
  "fmt"
)

const LibVer = "v0.7.1-beta"

// CopyFile 這個函式會將 src 檔案複製到 dst
// ，若成功回傳 error=nil，失敗則回傳 error!=nil
// 若開啟 ver(bose) 則使用詳細輸出模式。
func CopyFile(src string, dst string, ver bool) error {
  if ver {log.Printf(Info_StartCopy, src, dst)} // 若開啟 verbose，則顯示「開始複製」
  srcd, _ := os.OpenFile(src, os.O_RDONLY, 0444); defer srcd.Close() // 開啟 src
  dstd, _ := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0755); defer dstd.Close() // 建立 dst 檔案，權限跟 src 一樣
  cpbyte, err := io.Copy(dstd, srcd) // 複製檔案
  if err != nil { return err }
  if ver {log.Printf(Info_CopyCompleted, src, dst, cpbyte)} // 若開啟 verbose，則顯示「複製完成」
  return nil
}

// CopyDirectory 這個函式會遞迴將 src 目錄
// 複製到 dst 目錄，若成功則回傳 error=nil，
// 失敗則回傳 error!=nil。若開啟 ver(bose)
// 則使用詳細輸出模式。
func CopyDirectory(src string, dst string, ver bool) error {
  if ver {log.Printf(Info_StartCopy, src, dst)} // 若開啟 verbose，則顯示「開始複製」
  list, _ := ioutil.ReadDir(src) // 讀取 src 資料夾內的內容
  srcDat, _ := os.Stat(src) // 取得 src 的 *os.FileInfo 資訊
  os.Mkdir(dst, srcDat.Mode()) // 建立 dst 資料夾
  for _, name := range list {
    if name.IsDir() {
      CopyDirectory(src + "/" + name.Name(), dst + "/" + name.Name(), ver) // 若 src/(名稱) 仍為資料夾，則再次呼叫 CopyDirectory() 向下讀取
    } else {
      CopyFile(src + "/" + name.Name(), dst + "/" + name.Name(), ver) // 若 src/(名稱) 為檔案，則呼叫 CopyFile() 複製檔案到 dst
    }
  }
  if ver {log.Printf(Info_CopiedFolder, src, dst)} // 若開啟 verbose，則顯示「複製完成」
  return nil
}

// ErrorHandler 這個函式是用來替代 panic("")，不顯示
// 太多偵錯訊息，就是單純的報錯輸出。
func ErrorHandler(msg interface{}) {
  fmt.Printf("%s%v\n", ErrPrefix, msg)
  os.Exit(1)
}
