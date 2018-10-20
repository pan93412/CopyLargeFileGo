/*
 * Copy Large File (CLF) 字串檔案
 * 版本：v0.7.0-beta
 */
 
package main

// 請勿修改此處。
const StrVer = "v0.7.0-beta"

// 第一個 %s：主程式版本；第二個 %s：程式檔名
const Usage = `=! Copy Large File != 版本 %s
用法：%s (原始檔案位置) (目標檔案位置) [-r] [-v] [--check-{devel|stable}]
() 為必須、[] 為選用。

[-r] (可選)：遞迴複製模式 (若原始檔案位置為 目錄，則必須)
[-v] (可選)：詳細輸出模式
[--check-devel] (可選)：檢查 devel 分支上的更新
[--check-stable] (可選)：檢查 master 分支上的更新

此軟體 GitHub 網址：<http://www.github.com/pan93412/CopyLargeFileGo>
`

/* 
 * libs.go 與 main.go 字串區
 * 若能，請優先翻譯此區塊。
 * 調用時機：當使用者開啟編譯後程式
 */
  
// 若 string.go 之 Version 無效時。第一個 %s：主程式版本；第二個 %s：語言檔版本
const Err_LanguageFileVer = "string.go 版本無效！此程式使用 %s 版，但字串檔案為 %s 版。"

// 若 libs.go 之 Version 無效時。第一個 %s：主程式版本；第二個 %s：函式庫版本
const Err_LibFileVer = "libs.go 版本無效！此程式使用 %s 版，但函式庫版本為 %s 版。"

// 若目標位置已經有檔案。
const Err_dscExists = "目標位置已有檔案。"

// 若來源位置沒有檔案或資料夾。
const Err_srcNotExists = "來源位置沒有檔案或資料夾。"

// 若打算複製資料夾，卻沒使用遞迴 (-r) 引數
const Err_FolderNotRecursive = "欲複製資料夾，但未使用 -r 引數。"

// 若目標或來源不是個檔案或資料夾、或其他錯誤。第一個 %s：GitHub URL
const Err_unknownErrorWhenProcess = "目標或來源不是個檔案或資料夾、或發生了其他錯誤。"

// 發生錯誤時顯示的 GitHub Issue Tracker 網址。
const Err_GitHubIT = "\n若仍然不行，請回報錯誤到 <http://www.github.com/pan93412/CopyLargeFileGo/issues>。"

// 開始複製的訊息。第一個 %s：來源檔案名稱；第二個 %s：目標檔案名稱
const Info_StartCopy = "開始複製 %s 到 %s 的程序。\n"

// 複製檔案完成的訊息。第一個 %s：來源檔案名稱；第二個 %s：目標檔案名稱；第三個 %d：複製的 bytes。
const Info_CopyCompleted = "從 %s 檔案複製到 %s 完成，共複製 %d 位元組。\n"

// 複製資料夾完成的訊息。第一個 %s：來源檔案名稱；第二個 %s：目標檔案名稱
const Info_CopiedFolder = "從 %s 資料夾複製到 %s 完成。\n"

// 錯誤訊息的 prefix。
const ErrPrefix = "錯誤："

/*
 * version.go 字串區
 * 非重要區域。優先級第二。
 * 調用時機：當使用者執行：(執行檔案) --check-stable / --check-devel
 */

// 通常這字串不應該出現，若出現請發 Issue 到 Issue Tracker
const err_BranchInvaild = "所選取分支無效，請選擇正確分支。" + Err_GitHubIT

// 若接收到的 HTTP Status Code 不是 200
const err_HTTPSCErr = "URL 回傳代碼非 200，可能是因為分支網址錯誤、或是目標網站發生問題。" + Err_GitHubIT

// 接收到更新時的訊息
// 第一個 %s：伺服器上最新版本；第二個 %s：本機版本
// 第三個 %s：最新版本之發布日期；第四個 %s：最新版本之更新日誌
// 第五個 %s：引數 (devel/stable)
const updReceived = `找到新更新！
版本號碼：%s (目前版本：%s)
發布日期：%s
更新日誌：
%s

若要更新，請按下 [Enter] 更新，
反之，請按下 [Ctrl-C] 中止程式，並
不再加上 --check-%s 選項。
`

// 若目前版本為最新版本。第一個 %s：目前版本號碼
const nowVersionLatest = "目前版本 (%s) 為最新版本，不需更新 :-)\n"

// 若您的作業系統有可用的二進位檔案。第一個 %s：作業系統名稱；第二個 %s：作業系統架構
const supportYourComputer = "此自動更新程式支援您的作業系統：%s (%s)，將自動為您下載二進位檔案。\n下載後重新啟動程式即為最新版本！\n"

// 若您的作業系統沒有可用的二進位檔案。第一個 %s：作業系統名稱；第二個 %s：作業系統架構
const notSupportYourComputer = `此自動更新程式不支援您的作業系統：%s (%s)。

請下載適用於您作業系統的 Go 編譯工具：
https://golang.org/doc/install

稍候下載的原始碼檔案為 %s，解壓縮後即可編譯。` + "\n"
