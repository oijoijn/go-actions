package main

import (
	"context"
	// "fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-app/internal/config"
	"todo-app/internal/domain/model"
	"todo-app/internal/infrastructure/messaging"
	"todo-app/internal/infrastructure/persistence"
	"todo-app/internal/infrastructure/router"
	"todo-app/internal/interface/handler"
	"todo-app/internal/usecase"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 1. 設定の読み込み
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // 環境変数を優先
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// 2. データベース接続
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// 自動マイグレーション
	db.AutoMigrate(&model.Todo{})

	// 3. NATS JetStream接続
	nc, err := nats.Connect(cfg.Nats.URL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Failed to get JetStream context: %v", err)
	}
	// ストリームを作成（存在しない場合）
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     cfg.Nats.StreamName,
		Subjects: cfg.Nats.Subjects,
	})
	if err != nil {
		log.Fatalf("Failed to add NATS stream: %v", err)
	}

	// 4. 依存関係の注入 (Dependency Injection)
	todoRepo := persistence.NewGormTodoRepository(db)
	eventPublisher := messaging.NewNatsPublisher(js)
	todoUsecase := usecase.NewTodoUsecase(todoRepo, eventPublisher)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	// 5. ルーターの設定
	e := router.NewRouter(todoHandler)

	// 6. サーバーの起動とグレースフルシャットダウン
	go func() {
		if err := e.Start(":" + cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// シャットダウンシグナルを待つ
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("Server gracefully stopped")
}
