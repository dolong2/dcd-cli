#!/bin/bash
set -euo pipefail
trap 'rm -rf "${TEMP_DIR:-}"' EXIT

# 설정값
REPO_URL="https://github.com/dolong2/dcd-cli.git"
BINARY_NAME="dcd"
INSTALL_DIR="/usr/local/bin"

# 임시 작업 디렉토리
TEMP_DIR=$(mktemp -d)
# 루트 권한 확인
if [ "$EUID" -ne 0 ]; then
  echo "이 스크립트는 루트 권한이 필요합니다. sudo로 실행해주세요."
  exit 1
fi

echo "git 저장소에서 클론 중: $REPO_URL"
git clone "$REPO_URL" "$TEMP_DIR"

if [ $? -ne 0 ]; then
  echo "git clone 실패"
  exit 1
fi

cd "$TEMP_DIR" || exit 1

echo "Go 빌드 중..."
go build -ldflags="-X github.com/dolong2/dcd-cli/api.baseUrl=https://dcd-api.dolong2.co.kr -X github.com/dolong2/dcd-cli/websocket.baseUrl=wss://dcd-api.dolong2.co.kr" -o "$BINARY_NAME"

if [ ! -f "$BINARY_NAME" ]; then
  echo "빌드 실패"
  exit 1
fi

chmod +x "$BINARY_NAME"

echo "${INSTALL_DIR}에 설치 중..."
cp "$BINARY_NAME" "$INSTALL_DIR/"

echo "설치 완료: $(which $BINARY_NAME)"

# 임시 디렉토리 삭제
rm -rf "$TEMP_DIR"
