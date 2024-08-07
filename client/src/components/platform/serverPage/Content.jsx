import { useState, useEffect, Fragment, memo } from "react";
import { Button, Input } from "@material-tailwind/react";
import { API_URL } from "../../../constants";

const Content = ({ serverID, channelID, ws }) => {
  const [historyMsg, setHistoryMsg] = useState([]);
  const [msg, setMsg] = useState("");
  const [loading, setLoading] = useState(true);
  const onChange = ({ target }) => setMsg(target.value);

  let timeInMs = Date.now();

  const fetchChannelData = () => {
    setLoading(true);

    fetch(`${API_URL}/server/getHistoryMsgs`, {
      method: "POST",
      credentials: "include", // 確保cookie包含在內
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ channel_id: channelID }),
    }).then((res) => {
      res.json().then((data) => {
        // console.log(data);
        setHistoryMsg(data["history_msgs"]);
      });
    });

    setLoading(false);
  };

  useEffect(() => {
    fetchChannelData();
  }, [channelID]);

  useEffect(() => {
    const handleMessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.ChannelID === parseInt(channelID)) {
        setHistoryMsg((prevHistoryMsg) => [
          ...prevHistoryMsg,
          {
            Message: data.Message,
            Time: timeInMs,
            UserID: data.UserID,
            UserName: data.UserName,
          },
        ]);
        // console.log("historyMsg", historyMsg);
      }
    };

    ws.addEventListener("message", handleMessage);
    return () => {
      ws.removeEventListener("message", handleMessage);
    };
  }, [ws, channelID, historyMsg]);

  return (
    <div className="grow flex flex-col bg-grey-3 h-[calc(100vh-40px)] overflow-auto">
      <div className="flex flex-col-reverse overflow-auto m-3">
        {loading ? <div>Loading...</div> : <MsgList messages={historyMsg} />}
      </div>
      <div className="flex relative mt-auto p-3">
        <Input
          placeholder="輸入訊息"
          labelProps={{
            className: "hidden",
          }}
          value={msg}
          onChange={onChange}
          containerProps={{
            className: "min-w-[100px]",
          }}
          className="!border !border-gray-800 bg-gray-700 text-gray-300 placeholder:text-gray-600 placeholder:opacity-100 "
        />
        <Button
          size="sm"
          color={msg ? "gray" : "black"}
          disabled={!msg}
          onClick={() => {
            ws.send(
              JSON.stringify({
                ServerID: serverID,
                ChannelID: channelID,
                Message: msg,
              })
            );
            // console.log(msg);

            // 發送完畢後清除輸入框
            setMsg("");
          }}
          className="!absolute right-4 top-4 rounded"
        >
          傳送
        </Button>
      </div>
    </div>
  );
};

const MsgList = memo(({ messages }) => {
  // console.log("rerender");

  return !messages ? (
    <p>Nothing here...</p>
  ) : (
    <div>
      {messages.map((msg) => (
        <Fragment key={msg.UserID + msg.Time}>
          <div className="container min-w-60 pb-2">
            <div className="flex items-baseline">
              <div className=" text-lg text-green-300 pr-2">
                {msg["UserName"]}
              </div>
              <div className="text-xs">{msg["Time"]}</div>
            </div>
            <div className=" text-base">{msg["Message"]}</div>
          </div>
        </Fragment>
      ))}
    </div>
  );
});

export default Content;
