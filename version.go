/*
 * Copy Large File (CLF) 更新軟體
 * 版本：v0.7_beta (不跟隨 main.go 更新)
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
  "net/http"
  "os"
  "io"
  "io/ioutil"
  "errors"
  JSONProc "encoding/json"
  "runtime"
)

// 相關常數定義
const updStable = "https://raw.githubusercontent.com/pan93412/CopyLargeFileGo/master/" // GitHub 版本庫：master 分支上 raw 網址，尾端必須加 /
const updDevel = "https://raw.githubusercontent.com/pan93412/CopyLargeFileGo/devel/" // GitHub 版本庫：devel 分支上 raw 網址，尾端必須加 /
const versionJSON = "Version.json" // 版本資訊 JSON 檔名
const releaseDown = "https://github.com/pan93412/CopyLargeFileGo/releases/download/%s/%s" // 下載 Releases 的網址。第一個 %s 預計放 release 版本、第二個預計放檔案名稱
const archiveDown = "https://github.com/pan93412/CopyLargeFileGo/archive/%s.zip" // 下載 .zip 原始碼檔的網址。第一個 %s 預計放 release 版本 
const userAgent = `Mozilla/5.0 (X11; Linux x86_64; rv:62.0) Gecko/20100101 Firefox/62.0`
var releaseDown_filename = map[string]string{
  "linux-386": "clf-linux-386.out",          // Linux (i386)
  "linux-amd64": "clf-linux-amd64.out",      // Linux (amd64)
  "freebsd-386": "clf-freebsd-386.out",      // FreeBSD (i386)
  "freebsd-amd64": "clf-freebsd-amd64.out",  // FreeBSD (amd64)
  "windows-386": "clf-windows-386.exe",      // Windows (i386)
  "windows-amd64": "clf-windows-amd64.exe",  // Windows (amd64)
  "darwin-amd64": "clf-darwin-amd64.out",    // Darwin (macOS) (amd64)
}

// System 是用來擺放各發行版的
// VersionInfo 為 Version.json 的內容建構體
type VersionInfo struct {
  NowVer string   `json:"currentVersion"` // 目前版本
  RelDat string   `json:"releaseDate"`    // 釋出日期
  UpdLog string   `json:"updateLog"`      // 更新日誌
}

// download 模組可以從遠端伺服器下載 URL 檔案。
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
  if err != nil { return []byte(""), err }
  // 增加 Firefox 的 User-Agent
  rawHTTPRequest.Header.Add("User-Agent", userAgent)
  // 發送請求
  rawHTTPResponse, err = clientServ.Do(rawHTTPRequest)
  if err != nil { return []byte(""), err }
  if rawHTTPResponse.StatusCode != 200 { return []byte(""), errors.New(err_HTTPSCErr) }
  rawData = rawHTTPResponse.Body; defer rawData.Close()
  data, _ := ioutil.ReadAll(rawData)
  return data, nil
}
  
// checkUpdates 可以從遠端伺服器抓取 Version.json，並解析
// Version.json，並檢查目前版本是否最新。
//
// 第一個傳回參數：檢查更新過程中是否有錯誤
// 第二個傳回參數：要給使用者看的更新資訊
// 第三個傳回參數：判斷是否要更新
func checkUpdates(currentVer string, branch string, fetchWhere string) (error, string, bool, *VersionInfo){
  // 初始化變數
  var err error
  var jsonRaw []byte
  var theVersionInfo = &VersionInfo{}
  
  // 解析 Version.json
  jsonRaw, err = downloadFile(fetchWhere + versionJSON)
  if err != nil { return err, "", false, theVersionInfo}
  err = JSONProc.Unmarshal(jsonRaw, &theVersionInfo)
  if err != nil { return err, "", false, theVersionInfo}
  
  // 輸出結果
  if theVersionInfo.NowVer != currentVer {
    return nil, fmt.Sprintf(updReceived, theVersionInfo.NowVer, currentVer, theVersionInfo.RelDat, theVersionInfo.UpdLog, branch), true, theVersionInfo
  } else {
    return nil, fmt.Sprintf(nowVersionLatest, currentVer), false, theVersionInfo
  }
}
  
// Updater 可從 GitHub 庫抓取最新 CLF 程式，判斷系統
// 種類並下載對應二進位檔案
func Updater(currentVer string, branch string) error {
  // 初始化變數
  var fetchWhere, updMsg, scanStr, downWhat string
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

  err, updMsg, needUpdate, verInf = checkUpdates(currentVer, branch, fetchWhere)
  if err != nil { panic(err) }
  fmt.Println(updMsg)
  if needUpdate == true {
    fmt.Scanln(&scanStr)
    for id, version := range releaseDown_filename {
      if id == runtime.GOOS + "-" + runtime.GOARCH {
        downWhat = version
      }
    }
    if downWhat != "" {
      fmt.Printf(supportYourComputer, runtime.GOOS, runtime.GOARCH)
      file, err = downloadFile(fmt.Sprintf(releaseDown, verInf.NowVer, downWhat))
      filename := os.Args[0] + ".new"
      if _, err := os.Stat(filename); err == nil {
          os.Remove(filename)
        }
      clf_cachedat, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
      clf_cachedat.Write(file)
      clf_cachedat.Close()
      os.Remove("./" + os.Args[0])
      os.Rename("./" + os.Args[0] + ".new", "./" + os.Args[0])
      os.Exit(0)
    } else {
      theFilename := verInf.NowVer + ".zip"
      fmt.Printf(notSupportYourComputer, runtime.GOOS, runtime.GOARCH, theFilename)
      if _, err := os.Stat(theFilename); err == nil {
        os.Remove(theFilename)
      }
      file, err = downloadFile(fmt.Sprintf(archiveDown, verInf.NowVer))
      clf_archive, _ := os.OpenFile(theFilename, os.O_WRONLY|os.O_CREATE, 0755)
      clf_archive.Write(file)
      clf_archive.Close()
      os.Exit(0)
    }
  }
  return nil
}
