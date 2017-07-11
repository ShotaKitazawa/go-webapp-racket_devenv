# About

Racket 用 Web 開発環境

- フォームにコードを書いて Execute で実行
- defined に 過去実行したコードの define を定義
  - defined に定義されている関数は以降の Execute 時にも反映
  - defined に定義された関数名は再定義できない
    - TODO: Clear ボタン実装
    - 即値としてなら利用可能

# 環境

- go version go1.8.3 darwin/amd64

# TODO

- Curry 化

- 完: Bootstrap でキレイに

---

- define_list の clear ボタン実装

- 改行時、括弧のネスト数だけスペースを挿入

- sourcecode フォーム のパースが \n 区切り > 括弧対応でパースできるようにする

- 実行速度の向上
  - string += string > strings.Join
  - 他色々
  - sourcecode からの入力をまずフォーマットしたほうがよさげ

- 悪いことをエスケープしなきゃ使い物にならない (例: ローカルのファイル参照)
  - 現在: フォームから来た文字列を racket -e コマンドに食わせている
