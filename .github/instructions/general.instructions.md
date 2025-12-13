# General Instructions

## 目的
このファイルはリポジトリ共通の開発方針および自動化エージェント（Copilot 等）向けの簡潔な指示をまとめます。開発者と補助ツールが同じ前提で作業できるようにします。

## 想定読者
- リポジトリのコントリビュータ
- 自動化エージェント（Copilot 等）

## 技術スタック
- 言語: Go
- Web フレームワーク: Gin
- テンプレート: Go 標準の `html/template`
- ORM: GORM
- データベース移行: `golang-migrate`
- CSS: Tailwind CSS をフル導入し、UI ライブラリに daisyUI を使用
- アーキテクチャ: オニオンアーキテクチャ（Onion architecture）

## Tailwind + daisyUI（導入方針）
1. Tailwind をプロジェクトにフル導入する（開発依存として npm 管理）。
2. daisyUI は Tailwind のプラグインとして利用する。
3. 開発の最小手順（概略）:
   - `package.json` を用意し、`tailwindcss`, `postcss`, `autoprefixer`, `daisyui` を追加
   - `tailwind.config.js` で `daisyui` をプラグインに追加
   - ビルドパイプラインで CSS を生成し、テンプレートに組み込む

## マイグレーション方針
- スキーマ変更は `golang-migrate` を使用し、明示的なマイグレーションファイルを作成すること。
- GORM の `AutoMigrate` に頼り切らない（開発中の一時利用は可だが、本番は `golang-migrate` を優先）。
- `migrate` の基本コマンド例（ローカル開発）:

    migrate -path ./migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up

## オニオンアーキテクチャ（簡潔ガイド）
- レイヤー（内→外）: domain (entities, interfaces) → application (usecases) → infrastructure (DB, web)
- 依存ルール: 外側のレイヤーが内側に依存してはならない。インターフェースは内側に置く。
- GORM や Gin などの外部依存はインフラ層で実装し、アプリケーション層はインターフェース経由で呼び出す。

### 推奨ディレクトリ構成（例）
- `cmd/` - アプリケーションエントリ（main）
- `internal/domain/` - エンティティとドメインインターフェース
- `internal/application/` - ユースケース
- `internal/infrastructure/` - GORM リポジトリ、DB 接続、HTTP 層（Gin の初期化）
- `web/templates/` - `html/template` ファイル
- `web/assets/` - ビルド済み CSS/JS

## テンプレート方針
- Go の `html/template` を使用してサーバーサイドでレンダリングする。
- レイアウトテンプレートを採用して共通部品を分割する（ヘッダ、フッタ、コンポーネント）。

## GORM の使い方（方針）
- DB 操作はリポジトリ層に集約する（`internal/infrastructure/repository`等）。
- トランザクションはユースケースで制御し、必要に応じてリポジトリに受け渡す。

## テスト方針
- ユニットテストを重視する。外部DB依存はモック化もしくはインメモリ/テスト用 DB コンテナでの統合テストにする。
- 統合テストでは `golang-migrate` を使ってテスト DB にスキーマを適用する。

## CI / Lint / Formatting
- フォーマット: `gofmt`（または `gofumpt`）
- 静的解析: `go vet`, `golangci-lint` を推奨
- テスト実行: `go test ./...`

## PR チェックリスト
- 小さな差分に分けられているか
- 必要なテストが含まれているか（新機能なら新規テスト）
- マイグレーションがある場合、`migrations/` にファイルが追加されているか
- Lint とフォーマットに合致しているか
- HTML テンプレートで未エスケープの生データが出力されていないか

## Copilot / 自動化エージェント向け指示
- 言語: 日本語で簡潔に応答すること
- 出力方針: 小さな差分（最小限の変更）で提案すること
- コード提案時: 必ず Go の慣例に従い、必要ならテストコードを同時に提案すること
- データベース変更: `golang-migrate` 用のマイグレーションを提案し、直接 `AutoMigrate` に頼る変更は避けること
- フロントエンド: Tailwind + daisyUI を用いる前提でクラス名を生成すること
- テンプレート: `html/template` を使ったサーバーサイドレンダリングを前提にすること

## 付録: 早見コマンド（開発用）
- Tailwind をビルド（例）:

    npm install
    npx tailwindcss -i ./src/input.css -o ./web/assets/tailwind.css --watch

- マイグレーション適用:

    migrate -path ./migrations -database "$DATABASE_URL" up

---
このファイルはリポジトリの作業前提を示します。修正や追記があれば指示してください。
