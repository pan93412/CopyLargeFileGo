/*
 * Copy Large File (CLF) 更新軟體
 * 版本：v0.7.4 (不跟隨 main.go 更新)
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
  JSONProc "encoding/json" // 用來解析 JSON
  "errors"                 // 回傳 error 訊息
  "fmt"                    // 輸出文字
  "io"                     // 僅用於定義變數型態
  "io/ioutil"              // 讀取檔案
  "net/http"               // 下載內容
  "os"                     // 寫入到檔案
  "runtime"                // 顯示系統版本與架構
)

// 定義相關檔案
const verInfoFile = "Version.json" // 版本資訊 JSON 檔名

// VersionInfo 為 Version.json 的內容建構體
type VersionInfo struct {
  NowVer string `json:"currentVersion"` // 目前版本
  RelDat string `json:"releaseDate"`    // 釋出日期
  UpdLog string `json:"updateLog"`      // 更新日誌
}

// downloadFile 模組可以從遠端伺服器下載 theURL 檔案。
// []byte: 下載到的檔案資料
// error: 下載時發生的錯誤
func downloadFile(theURL string) ([]byte, error) {
  // 定義變數
  var rawHTTPRequest *http.Request
  var rawHTTPResponse *http.Response
  var clientServ = &http.Client{}
  var rawData io.ReadCloser
  var err error

  // 建立請求
  rawHTTPRequest, err = http.NewRequest("GET", theURL, nil)
  if err != nil {
    return []byte(""), err
  }

  // 增加 User-Agent
  rawHTTPRequest.Header.Add("User-Agent", userAgent)

  // 發送請求
  rawHTTPResponse, err = clientServ.Do(rawHTTPRequest)
  if err != nil {
    return []byte(""), err
  }
  if rawHTTPResponse.StatusCode != 200 {
    return []byte(""), errors.New(err_HTTPSCErr)
  }
  rawData = rawHTTPResponse.Body
  defer rawData.Close()
  data, _ := ioutil.ReadAll(rawData)
  return data, nil
}

// checkUpdates 可以從遠端伺服器抓取 Version.json，並解析
// Version.json，並檢查目前版本是否最新。
//
// 第一個傳回參數：檢查更新過程中是否有錯誤
// 第二個傳回參數：判斷是否要更新
// 第三個傳回參數：version.json 資料
func checkUpdates(currentVer string, branch string, fetchWhere string) (error, bool, *VersionInfo) {
  // 初始化變數
  var err error
  var jsonRaw []byte
  var theVersionInfo = &VersionInfo{}

  // 解析 Version.json
  jsonRaw, err = downloadFile(fetchWhere + verInfoFile)
  if err != nil {
    return err, false, theVersionInfo
  }
  err = JSONProc.Unmarshal(jsonRaw, &theVersionInfo)
  if err != nil {
    return err, false, theVersionInfo
  }

  // 回傳結果
  if theVersionInfo.NowVer != currentVer {
    return nil, true, theVersionInfo
  } else {
    return nil, false, theVersionInfo
  }
}

// Updater 可從 GitHub 庫抓取最新 CLF 程式，判斷系統
// 種類並下載對應二進位檔案
func Updater(currentVer string, branch string) error {
  // 初始化變數
  var fetchWhere, scanStr string
  var err error
  var needUpdate bool
  var file []byte
  var verInf *VersionInfo

  // 判斷是否要用 stable 分支還是 devel 分支
  switch branch {
  case "stable":
    fetchWhere = updStable
  case "devel":
    fetchWhere = updDevel
  default:
    ErrorHandler(err_BranchInvaild)
  }

  // 使用 checkUpdates 檢查更新
  err, needUpdate, verInf = checkUpdates(currentVer, branch, fetchWhere)
  if err != nil {
    PanicHandler(err)
  }

  // 如果需要更新
  if needUpdate {
    // 顯示需要更新訊息
    fmt.Printf(updReceived, verInf.NowVer, currentVer, verInf.RelDat, verInf.UpdLog, branch)
    fmt.Scanln(&scanStr) // 等待使用者按下 [Enter]

    // 判斷系統
    if downfile, ok := releaseDown_filename[runtime.GOOS+"-"+runtime.GOARCH]; ok {
      filename := os.Args[0] + ".new"                               // 預計檔案名稱為 (原執行檔案名).new
      fmt.Printf(supportYourComputer, runtime.GOOS, runtime.GOARCH) // 顯示「支援您的電腦」訊息

      // 調用下載檔案函數來下載二進位檔案
      file, err = downloadFile(fmt.Sprintf(releaseDown, verInf.NowVer, downfile))
      // 如預計檔案名稱存在，則移除該檔案
      if _, err := os.Stat(filename); err == nil {
        os.Remove(filename)
      }

      // 開始將下載內容寫入檔案
      cache, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
      cache.Write(file)
      defer cache.Close()

      // 刪除原執行檔案
      os.Remove("./" + os.Args[0])
      // 將 (執行檔案名稱).new 改成原執行檔案名稱
      os.Rename("./"+os.Args[0]+".new", "./"+os.Args[0])
    } else {
      theFilename := verInf.NowVer + ".zip" // 預計檔案名稱為 (新版本號碼).zip
      // 顯示「不支援您的電腦系統，請自行編譯」
      fmt.Printf(notSupportYourComputer, runtime.GOOS, runtime.GOARCH, theFilename)

      // 下載原始碼檔案
      file, err = downloadFile(fmt.Sprintf(archiveDown, verInf.NowVer))

      // 如預計檔案名稱存在，則移除該檔案
      if _, err := os.Stat(theFilename); err == nil {
        os.Remove(theFilename)
      }

      // 開始寫入 downloadFile() 得到的結果
      cache, _ := os.OpenFile(theFilename, os.O_WRONLY|os.O_CREATE, 0755)
      cache.Write(file)
      cache.Close()
    }
    fmt.Println(downloadDone)
    return nil
  }
  // 如果沒有更新，則什麼也不做。
  fmt.Printf(nowVersionLatest, ProgramVer)
  return nil
}
