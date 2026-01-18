# Go + htmx + daisyUI Practice

Go言語でWebアプリケーションを開発する練習用プロジェクトです。サーバーサイドレンダリング（SSR）とモダンなフロントエンドインタラクションを組み合わせたアーキテクチャを学習します。

## 技術スタック

- **バックエンド**: Go 1.x + Gin Web Framework
- **テンプレート**: Go html/template
- **フロントエンド**: htmx（動的インタラクション）
- **UI**: daisyUI + Tailwind CSS
- **アーキテクチャ**: レイヤードアーキテクチャ（Clean Architecture風）

## プロジェクト構造

```
/workspaces/go-htmx-practice/
├── application/services/    # ビジネスロジック・ユースケース
├── domain/models/           # ドメインモデル・エンティティ
├── infrastructure/repository/ # データアクセス層
├── presentation/
│   ├── handlers/           # HTTPリクエストハンドラー
│   ├── router.go          # ルーティング設定
│   └── templates/         # Goテンプレート
│       ├── layouts/       # ベースレイアウト
│       ├── pages/         # ページ全体のテンプレート
│       ├── partials/      # 部分テンプレート（htmx用）
│       └── components/    # 再利用可能なコンポーネント
└── main.go               # アプリケーションエントリーポイント
```

## 主な特徴

- **レイヤードアーキテクチャ**: 関心の分離により保守性と拡張性を向上
- **htmxによる部分更新**: JavaScriptを最小限に抑え、HTMLベースのインタラクション
- **daisyUI**: Tailwind CSSベースの美しいUIコンポーネント
- **サーバーサイドレンダリング**: シンプルで高速なWebアプリケーション

## セットアップ

### 前提条件

- Go 1.x以上
- Git

### インストール

```bash
# リポジトリのクローン
git clone <repository-url>
cd go-htmx-practice

# 依存関係のインストール
go mod download
```

### 実行

```bash
# 開発サーバーの起動
go run main.go
```

ブラウザで http://localhost:8080 を開いてアプリケーションにアクセスできます。

## 開発ガイド

プロジェクトの詳細なコーディング規約やアーキテクチャパターンについては、[.copilot-instructions.md](.copilot-instructions.md) を参照してください。

### htmxの使い方

htmxリクエストと通常のリクエストを判別して、適切なテンプレートを返します：

```go
if c.GetHeader("HX-Request") != "" {
    // 部分テンプレートのみ返す
    c.HTML(http.StatusOK, "partials/user-list", data)
} else {
    // 完全なページを返す
    c.HTML(http.StatusOK, "pages/users", data)
}
```

### テンプレートの命名規則

- レイアウト: `layouts/base.go.tmpl`
- ページ: `pages/home.go.tmpl`
- パーシャル: `partials/navbar.go.tmpl`
- コンポーネント: `components/button.go.tmpl`

## ライセンス

このプロジェクトは学習目的で作成されています。
