# Copy Large File (CLF)
這是一個能遞迴複製資料夾、或只單一複製檔案的一個程式。

## 命令列
```
=! Copy Large File != 版本 v0.5_beta
用法：(程式名稱) (原始檔案位置) (目標檔案位置) [-r] [-v]
() 為必須、[] 為選用。

[-r] (可選)：遞迴複製模式 (若原始檔案位置為 目錄，則必須)
[-v] (可選)：詳細輸出模式

此軟體 GitHub 網址：<http://www.github.com/pan93412/CopyLargeFileGo>
```

## 速度
```
[pan93412@archlinux ~]$ du -sh ~/下載
824M	/home/pan93412/下載

[pan93412@archlinux CopyLargeFile]$ time ./clf.out ~/下載 ~/test_folder -r

real  0m16.591s
user  0m0.209s
sys  0m4.684s
```

## 作者
- pan93412 \<<https://github.com/pan93412>\>, 2018.
- Yami Odymel (技術指導) \<<https://github.com/yamiodymel>\>, 2018.
- 和其他朋友們 ;-)
