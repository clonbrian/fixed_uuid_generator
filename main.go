package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var fixedUUIDs = []string{
	"2284:140",
	"2284:194",
	"2284:275",
	"2284:146",
	"2284:203",
	"2284:281",
	"2284:152",
	"2284:212",
	"2284:287",
	"2284:158",
	"2284:221",
	"2284:293",
	"2284:164",
	"2284:230",
	"2284:299",
	"2284:170",
	"2284:239",
	"2284:305",
	"2284:176",
	"2284:248",
	"2284:311",
	"2284:182",
	"2284:257",
	"2284:317",
	"2284:188",
	"2284:266",
	"2284:323",
	"2284:142",
	"2284:197",
	"2284:277",
	"2284:148",
	"2284:206",
	"2284:283",
	"2284:154",
	"2284:215",
	"2284:289",
	"2284:160",
	"2284:224",
	"2284:295",
	"2284:166",
	"2284:233",
	"2284:301",
	"2284:172",
	"2284:242",
	"2284:307",
	"2284:178",
	"2284:251",
	"2284:313",
	"2284:184",
	"2284:260",
	"2284:319",
	"2284:190",
	"2284:269",
	"2284:325",
	"2284:144",
	"2284:200",
	"2284:279",
	"2284:150",
	"2284:209",
	"2284:285",
	"2284:156",
	"2284:218",
	"2284:291",
	"2284:162",
	"2284:227",
	"2284:297",
	"2284:168",
	"2284:236",
	"2284:303",
	"2284:174",
	"2284:245",
	"2284:309",
	"2284:180",
	"2284:254",
	"2284:315",
	"2284:186",
	"2284:263",
	"2284:321",
	"2284:192",
	"2284:272",
	"2284:327",
	"2284:330",
	"2284:348",
	"2284:375",
	"2284:332",
	"2284:351",
	"2284:377",
	"2284:334",
	"2284:354",
	"2284:379",
	"2447:2",
	"2447:28",
	"2447:67",
	"2447:6",
	"2447:34",
	"2447:71",
	"2447:10",
	"2447:40",
	"2447:75",
	"2447:14",
	"2447:46",
	"2447:79",
	"2447:18",
	"2447:52",
	"2447:83",
	"2447:20",
	"2447:55",
	"2447:85",
	"2447:22",
	"2447:58",
	"2447:87",
	"2447:24",
	"2447:61",
	"2447:89",
	"2447:26",
	"2447:64",
	"2447:91",
	"2447:4",
	"2447:31",
	"2447:69",
	"2447:8",
	"2447:37",
	"2447:73",
	"2447:12",
	"2447:43",
	"2447:77",
	"2447:16",
	"2447:49",
	"2447:81",
	"2447:101",
	"2447:133",
	"2447:171",
	"2447:105",
	"2447:139",
	"2447:175",
	"2447:109",
	"2447:145",
	"2447:179",
	"2447:113",
	"2447:151",
	"2447:183",
	"2447:117",
	"2447:157",
	"2447:187",
	"2447:93",
	"2447:121",
	"2447:163",
	"2447:95",
	"2447:124",
	"2447:165",
	"2447:97",
	"2447:127",
	"2447:167",
	"2447:99",
	"2447:130",
	"2447:169",
	"2447:103",
	"2447:136",
	"2447:173",
	"2447:107",
	"2447:142",
	"2447:177",
	"2447:111",
	"2447:148",
	"2447:181",
	"2447:115",
	"2447:154",
	"2447:185",
	"2447:119",
	"2447:160",
	"2447:189",
	"2447:191",
	"2447:197",
	"2447:206",
	"2447:193",
	"2447:200",
	"2447:208",
	"2447:195",
	"2447:203",
	"2447:210",
	"2447:213",
	"2447:239",
	"2447:278",
	"2447:217",
	"2447:245",
	"2447:282",
	"2447:221",
	"2447:251",
	"2447:286",
	"2447:225",
	"2447:257",
	"2447:290",
	"2447:229",
	"2447:263",
	"2447:294",
	"2447:231",
	"2447:266",
	"2447:296",
	"2447:233",
	"2447:269",
	"2447:298",
	"2447:235",
	"2447:272",
	"2447:300",
	"2447:237",
	"2447:275",
	"2447:302",
	"2447:215",
	"2447:242",
	"2447:280",
	"2447:219",
	"2447:248",
	"2447:284",
	"2447:223",
	"2447:254",
	"2447:288",
	"2447:227",
	"2447:260",
	"2447:292",
	"2447:312",
	"2447:344",
	"2447:382",
	"2447:316",
	"2447:350",
	"2447:386",
	"2447:320",
	"2447:356",
	"2447:390",
	"2447:324",
	"2447:362",
	"2447:394",
	"2447:328",
	"2447:368",
	"2447:398",
	"2447:304",
	"2447:332",
	"2447:374",
	"2447:306",
	"2447:335",
	"2447:376",
	"2447:308",
	"2447:338",
	"2447:378",
	"2447:310",
	"2447:341",
	"2447:380",
	"2447:314",
	"2447:347",
	"2447:384",
	"2447:318",
	"2447:353",
	"2447:388",
	"2447:322",
	"2447:359",
	"2447:392",
	"2447:326",
	"2447:365",
	"2447:396",
	"2447:330",
	"2447:371",
	"2447:400",
	"2447:402",
	"2447:408",
	"2447:417",
	"2447:404",
	"2447:411",
	"2447:419",
	"2447:406",
	"2447:414",
	"2447:421",
	"2447:424",
	"2447:438",
	"2447:459",
	"2447:426",
	"2447:441",
	"2447:461",
	"2447:428",
	"2447:444",
	"2447:463",
	"2447:430",
	"2447:447",
	"2447:465",
	"2447:432",
	"2447:450",
	"2447:467",
	"2447:434",
	"2447:453",
	"2447:469",
	"2447:436",
	"2447:456",
	"2447:471",
	"2447:473",
	"2447:479",
	"2447:488",
	"2447:475",
	"2447:482",
	"2447:490",
	"2447:477",
	"2447:485",
	"2447:492",
}

type Game struct {
	Platform string
	GameKey  string
}

func main() {
	gamesPath := flag.String("games", "game.txt", "Path to game.txt (tab-separated with headers Platform and GameKey)")
	outPath := flag.String("out", "uuid_new.txt", "Output txt path")
	flag.Parse()

	games, err := readGames(*gamesPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "錯誤：%v\n", err)
		os.Exit(1)
	}

	expanded := expandGames(games)
	if len(expanded) > len(fixedUUIDs) {
		fmt.Fprintf(os.Stderr, "錯誤：產生的資料列過多（%d 筆），程式內建 UUID 僅支援 %d 筆\n", len(expanded), len(fixedUUIDs))
		os.Exit(1)
	}

	if err := writeOutput(*outPath, fixedUUIDs[:len(expanded)], expanded); err != nil {
		fmt.Fprintf(os.Stderr, "錯誤：寫入輸出檔失敗：%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Done. games=%d, rows=%d, output=%s\n", len(games), len(expanded), *outPath)
}

func readGames(path string) ([]Game, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("找不到 game.txt，請確認檔案與執行檔放在同一資料夾")
		}
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNo := 0
	var games []Game

	for scanner.Scan() {
		lineNo++
		raw := strings.TrimSpace(scanner.Text())
		if raw == "" {
			continue
		}

		fields := splitFields(raw)
		if lineNo == 1 {
			if len(fields) < 2 {
				return nil, fmt.Errorf("game.txt 格式錯誤：第一行標題至少要有兩欄（Platform、GameKey），目前：%q", raw)
			}
			continue
		}

		if len(fields) < 2 {
			return nil, fmt.Errorf("game.txt 第 %d 行格式錯誤，應為兩欄（平台與遊戲代碼）：%q", lineNo, raw)
		}

		platform := strings.TrimSpace(fields[0])
		gameKey := strings.TrimSpace(fields[1])
		if platform == "" || gameKey == "" {
			return nil, fmt.Errorf("game.txt 第 %d 行 Platform 或 GameKey 為空：%q", lineNo, raw)
		}

		games = append(games, Game{Platform: platform, GameKey: gameKey})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(games) == 0 {
		return nil, errors.New("game.txt 內沒有任何遊戲資料")
	}

	maxGames := len(fixedUUIDs) / 3
	if len(games) > maxGames {
		return nil, fmt.Errorf("遊戲數量超過上限：目前 %d 款，上限為 100 款", len(games))
	}

	return games, nil
}

func splitFields(line string) []string {
	if strings.Contains(line, "	") {
		return strings.Split(line, "	")
	}
	return strings.Fields(line)
}

func expandGames(games []Game) []string {
	out := make([]string, 0, len(games)*3)
	for _, g := range games {
		base := fmt.Sprintf("%s/%s", g.Platform, g.GameKey)
		out = append(out, base)
		out = append(out, base+"-w")
		out = append(out, base+"-t")
	}
	return out
}

func writeOutput(path string, uuids []string, names []string) error {
	if len(uuids) != len(names) {
		return fmt.Errorf("uuid count (%d) != name count (%d)", len(uuids), len(names))
	}

	dir := filepath.Dir(path)
	if dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for i := range uuids {
		if _, err := fmt.Fprintf(w, "%s	%s\n", uuids[i], names[i]); err != nil {
			return err
		}
	}
	return w.Flush()
}
