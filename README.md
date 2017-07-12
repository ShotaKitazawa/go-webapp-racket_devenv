# About

Racket 用 Web 開発環境

- フォームにコードを書いて Execute で実行
  - 裏で `racket -e "Source"` が動いている
- defined に 過去実行したコードの define を定義
  - defined に定義されている関数は以降の Execute 時にも反映
  - defined に定義された関数名は再定義できない
    - 即値としてなら利用可能
    - Clear にて関数削除
  - Curry で関数をCurry化

# 実行方法

- `make` コマンドで main 実行ファイルを生成
- コマンドラインにて main を実行
  - http://localhost:8000/ にてアクセス可能

# 環境

- macOS 10.12.5
  - Darwin Kernel Version 16.6.0
- go version go1.8.3 darwin/amd64
- Racket v6.9
- Bootstrap v3.3.7

# TODO

- 完: Curry 化

- 完: Bootstrap でキレイに

- 完: define_list の clear ボタン実装

- 完: 不完全な形の関数も Defined に突っ込んでしまう
  - 例: `(define (fc x y) ()`

- #lang による言語切替に対応するため、tmp.rktファイルにSourceを書き込んだあと、`racket tmp.rkt` で実行出来るようにする

- 出力結果の見直し
  - 改行コードが変なところに入っている?

- SourceCodeフォームにてクライアントが改行時、括弧のネスト数だけスペースを挿入

- sourcecode フォーム のパースが \n 区切り > 括弧対応でパースできるようにする

- 実行速度の向上
  - string += string > strings.Join
  - 他色々

- sourcecode からの入力をまずフォーマットしたほうがよさそう
  - Execute 時に以下をする
    - 1つ以上のスペースを1つのスペースに置換
    - ( , ) の周りのSpace排除

- 悪いことをエスケープしなきゃマズい (例: ローカルのファイル参照)
  - 現在: フォームから来た文字列を racket -e コマンドに食わせている
