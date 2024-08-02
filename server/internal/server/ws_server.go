package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn // The Conn type represents a WebSocket connection.
	Servers []int
	Message chan CreateMsgReq // 訊息通道，向客戶端發送訊息
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/*
 * 以下宣告之變數是為了WS處理不同使用者加入不同伺服器的設置
 * 舉例說明，以陣列表示某個使用者所加入的伺服器。
 * 使用者A:[0, 1]; 使用者B:[1, 2]; 使用者C:[1, 2]
 * 假設使用者A在伺服器1發送了訊息，使用者B及使用者C都應該要接收到。
 */
var (
	clients       = make(map[*Client]bool)   // 確定使用者連線狀態
	serverClients = make(map[int][]*Client)  // 使用者加入到伺服器客戶端映射
	broadcast     = make(chan *CreateMsgReq) // 廣播訊息通道，所有訊息都通過該通道來廣播給客戶端。
	mu            sync.Mutex                 // 互斥鎖
	once          sync.Once                  // broadcastMessages 只執行一次，所有使用者都靠這條
)

func (h *Handler) HandleWS(ctx *gin.Context) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer ws.Close()

	mu.Lock()
	res, err := h.service.GetServerByEmail(ctx.Request.Context(), ctx)
	if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	var servers []int
	for server_id := range res.Servers {
		servers = append(servers, server_id)
	}

	client := &Client{
		Conn:    ws,
		Servers: servers,
		Message: make(chan CreateMsgReq),
	}

	clients[client] = true
	for _, server := range servers {
		serverClients[server] = append(serverClients[server], client)
	}
	mu.Unlock()

	go writeMessages(client)

	// 確保 broadcastMessages 只被啟動一次
	once.Do(func() {
		go broadcastMessages()
	})

	for {
		log.Printf("Waiting for message from client...")
		_, msg, err := ws.ReadMessage()	// ws.ReadMessage means receive message
		if err != nil {
			log.Printf("error: %v", err)
			mu.Lock()
			delete(clients, client)
			for _, server := range servers {
				serverClients[server] = removeClient(serverClients[server], client)
			}
			mu.Unlock()
			break
		}

		var message CreateMsgReq
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("unmarshal:", err)
			continue
		}

		log.Printf("Broadcasting message: %+v\n", message)
		broadcast <- &message
	}
}

func writeMessages(client *Client) {
	for msg := range client.Message {
		msgBytes, err := json.Marshal(msg)
		log.Print("msgBytes:", msgBytes)
		if err != nil {
			log.Printf("marshal error: %v", err)
			continue
		}
		err = client.Conn.WriteMessage(websocket.TextMessage, msgBytes)	// ws.WriteMessage means send message
		if err != nil {
			log.Printf("error: %v", err)
			client.Conn.Close()
			mu.Lock()
			delete(clients, client)
			mu.Unlock()
		}
		log.Printf("Message sent to client: %+v\n", msg)
	}
}

func broadcastMessages() {
	for {
		message := <-broadcast
		log.Printf("Broadcasting message to channel %d: %+v\n", message.ChannelID, message)

		mu.Lock()
		if clients, ok := serverClients[message.ChannelID]; ok {
			for _, client := range clients {
				select {
				case client.Message <- *message:
				default:
				}
			}
		}
		mu.Unlock()
	}
}

func removeClient(clients []*Client, client *Client) []*Client {
	for i, c := range clients {
		if c == client {
			return append(clients[:i], clients[i+1:]...)
		}
	}
	return clients
}
