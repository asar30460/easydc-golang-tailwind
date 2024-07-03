// 該「Server」指DC伺服器
import { useState, useEffect } from "react";

import { Header, Channel, Content, Member } from "./serverPage";
import { server_data } from "..";

const Server = ({ serverID }) => {
  const [server, setServer] = useState(null);
  const [channel, setChannel] = useState("");
  const [member, serMember] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchServerData = () => {
      const serverData = server_data.find((item) => item.serverID === serverID);
      setServer(serverData);
      setChannel(serverData.channels[0].name);

      const memberMap = new Map();
      serverData.users.forEach((element) => {
        // console.log(memberMap.get(element.email));
        memberMap.set(element.email, element.name);
      });
      serMember(memberMap);

      setLoading(false);
    };

    fetchServerData();
  }, [serverID]);

  return loading ? (
    <div>Loading...</div>
  ) : (
    <div className="grow flex flex-col h-screen">
      <Header serverName={server.name} channelName={channel} />
      <div className="grow flex bg-grey-2">
        <Channel server={server} channel={channel} setChannel={setChannel} />
        <Content
          channelList={server.channels}
          channel={channel}
          speaker={member}
        />
        <Member member={member} />
      </div>
    </div>
  );
};

export default Server;
