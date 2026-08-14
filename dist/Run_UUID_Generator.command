#!/bin/bash
# 將此檔、uuid_generator_mac（或 uuid_generator_mac_intel）、game.txt 放在同一資料夾，雙擊即可執行。
DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$DIR" || exit 1

# 與手動執行 chmod +x 相同：確保本腳本與執行檔有可執行位元（複製 / 解壓後常會被清掉）
chmod +x "$0" 2>/dev/null || true

ARCH="$(uname -m)"
if [[ "$ARCH" == "arm64" ]]; then
	PREFERRED="./uuid_generator_mac"
	FALLBACK="./uuid_generator_mac_intel"
else
	PREFERRED="./uuid_generator_mac_intel"
	FALLBACK="./uuid_generator_mac"
fi

BIN=""
if [[ -f "$PREFERRED" ]]; then
	BIN="$PREFERRED"
elif [[ -f "$FALLBACK" ]]; then
	BIN="$FALLBACK"
else
	echo "找不到 uuid_generator_mac 或 uuid_generator_mac_intel。"
	echo "請把執行檔放在同一資料夾：${DIR}"
	read -r -p "按 Enter 關閉…"
	exit 1
fi

chmod +x "$BIN" 2>/dev/null || true

if [[ ! -f "game.txt" ]]; then
	echo "找不到 game.txt，請放在同一資料夾：${DIR}"
	read -r -p "按 Enter 關閉…"
	exit 1
fi

"$BIN"
EXITCODE=$?
echo
read -r -p "結束（離開碼 ${EXITCODE}）。按 Enter 關閉視窗…"
exit "$EXITCODE"
