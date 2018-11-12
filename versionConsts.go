/*
 * Copy Large File (CLF) 更新軟體相關變數
 * 版本：v0.7.4 (不跟隨 main.go 更新)
 *
 * 編譯方式：  go build -o clf.out *.go
 * (Windows) go build -o clf.exe *.go
 * 如果是 Linux，則直接輸入 bash build.sh 即可。
 *
 * 語言：zh_TW。若要變更語系，請翻譯 string.go 中字串。
 */

package main

// GitHub 版本庫：master 分支上 raw 網址，尾端必須加 /
const updStable = "https://raw.githubusercontent.com/pan93412/CopyLargeFileGo/master/"

// GitHub 版本庫：devel 分支上 raw 網址，尾端必須加 /
const updDevel = "https://raw.githubusercontent.com/pan93412/CopyLargeFileGo/devel/"

// 下載 Releases 的網址。第一個 %s 預計放 release 版本、第二個預計放檔案名稱
const releaseDown = "https://github.com/pan93412/CopyLargeFileGo/releases/download/%s/%s"

// 下載 .zip 原始碼檔的網址。第一個 %s 預計放 release 版本
const archiveDown = "https://github.com/pan93412/CopyLargeFileGo/archive/%s.zip"

// 要下載檔案的檔名
var releaseDown_filename = map[string]string{
  "linux-386":     "clf-linux-386.out",     // Linux (i386)
  "linux-amd64":   "clf-linux-amd64.out",   // Linux (amd64)
  "freebsd-386":   "clf-freebsd-386.out",   // FreeBSD (i386)
  "freebsd-amd64": "clf-freebsd-amd64.out", // FreeBSD (amd64)
  "windows-386":   "clf-windows-386.exe",   // Windows (i386)
  "windows-amd64": "clf-windows-amd64.exe", // Windows (amd64)
  "darwin-amd64":  "clf-darwin-amd64.out",  // Darwin (macOS) (amd64)
}

const userAgent = `Mozilla/5.0 (X11; Linux x86_64; rv:62.0) Gecko/20100101 Firefox/62.0`
