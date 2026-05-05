# fixed_uuid_generator

這支工具把 `uuid.txt` 的第一欄 uuid/config **直接寫死在程式裡**。
你之後每次只要更新 `game.txt`，不用再提供 `uuid.txt`。

## 輸入格式
`game.txt` 必須是 tab 分隔，第一行 header：

```txt
Platform	GameKey
YB	Goldliner
JILI	MinesCricket
```

## 規則
- 每款遊戲自動展開成 3 筆：
  - `Platform/GameKey`
  - `Platform/GameKey-w`
  - `Platform/GameKey-t`
- 固定 uuid config 共 300 筆
- 所以最多一次 90 款遊戲

## 最簡單的跑法
把 `main.go` 和你的 `game.txt` 放同一層：

```bash
go run . -games game.txt -out uuid_new.txt
```

跑完會得到：

```txt
2284:140	YB/Goldliner
2284:194	YB/Goldliner-w
2284:275	YB/Goldliner-t
```

## 編譯成 exe

```bash
go build -o fixed_uuid_generator.exe .
./fixed_uuid_generator.exe -games game.txt -out uuid_new.txt
```

## 給 Cursor 的一句話
請看 `cursor_prompt.txt`
