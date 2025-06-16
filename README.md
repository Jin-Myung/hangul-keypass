# Hangul Keypass

기억하기 쉬운 한글 자판 기반 비밀번호 생성기

## 개요

쌍모임이나 쌍자음이 포함된 단어를 조합하여 대문자가 포함된 기억하기 쉬운 비밀번호를 생성합니다.

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

## 단어 소스

한국어 학습용 어휘 목록에서 기본 단어들을 추출하였습니다: https://www.korean.go.kr/front/etcData/etcDataView.do?mn_id=208&etc_seq=71&pageIndex=1
