import React from "react";
import { Button } from "@material-tailwind/react";

const Channel = ({ server, channel, setChannel }) => {
  return (
    <div className="flex-none min-w-56 mx-2 mt-5">
      <div className="flex justify-between mb-2">
        <div className="text-xs">文字頻道</div>
        <div className="text-xs">+</div>
      </div>
      <div className="grid">
        {server.channels.map((item) => (
          <Button
            className={`flex items-center rounded-none ${
              item.name === channel ? "bg-gray-700" : "bg-grey-2"
            } gap-2 px-2`}
            key={item.name}
            onClick={() => {
              setChannel(item.name);
            }}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth="1.5"
              stroke="currentColor"
              className=" size-4"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M5.25 8.25h15m-16.5 7.5h15m-1.8-13.5-3.9 19.5m-2.1-19.5-3.9 19.5"
              />
            </svg>

            {item.name}
          </Button>
        ))}
      </div>

      {/* <div className="flex justify-between mt-8 mb-2">
          <div className="text-xs">語音頻道</div>
          <div className="text-xs">+</div>
        </div>
        <div className="grid gap-2">
          <Button>聊天室1</Button>
          <Button>聊天室2</Button>
        </div> */}
    </div>
  );
};

export default Channel;
