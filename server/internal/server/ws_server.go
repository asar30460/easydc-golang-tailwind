package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

/*
 * 以下宣告之變數是為了WS處理不同使用者加入不同伺服器的設置
 * 舉例說明，以陣列表示某個使用者所加入的伺服器。
 * 使用者A:[0, 1]; 使用者B:[1, 2]; 使用者C:[1, 2]
 * 假設使用者A在伺服器1發送了訊息，使用者B及使用者C都應該要接收到。
 */
var (
	clients       = make(map[*Client]bool)         // 確定使用者連線狀態
	serverClients = make(map[int]map[*Client]bool) // 使用者加入到伺服器客戶端映射
	broadcast     = make(chan WsMessage, 1024)     // 廣播訊息通道，所有訊息都通過該通道先發送給伺服器端，伺服器端再來決定要發送給哪些使用者。
	mu            sync.Mutex                       // 互斥鎖
)

type Client struct {
	UserEmail string          // Log Websocket 連線使用者
	UserName  string          // 紀錄發送訊息使用者
	Conn      *websocket.Conn // The Conn type represents a WebSocket connection.
	Servers   []int           // 該使用者加入的伺服器清單
	Receive   chan WsMessage  // 每個使用者只需要有接收通道，而發送是發給伺服器，讓伺服器決定這筆訊息要讓哪些使用者接收
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/*
 * 在使用WebSocket協議時，無法直接通過標準的HTTP Cookie機制來傳遞cookie。
 * 但我們可以將參數作為URL部分，查詢參數或 WebSocket 協議的子協議來傳遞
 */
func (h *Handler) HandleWS(ctx *gin.Context) {
	/* 
	 * 這邊貪圖方便，正確的做法應該是接收使用者JWT，
	 * 接著，解析出使用者ID，並檢查是否存放在資料庫中
	 */
	userId := ctx.Query("userId") // Param for URL[:], Query for URL[?]
	intUserId, _ := strconv.Atoi(userId)

	res, err := h.service.WsGetClientInfo(ctx, intUserId)
	if err != nil {
		log.Fatalln(err)
	}

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer ws.Close()

	client := &Client{
		UserEmail: res.UserEmail,
		UserName:  res.UserName,
		Conn:      ws,
		Servers:   make([]int, 0),
		Receive:   make(chan WsMessage, 64),
	}

	log.Printf("A new client %s is connected WS...", client.UserEmail)

	mu.Lock()
	for server_id := range res.Servers {
		if _, ok := serverClients[server_id]; !ok {
			serverClients[server_id] = make(map[*Client]bool)
		}
		serverClients[server_id][client] = true
		client.Servers = append(client.Servers, server_id)
	}
	mu.Unlock()

	go writeMessages(client)

	for {
		_, msg, err := ws.ReadMessage() // ws.ReadMessage means receive message from client
		if err != nil {
			log.Printf("Error reading message from client %s: %v", client.UserEmail, err)
			mu.Lock()
			delete(clients, client)
			for server_id := range client.Servers {
				delete(serverClients[server_id], client)
			}
			mu.Unlock()
			log.Printf("User %s disconnected successfully:", client.UserEmail)
			break
		}

		var result map[string]string
		err = json.Unmarshal(msg, &result)
		if err != nil {
			log.Println(err)
			return
		}
		intServerID, err := strconv.Atoi(result["ServerID"])
		if err != nil {
			log.Println(err)
			return
		}
		intChannelID, err := strconv.Atoi(result["ChannelID"])
		if err != nil {
			log.Println(err)
			return
		}

		// 將訊息封裝成 JSON 結構
		message := WsMessage{
			UserID:    intUserId,
			UserEmail: res.UserEmail,
			UserName:  res.UserName,
			ServerID:  intServerID,
			ChannelID: intChannelID,
			Message:   result["Message"],
		}

		log.Printf("Message received from %s: %s", client.UserEmail, message.Message)

		// 建議開執行緒
		h.service.WsSendMessage(ctx, message)
		broadcast <- message
	}
}

func writeMessages(client *Client) {
	for {
		msg := <-client.Receive
		messageJSON, err := json.Marshal(msg)
		// log.Print("msgBytes:", messageJSON)
		if err != nil {
			log.Printf("marshal error: %v", err)
			continue
		}
		err = client.Conn.WriteMessage(websocket.TextMessage, messageJSON) // ws.WriteMessage means send message to client
		if err != nil {
			log.Printf("error: %v", err)
			mu.Lock()
			client.Conn.Close()
			delete(clients, client)
			mu.Unlock()
		}
		log.Printf("Message sent to client %s at server %v: %s", client.UserEmail, msg.ServerID, msg.Message)
	}
}

func (h *Handler) BroadcastMessages() {
	for {
		message := <-broadcast
		log.Printf("Broadcasted message from %s at server %v: %s", message.UserEmail, message.ServerID, message.Message)

		mu.Lock()
		if len(serverClients[message.ServerID]) > 0 {
			for client := range serverClients[message.ServerID] {
				select {
				case client.Receive <- message:
				default:
				}
			}
		}
		mu.Unlock()
	}
}
