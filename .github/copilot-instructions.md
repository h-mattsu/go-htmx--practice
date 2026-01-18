# プロジェクト指示書

## 技術スタック
- **バックエンド**: Go 1.x + Gin Web Framework
- **テンプレート**: Go html/template
- **フロントエンド**: htmx + daisyUI (Tailwind CSS)
- **アーキテクチャ**: レイヤードアーキテクチャ（Clean Architecture風）

## プロジェクト構造

```
/workspaces/go-htmx-practice/
├── application/services/    # アプリケーションロジック・ユースケース
├── domain/models/           # ドメインモデル・エンティティ
├── infrastructure/repository/ # データアクセス層
├── presentation/
│   ├── handlers/           # HTTPハンドラー
│   │   ├── handler.go     # すべてのハンドラーをまとめるメインハンドラー
│   │   ├── page_handler.go # ページ関連のハンドラー
│   │   └── user_handler.go # ユーザー関連のハンドラー（例）
│   └── templates/         # Goテンプレート
│       ├── layouts/       # レイアウトテンプレート
│       ├── pages/         # ページテンプレート
│       ├── partials/      # 部分テンプレート
│       └── components/    # 再利用可能なコンポーネント
└── main.go               # エントリーポイント
```

## ハンドラー構造

### ファイル構成
- `handlers/handler.go`: すべてのハンドラーをまとめ、ルーティングを定義
- `handlers/<機能>_handler.go`: 各機能ごとのハンドラー実装

### ハンドラーの構成例
```go
// handler.go - メインハンドラー
type Handler struct {
    PageHandler *PageHandler
    UserHandler *UserHandler
}

func NewHandler(pageService *services.PageService, userService *services.UserService) *Handler {
    return &Handler{
        PageHandler: NewPageHandler(pageService),
        UserHandler: NewUserHandler(userService),
    }
}

func (h *Handler) SetupRoutes(router *gin.Engine) {
    // ルーティング設定
}
```

```go
// page_handler.go - ページハンドラー
type PageHandler struct {
    pageService *services.PageService
}

func NewPageHandler(pageService *services.PageService) *PageHandler {
    return &PageHandler{pageService: pageService}
}
```

## コーディング規約

### Go全般
- 標準的なGoのコーディングスタイルに従う（gofmt, golint準拠）
- エラーハンドリングは明示的に行う
- コンテキストは第一引数として渡す
- パッケージ名は小文字の単数形を使用

### レイヤー責務
- **presentation**: HTTPリクエスト/レスポンスの処理、テンプレートレンダリング
- **application/services**: ビジネスロジック、トランザクション制御
- **domain/models**: ドメインモデル、ビジネスルール
- **infrastructure/repository**: データベースアクセス、外部API連携

### 依存関係の方向
- presentation → application → domain ← infrastructure
- 下位レイヤーは上位レイヤーに依存しない
- infrastructureはinterfaceを通じてdomainに依存

## テンプレート規約

### ファイル命名
- レイアウト: `layouts/base.go.tmpl`, `layouts/admin.go.tmpl`
- ページ: `pages/home.go.tmpl`, `pages/users.go.tmpl`
- パーシャル: `partials/navbar.go.tmpl`, `partials/footer.go.tmpl`
- コンポーネント: `components/button.go.tmpl`, `components/modal.go.tmpl`

### テンプレート構造
- 全てのページは基本レイアウトを継承
- `{{define "ページ名"}}` でページコンテンツを定義
- `{{template "コンポーネント名" .}}` で部分テンプレートを読み込み

### データ渡し
- ハンドラーからテンプレートへのデータは構造体で渡す
- テンプレート用のViewModelを定義することを推奨

## htmx使用パターン

### 基本方針
- ページ全体のリロードを避け、部分更新を優先
- `hx-get`, `hx-post`, `hx-put`, `hx-delete` を適切に使い分け
- `hx-target` で更新対象を明示
- `hx-swap` でスワップ方法を指定（innerHTML, outerHTML, beforeend等）

### エンドポイント設計
- 完全なページ: `/users` (GET)
- 部分更新用: `/users/list` (htmxリクエスト用のフラグメント)
- `c.GetHeader("HX-Request")` でhtmxリクエストを判定

### レスポンスパターン
```go
if c.GetHeader("HX-Request") != "" {
    // htmxリクエスト: 部分テンプレートのみ返す
    c.HTML(http.StatusOK, "partials/user-list", data)
} else {
    // 通常リクエスト: 完全なページを返す
    c.HTML(http.StatusOK, "pages/users", data)
}
```

## daisyUI使用ガイドライン

### コンポーネント使用
- daisyUIのコンポーネントクラスを優先使用
- カスタムスタイルが必要な場合はTailwindユーティリティクラスを追加
- 主要コンポーネント: `btn`, `card`, `modal`, `navbar`, `form-control`, `input`, `select`

### テーマ
- `data-theme` 属性でテーマを指定
- ダークモード対応を考慮

### レスポンシブデザイン
- Tailwindのレスポンシブプレフィックスを使用: `sm:`, `md:`, `lg:`, `xl:`

## ハンドラー実装パターン

### 基本的なハンドラー実装
```go
// user_handler.go
type UserHandler struct {
    userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
    // 1. リクエストパラメータの取得とバリデーション
    
    // 2. サービス層の呼び出し
    users, err := h.userService.GetAll(c.Request.Context())
    if err != nil {
        c.HTML(http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
        return
    }
    
    // 3. ViewModelの構築
    viewModel := UsersViewModel{
        Users: users,
        Title: "ユーザー一覧",
    }
    
    // 4. htmxリクエストの判定とレンダリング
    if c.GetHeader("HX-Request") != "" {
        c.HTML(http.StatusOK, "partials/user-list", viewModel)
    } else {
        c.HTML(http.StatusOK, "pages/users", viewModel)
    }
}
```

## エラーハンドリング

- エラーは適切なHTTPステータスコードで返す
- ユーザー向けエラーメッセージはテンプレートで表示
- htmxリクエストのエラーは `HX-Trigger` ヘッダーでイベント通知も検討

## セキュリティ
- CSRFトークンをフォームに含める
- XSS対策: html/templateの自動エスケープを活用
- バリデーションはサーバーサイドで必ず実施

## パフォーマンス
- テンプレートは起動時に一度だけパース
- 必要に応じてキャッシュ機構を実装
- htmxで不要なデータ転送を削減

## 命名規則
- ハンドラーファイル: `<機能>_handler.go` (例: `user_handler.go`, `page_handler.go`)
- ハンドラー構造体: `<機能>Handler` (例: `UserHandler`, `PageHandler`)
- ハンドラー関数: `GetUsers`, `CreateUser`, `UpdateUser`, `DeleteUser`
- サービスファイル: `<機能>_service.go` (例: `user_service.go`, `page_service.go`)
- サービスメソッド: `GetAll`, `GetByID`, `Create`, `Update`, `Delete`
- URL: `/users`, `/users/:id`, `/users/new`, `/users/:id/edit`
- テンプレート定義名: `pages/users`, `partials/user-card`
