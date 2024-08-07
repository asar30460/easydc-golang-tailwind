/* 後臺頁面（組件：側邊攔 + 伺服器內容），根據使用者加入DC伺服器狀況:
 * 1. 未加入任何伺服器 -> 路由至引導頁面
 * 2. 加入一個或以上伺服器 -> 路由至第一個伺服器頁面
 */
import { Routes, Route, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { Sidebar, Server, NoServer, SearchResult } from "./";
import { API_URL, WS_URL } from "../../constants";

import Cookies from "universal-cookie";

const Platform = () => {
  // 列出該使用者有的伺服器清單，預設值是沒有參加任何伺服器
  const [serverList, setServerList] = useState([]);
  const [serverID, setServerID] = useState();
  const [ws, setWs] = useState(null);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();
  const cookie = new Cookies();

  // 檢測伺服器列表是否變動以重新載入，當發生變動時setCruding(server_id)
  const [cruding, setCruding] = useState(0);

  // 切換伺服器時，觸發更新伺服器component內容
  const [switchServer, setSwitchServer] = useState(0);

  useEffect(() => {
    const fetchAddedServer = () => {
      fetch(`${API_URL}/server/getServers`, {
        method: "GET",
        credentials: "include", // 確保cookie包含在內
        headers: {
          "Content-Type": "application/json",
        },
      }).then((res) => {
        // console.log(res);
        res.json().then((server_data) => {
          let data = server_data["servers"];
          // console.log(Object.keys(data).length);
          if (Object.keys(data).length !== 0) {
            setServerList(data);
            // console.log(Object.keys(data)[0]);
            setServerID(Object.keys(data)[0]);
          } else {
            setServerList([]);
          }
        });
      });

      const socket = new WebSocket(
        `${WS_URL}/server/handleWs?userId=${cookie.get("user_id")}`
      );
      setWs(socket);
      console.log("connect to websocket", socket);

      setLoading(false);
    };

    fetchAddedServer();
  }, [cruding]);

  // 如果該使用者存在任何伺服器，則起始畫面直接導到列表中的第一個伺服器
  useEffect(() => {
    if (!loading && serverList.length !== 0) {
      navigate(`${serverID}`);
    }
  }, [loading, serverList]);

  const renderComponet = () => {
    return (
      <Routes>
        <Route
          path="/search"
          element={<SearchResult serverList={serverList} />}
        />
        // 有無加入Server的結果判斷
        {serverList.length === 0 ? (
          <Route path="/" element={<NoServer />}></Route>
        ) : (
          Object.keys(serverList).map((key) => (
            <Route
              key={key}
              path={`/${key}`}
              element={
                <Server
                  serverID={serverID}
                  serverList={serverList}
                  switchServer={switchServer}
                  ws={ws}
                />
              }
            />
          ))
        )}
      </Routes>
    );
  };

  return (
    <div className="flex">
      <Sidebar
        serverList={serverList}
        server={serverID}
        setServer={setServerID}
        setCruding={setCruding}
        setSwitchServer={setSwitchServer}
        ws={ws}
      />
      {loading ? <div>loading...</div> : renderComponet()}
    </div>
  );
};

export default Platform;
