import { useState, Fragment } from "react";
import { Button, Input } from "@material-tailwind/react";

const Content = ({ channelList, channel, speaker }) => {
  const historyMsg = channelList.find((item) => item.name === channel);
  const [msg, setMsg] = useState("");
  const onChange = ({ target }) => setMsg(target.value);

  return (
    <div className="grow flex flex-col bg-grey-3">
      <div className="box-content overflow-auto m-3">
        {historyMsg.record.map((item) => {
          return (
            <Fragment key={item.time + item.speaker}>
              <div className="container min-w-60 pb-2">
                <div className="flex items-baseline">
                  <div className=" text-lg text-green-300 pr-2">
                    {speaker.get(item.speaker)}
                  </div>
                  <div className="text-xs">{item.time}</div>
                </div>
                <div className=" text-base">{item.content}</div>
              </div>
            </Fragment>
          );
        })}
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
