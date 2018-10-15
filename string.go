/*
 * I18N 字串
 * for version: snapshot_20181015:1735(dev)
 */
 
package main

// 請勿修改此處。
const Version = "snapshot_20181015:1735(dev)"

// 第一個 %s：主程式版本；第二個 %s：程式檔名
var Usage = `=! Copy Large File != 版本 %s
用法：%s (原始檔案位置) (目標檔案位置) [-r] [-v]
() 為必須、[] 為選用。

[-r] (可選)：遞迴複製模式 (若原始檔案位置為 目錄，則必須)
[-v] (可選)：詳細輸出模式`

// 若 string.go 之 Version 無效時。第一個 %s：主程式版本；第二個 %s：語言檔版本
var Err_LanguageFileVer = "string.go 版本無效！此程式使用 %s 版，但字串檔案為 %s 版。"

// 若不支援現有系統的複製時。第一個 %s：系統名稱；第二個 %s：系統架構；第三個 %s：GitHub URL
var Err_NotSupport = "目前所使用的系統：%s (%s) 不在支援行列中。提交 Pull Request 到：%s\n"

// 若目標位置已經有檔案。
var err_dscExists = "目標位置已有檔案。"

// 若打算複製資料夾，卻沒使用遞迴 (-r) 引數
var err_FolderNotRecursive = "欲複製資料夾，但未使用 -r 引數。"

// 若目標或來源不是個檔案或資料夾、或其他錯誤。第一個 %s：GitHub URL
var err_unknownErrorWhenProcess = "目標或來源不是個檔案或資料夾、或發生了其他錯誤。若仍然不行，請回報到 %s。"


