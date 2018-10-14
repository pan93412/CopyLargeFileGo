/*
 * I18N 字串
 * for version: snapshot_20181015:0053(dev)
 */
 
package main

// 請勿修改此處。
const Version = "snapshot_20181015:0053(dev)"

// 第一個 %s：主程式版本；第二個 %s：程式檔名
var Usage = `=! Copy Large File != 版本 %s
用法：%s (原始檔案位置) (目標檔案位置) [-r] [-v]
() 為必須、[] 為選用。

[-r] (可選)：遞迴複製模式 (若原始檔案位置為 目錄，則必須)
[-v] (可選)：詳細輸出模式`

// 第一個 %s：主程式版本；第二個 %s：語言檔版本
var Err_LanguageFileVer = "string.go 版本無效！此程式使用 %s 版，但字串檔案為 %s 版。"

// 第一個 %s：系統名稱；第二個 %s：系統架構；第三個 %s：GitHub URL
var Err_NotSupport = "目前所使用的系統：%s (%s) 不在支援行列中。提交 Pull Request 到：%s\n"
