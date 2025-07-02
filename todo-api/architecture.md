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
