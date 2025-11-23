* 環境構築
  * nosqlworkbenchのインストール
  dynamoDBのローカル起動で一番楽ちんだと思ったので
  * jreのインストール
  nosqlworkbenchでローカル起動するために必要
  * aws cliのインストール
  dynamoDBのテーブル作成のために必要(nosqlworkbenchだと先にデータモデルを作る必要があり今回のケースだと面倒だった)
  cliからのテーブル作成はこちらを参考にした。
  https://future-architect.github.io/articles/20200225/

* gomockコマンド
pwshではだめでbashで実行する必要があった。
実行ディレクトリ：/c/git/go-sandbox/go-mock (main) 
コマンド： mockgen -source=service/foo.go -destination=./mock/mock_foo.go

* moqコマンド
実行ディレクトリ：/c/git/go-sandbox/go-mock/service (main)
コマンド：$ moq -out foo_moq.go . XxxService

* gomock vs moq
どっちもどっちという印象
gomockの方が生成コマンドがシンプル、moqは型安全らしい。
インターフェースのメソッド名を変えた場合moqは追随できる（ただ再生成するとmoqで生成されるメソッド名が変わってしまうのでそのタイミングで結局ダメになりそう）

* Mockio
interfaceさえ作っておけばDIも必要なく処理を書き換えることができ最強  
ただ実行環境やコンパイラの内部構造に依存しており将来的に使えなくなる可能性がある。  
※[Add support for code generation #79](https://github.com/ovechkin-dm/mockio/issues/79)このIssueの解消を待つか自分で対処方法を検討しておく必要がありそう。（そのタイミングで別Mockライブラリへの切替もある程度機械的な修正で切替できそうな気もする）

* GoのMockツールの比較情報
https://gist.github.com/maratori/8772fe158ff705ca543a0620863977c2

* ビルド実行

