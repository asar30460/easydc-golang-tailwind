# EasyDC ─ Static Easy Discord Front─end (React Tailwind)

## 簡介

這個專案是模擬 Discord 使用畫面，資料來源是[靜態的 JSON 伺服器資料](./client/static_data/example.json)，為後續連接 Golang 後端所預先建立模板用途，已完成連接後端進度約50%。頁面元件（React components）如下:

    |───EasyDC
    |   |───登入
    |   |───後臺平台（Platform）
    |       |───SideBar
    |       |───新使用者引導（無伺服器）
    |       |───DC伺服器
    |           |───頻道列表
    |           |───聊天室
    |           └───成員列表
    |   |─── 搜尋伺服器
    |   |─── 伺服器類別
    |   └─── 搜尋結果
    └───────|─── 建立伺服器

本專案開發順序：靜態前端(React Tailwind) -> 規劃schema -> 連接後端
### 後端應用技術如下列舉：
1. service repository pattern(SRP)架構
2. JWT驗證
3. Gin context(JSON)、WebSocket、Goroutine


## 預覽畫面

<img alt="server" src="./client/previews/server.png">

| 登入                                         | 搜尋伺服器                                                 | 建立伺服器                                                 |
| -------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| <img alt="login" src="./client/previews/login.png"> | <img alt="searchServer" src="./client/previews/searchServer.png"> | <img alt="createServer" src="./client/previews/createServer.png"> |
