# Backupシステム

## 設計

- ファイルのスナップショットを定期的に作成し、ソースコードを含むプロジェクトへの変更を記録する
- 変更の有無をチェックする間隔を変更できる
- 主にテキストベースのプロジェクトをZIP圧縮する
- ビルドは早期に行いつつ、将来的に改善の可能性を検討する
- 実装上の判断は容易に修正できるようにし、今後の変更に備える
- 2つのコマンドラインツール
  - 1つは実際に処理を行うバックエンドのデーモン
  - 1つはバックアップ対象のパスの一覧商事や追加・削除を行うユーザ向けのユーティリティ

## 参照

[Go言語によるWebアプリケーション開発 Mat Ryer著、鵜飼 文敏 監訳、牧野 聡 訳](https://www.oreilly.co.jp/books/9784873117522/)