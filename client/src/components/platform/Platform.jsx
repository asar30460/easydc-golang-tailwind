/* 後臺頁面，根據使用者加入DC伺服器狀況:
 * 1. 未加入任何伺服器 -> 路由至引導頁面
 * 2. 加入一個或以上伺服器 -> 路由至第一個伺服器頁面
 */
import { Routes, Route, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { Sidebar, Server, NoServer, SearchResult } from "./";
import { server_data } from "..";

const Platform = () => {
  // 列出該使用者有的伺服器清單，預設值是沒有參加任何伺服器
  const [serverList, setServerList] = useState([]);
  const [serverID, setServerID] = useState();
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchAddedServer = () => {
      if (server_data.length !== 0) {
        setServerList(server_data);
        setServerID(1);
      }

      setLoading(false);
    };

    fetchAddedServer();
  }, [serverList]);

  // 如果該使用者存在任何伺服器，則起始畫面直接導到伺服器1
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
          serverList.map((item) => (
            <Route
              key={item.serverID}
              path={`/${item.serverID}`}
              element={<Server serverID={serverID} />}
            />
          ))
        )}
      </Routes>
    );
  };

  return (
    <div className="flex">
      <Sidebar server={serverID} setServer={setServerID} />
      {loading ? <div>loading...</div> : renderComponet()}
    </div>
  );
};

export default Platform;
