// 該「Server」指DC伺服器
import { useState, useEffect } from "react";

import { Header, Channel, Content, Member } from "./serverPage";
import { API_URL } from "../../constants";

const Server = ({ serverID, serverList, switchServer, ws }) => {
  const [channelList, setChannelList] = useState([]);
  const [channelID, setChannelID] = useState();
  const [member, serMember] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchServerData = () => {
      setLoading(true);

      fetch(`${API_URL}/server/${serverID}/getChannels`, {
        method: "GET",
        credentials: "include", // 確保cookie包含在內
        headers: {
          "Content-Type": "application/json",
        },
      }).then((res) => {
        // console.log(res);
        res.json().then((channel_data) => {
          let data = channel_data["channels"];
          if (Object.keys(data).length !== 0) {
            // console.log(data);
            setChannelList(data);
            setChannelID(Object.keys(data)[0]);
          } else {
            setChannelList([]);
          }
        });
      });

      fetch(`${API_URL}/server/${serverID}/getMembers`, {
        method: "GET",
        credentials: "include", // 確保cookie包含在內
        headers: {
          "Content-Type": "application/json",
        },
      }).then((res) => {
        res.json().then((member_data) => {
          // console.log(member_data["members"]);
          serMember(member_data["members"]);
        });
      });
      setLoading(false);
    };

    fetchServerData();
  }, [switchServer]);

  return loading ? (
    <div>Loading...</div>
  ) : (
    <div className="grow flex flex-col h-screen">
      <Header
        serverName={serverList[serverID]}
        channelName={channelList[channelID]}
      />
      <div className="grow flex bg-grey-2">
        <Channel
          channelList={channelList}
          channelID={channelID}
          setChannelID={setChannelID}
        />
        <Content serverID={serverID} channelID={channelID} ws={ws} />
        <Member member={member} />
      </div>
    </div>
  );
};

export default Server;
