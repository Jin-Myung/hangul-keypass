# Hangul Keypass

기억하기 쉬운 한글 자판 기반 비밀번호 생성기

## 개발 환경

- Go (WebAssembly)
- Static Web (HTML + JS)
- Local test server: Python or Go

## 실행 방법

### 1. Go → WASM 빌드

```bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
go run serve.go
```
