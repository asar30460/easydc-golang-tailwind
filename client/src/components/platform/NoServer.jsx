import React from "react";
import { Header } from "./serverPage";

const NoServer = () => {
  return (
    <div className="grow flex flex-col h-screen">
      <Header serverName="" channelName="很高興見到你 !" />
      <div className="grow flex bg-grey-2">
        <div className="flex-none min-w-56 mx-2 mt-5"></div>
        <div className="grow flex flex-col items-center bg-grey-3 p-3">
          <div className="text-md">
            你還尚未加入任何伺服器，點擊側邊欄搜尋關鍵字尋找或建立吧。
          </div>
        </div>
        <div className="flex-none min-w-56 mx-2 mt-5"></div>
      </div>
    </div>
  );
};

export default NoServer;
