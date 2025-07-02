# 1. ビルドステージ
FROM golang:1.22-alpine AS builder

WORKDIR /src

# go.modだけ先にコピー
COPY go.mod ./
# ソースコードをコピー
COPY . .

# go.sumをコンテナ内で生成する
RUN go mod tidy

# ★★★ 修正点 ★★★
# 出力先をワーキングディレクトリ(/src)配下に変更する
RUN CGO_ENABLED=0 GOOS=linux go build -o /src/main ./cmd/api

# --- 実行ステージ ---
FROM alpine:latest
WORKDIR /app

COPY --from=builder /src/main .
COPY config.yaml .

# ソースも参照したいなら（任意）
COPY --from=builder /src /src

# アプリケーションを実行
EXPOSE 8080
CMD ["./main"]
