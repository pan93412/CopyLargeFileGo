package main

// CopyFile 這個函式會將 src 檔案複製到 dst
// ，若成功回傳 error=nil，失敗則回傳 error!=nil
// 若開啟 ver(bose) 則使用詳細輸出模式。
func CopyFile(src string, dst string, ver bool) error {}

// CopyDirectory 這個函式會遞迴將 src 目錄
// 複製到 dst 目錄，若成功則回傳 error=nil，
// 失敗則回傳 error!=nil。若開啟 ver(bose)
// 則使用詳細輸出模式。
func CopyDirectory(src string, dst string, ver bool) error {}
