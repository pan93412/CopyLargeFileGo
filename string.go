/*
 * Copy Large File (CLF) 字串檔案
 * 版本：v0.6_beta
 */
 
package main

// 請勿修改此處。
const StrVer = "v0.6_beta"

// 第一個 %s：主程式版本；第二個 %s：程式檔名
const Usage = `=! Copy Large File != 版本 %s
用法：%s (原始檔案位置) (目標檔案位置) [-r] [-v]
() 為必須、[] 為選用。

[-r] (可選)：遞迴複製模式 (若原始檔案位置為 目錄，則必須)
[-v] (可選)：詳細輸出模式

此軟體 GitHub 網址：<http://www.github.com/pan93412/CopyLargeFileGo>
`

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
