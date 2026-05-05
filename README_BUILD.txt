================================================================================
  UUID Generator - 建置與使用說明
================================================================================

一、如何在 Windows 上執行 build_all.bat
--------------------------------------------------------------------------------
1. 確認已安裝 Go（在命令提示字元輸入 go version 有版本號即可）。
2. 在檔案總管中雙擊 build_all.bat，
   或在「命令提示字元 (cmd)」裡切到專案資料夾後輸入：
      build_all.bat
3. 若成功，畫面上會顯示「Build OK.」以及三個輸出檔名。
4. 若失敗，會顯示「Build FAILED」並停在錯誤處，請依訊息排除問題。


二、Build 完後三個檔案各自給誰用
--------------------------------------------------------------------------------
- dist\uuid_generator.exe       → 給「Windows」電腦用。
- dist\uuid_generator_mac       → 給「Mac Apple Silicon（M1/M2/M3 等）」用。
- dist\uuid_generator_mac_intel → 給「Mac Intel」用。


三、Mac 使用者怎麼執行（Apple Silicon）
--------------------------------------------------------------------------------
1. 把 dist\uuid_generator_mac 複製到你要用的資料夾（例如和 game.txt 同一層）。
2. 開啟「終端機 (Terminal)」，切到該資料夾，例如：
      cd /Users/你的帳號/Desktop/我的資料夾
3. 第一次使用要先給執行權限：
      chmod +x uuid_generator_mac
4. 執行：
      ./uuid_generator_mac
5. 同資料夾下要有 game.txt，執行完會產生 uuid_new.txt。


四、Mac Intel 使用者怎麼執行
--------------------------------------------------------------------------------
1. 把 dist\uuid_generator_mac_intel 複製到你要用的資料夾（和 game.txt 同一層）。
2. 開啟「終端機 (Terminal)」，切到該資料夾。
3. 第一次使用要先給執行權限：
      chmod +x uuid_generator_mac_intel
4. 執行：
      ./uuid_generator_mac_intel
5. 同資料夾下要有 game.txt，執行完會產生 uuid_new.txt。


五、重要：執行檔要和 game.txt 放在同一個資料夾
--------------------------------------------------------------------------------
- Windows：請把 uuid_generator.exe 和 game.txt 放在同一個資料夾。
- Mac：在終端機裡 cd 到「放 game.txt 的那個資料夾」再執行 ./uuid_generator_mac 或 ./uuid_generator_mac_intel。
這樣程式才會正確讀到 game.txt，並在「同一資料夾」產生 uuid_new.txt。


六、Windows 使用方式
--------------------------------------------------------------------------------
- 雙擊 uuid_generator.exe 就會執行。
- 程式會自動讀取「同資料夾」的 game.txt，並在同一資料夾產生 uuid_new.txt。
- 若沒有 game.txt 或格式錯誤，會跳出錯誤訊息視窗。


七、Mac 版本需要從 Terminal 執行
--------------------------------------------------------------------------------
- Mac 的 uuid_generator_mac / uuid_generator_mac_intel 沒有視窗介面，
  必須從「終端機 (Terminal)」用指令執行（見上面第三、四點）。
- 執行成功時會在終端機顯示：Done. games=... output=uuid_new.txt

================================================================================
