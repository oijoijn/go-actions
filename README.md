### クリーンアーキテクチャ  何を考えるか                filetree
Entities                  domein                    todo構造体
↓
Use Cases                 business logic            app層(TodoService)
↓
Interface Adapters        db, web to connect        hander, ripository, eventpublisher
↓
Frameworks & Drivers      db, webframework itself   echo, gorm, PostgreSQL, nats itself


/internal/domain: Entities層。最も内側。
/internal/app: Use Cases層。ビジネスロジック。
/internal/infra:
    Interface Adapters 
    Frameworks層。GORMやEcho、NATSの実装。


### Interface Adapters層
ドメイン層で定義されたインターフェースをAdaptする具体的なコードが置かれる層

### Repository Pattern
ビジネスロジックとデータベースを分離するための最重要パターン
domein層 (内側): したいことの定義(interface)
infra層 (外側): code実装

### 依存性逆転の原則 (DIP)
Repository Patternのドメイン層に依存しようという考え方

### アウトボックスパターン
信頼性を劇的に高めるための超重要パターン
デュアルライト問題を解決するためのsql実行パターン
アトミック性
    結果が両方成功するか、両方失敗するか

データベースのトランザクションを利用してアトミック性を保証するパターン

### CQRS
システム内部構造をシンプルで効率的にするための設計パターン
データの書き込み, 読み取りで求められる要求が違う
これを1つのモデル（同じデータベース、同じテーブル、同じロジック）で両方やろうとすると、非常に複雑で中途半端なものが出来上がりがちです。
CQRSは書き込み読み取りの処理経路を、モデルレベルで完全に分離

コマンド側 (書き込みサイド)
    役割: システムの状態を変更する責任を持つ。
クエリ側 (読み取りサイド)
    役割: システムの状態を読み取る責任を持つ。

### 作り方
ドメイン層から外側に向かって作っていく
依存関係が少ない順に作ることで、一つ一つのパーツを独立させる

### ステップ1: 「何を作るか」を決める (ドメイン層) 🧱

1.  **モデル定義:** `internal/domain/model/todo.go` を作り、`Todo` というデータがどんな情報を持つのか（ID, Titleなど）を定義
2.  **リポジトリの「契約」定義:** `internal/domain/repository/todo_repository.go` を作る。ここでは「TODOを保存する」「TODOを探す」といった**機能のインターフェース（interface）**だけを定義します。
具体的なデータベースの処理はかかない。

---
### ステップ2: 「どうやるか」を決める (インフラ層) ⚙️

1.  **データベース処理:** `internal/infrastructure/persistence/gorm_todo_repository.go` を作り、ステップ1のインターフェースをGORMで実装します。
2.  **イベント発行処理:** `internal/infrastructure/messaging/nats_publisher.go` を作り、NATSでイベントを発行する具体的な処理を実装します。

---
### ステップ3: 「ビジネスの流れ」を作る (ユースケース層) 🧠

ドメインとインフラを繋ぎ、実際の業務ロジックを組み立てます。

1.  **ユースケース作成:** `internal/usecase/todo_usecase.go` を作り、リポジトリのインターフェースを使ってTODOを作成し、イベントを発行するといった一連の流れを記述します。

---
### ステップ4: 「リクエストの窓口」を作る (インターフェース層) 🔌
ユーザーからのHTTPリクエストを受け取り、ユースケースに処理を渡す部分です。

1.  **ハンドラ作成:** `internal/interface/handler/todo_handler.go` を作り、HTTPリクエストを解釈して、対応するユースケースを呼び出す処理を書きます。
2.  **ルーター設定:** `internal/infrastructure/router/router.go` で、どのURLがどのハンドラを呼び出すかを設定します。

---
### ステップ5: 全てを繋ぎ合わせる (`main.go` とコンテナ) 🚀
最後に、これまで作った全ての部品を組み立てて、アプリケーションを起動できるようにします。

1.  **`main.go`作成:** `cmd/api/main.go` で、設定ファイルの読み込み、DB接続、各層のインスタンス生成を行い、**依存関係の注入（DI）** をしてサーバーを起動します。
2.  **コンテナ設定:** `Dockerfile` と `docker-compose.yml` を作成し、`docker-compose up` コマンドだけで開発環境が立ち上がるようにします。
