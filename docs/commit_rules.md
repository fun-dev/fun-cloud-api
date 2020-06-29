# Commit Rules
## TL;DR（とりあえずここ読んで！！）
- Commit Template
    - `git commit -m [${commit category}] ${commit message}` 
    - ex: `git commit -m [fix] hogehoge`

## Commit Category
- 機能の追加のタスク完了時：コミットの前に `feature`と付ける
ex): `git commit -m "[feature] $何をしたか"`
- プログラムのファイルを追加したとき：`add` 
- プログラムには関係ないファイルを作ったとき（設定ファイル等）：`chore`
- プログラムに修正をかけた場合：`fix`