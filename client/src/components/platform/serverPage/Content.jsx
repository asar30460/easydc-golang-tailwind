import { useState, useEffect, Fragment } from "react";
import { Button, Input } from "@material-tailwind/react";
import { API_URL } from "../../../constants";

const Content = ({ channelID }) => {
  const [historyMsg, setHistoryMsg] = useState([]);
  const [msg, setMsg] = useState("");
  const [loading, setLoading] = useState(true);
  const onChange = ({ target }) => setMsg(target.value);

  useEffect(() => {
    const fetchChannelData = () => {
      fetch(`${API_URL}/server/getHistoryMsgs`, {
        method: "POST",
        credentials: "include", // 確保cookie包含在內
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ channel_id: channelID }),
      }).then((res) => {
        // console.log(res);
        res.json().then((data) => {
          // console.log(data["history_msgs"]);
          setHistoryMsg(data["history_msgs"]);
        });
      });

      setLoading(false);
    };

    fetchChannelData();
  }, [channelID]);

  const msgList = () => {
    return !historyMsg ? (
      <div>Nothing here...</div>
    ) : (
      historyMsg.map((item) => {
        // console.log(item);
        return (
          <Fragment key={item["UserID"] + item["Time"]}>
            <div className="container min-w-60 pb-2">
              <div className="flex items-baseline">
                <div className=" text-lg text-green-300 pr-2">
                  {item["UserName"]}
                </div>
                <div className="text-xs">{item["Time"]}</div>
              </div>
              <div className=" text-base">{item["Message"]}</div>
            </div>
          </Fragment>
        );
      })
    );
  };

  return (
    <div className="grow flex flex-col bg-grey-3">
      <div className="box-content overflow-auto m-3">
        {loading ? <div>Loading...</div> : msgList()}
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
            console.log(msg);

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

export default Content;
