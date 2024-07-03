import { useState, useEffect } from "react";
import { Button, Input } from "@material-tailwind/react";

const SearchResult = ({ serverList }) => {
  const [category, setCategory] = useState("all");

  // 搜尋伺服器對話框參數設置
  const [searchKey, setSearchKey] = useState("");
  const [filteredServerList, setFilteredServerList] = useState(serverList);

  return (
    <div className="grow flex flex-col h-screen">
      <div className="grow flex bg-grey-2 rounded-tl-3xl">
        <CategoryBar
          setCategory={setCategory}
          serverList={serverList}
          setFilteredServerList={setFilteredServerList}
        />
        <Content
          serverList={serverList}
          searchKey={searchKey}
          setSearchKey={setSearchKey}
          filteredServerList={filteredServerList}
          setFilteredServerList={setFilteredServerList}
        />
      </div>
    </div>
  );
};

const CategoryBar = ({
  category,
  setCategory,
  serverList,
  setFilteredServerList,
}) => {
  return (
    <div className="flex flex-none flex-col min-w-56 mx-2 mt-5">
      <div className="text-xs mb-2 p-1">探索</div>
      <Button
        className={
          "flex items-center rounded-md bg-grey-2 gap-2 hover:bg-grey-1 px-2"
        }
        key={0}
        onClick={() => {
          setCategory(category);
          setFilteredServerList(serverList);
        }}
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="currentColor"
          className="size-6"
        >
          <path d="M5.566 4.657A4.505 4.505 0 0 1 6.75 4.5h10.5c.41 0 .806.055 1.183.157A3 3 0 0 0 15.75 3h-7.5a3 3 0 0 0-2.684 1.657ZM2.25 12a3 3 0 0 1 3-3h13.5a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3v-6ZM5.25 7.5c-.41 0-.806.055-1.184.157A3 3 0 0 1 6.75 6h10.5a3 3 0 0 1 2.683 1.657A4.505 4.505 0 0 0 18.75 7.5H5.25Z" />
        </svg>
        全部
      </Button>
    </div>
  );
};

const Content = ({
  searchKey,
  setSearchKey,
  serverList,
  filteredServerList,
  setFilteredServerList,
}) => {
  return (
    <div className="grow flex flex-col bg-grey-3 p-5">
      <div className="flex items-center w-full gap-2 mb-3">
        <Input
          type="text"
          label="關鍵字"
          labelProps={{
            className: "hidden",
          }}
          value={searchKey}
          onChange={(e) => {
            setSearchKey(e.target.value);
          }}
          className="!border !border-gray-900 !bg-grey-1 text-gray-300 placeholder:text-gray-600 placeholder:opacity-100 rounded-sm"
        />
        <Button
          variant="filled"
          className="flex h-10 bg-indigo-600 items-center rounded-md gap-2"
          disabled={!searchKey}
          onClick={() => {
            const result = serverList.find((item) => item.name === searchKey);
            setFilteredServerList(result ? [result] : []);
          }}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            className="size-6"
          >
            <path
              fillRule="evenodd"
              d="M10.5 3.75a6.75 6.75 0 1 0 0 13.5 6.75 6.75 0 0 0 0-13.5ZM2.25 10.5a8.25 8.25 0 1 1 14.59 5.28l4.69 4.69a.75.75 0 1 1-1.06 1.06l-4.69-4.69A8.25 8.25 0 0 1 2.25 10.5Z"
              clipRule="evenodd"
            />
          </svg>
        </Button>
      </div>
      <div className="flex flex-wrap gap-4">
        {filteredServerList.length !== 0 ? (
          filteredServerList.map((item) => (
            <ServerCard
              key={item.serverID}
              serverID={item.serverID}
              serverName={item.name}
              creator={item.users[0].name}
            />
          ))
        ) : (
          <div>無結果</div>
        )}
      </div>
    </div>
  );
};

const ServerCard = ({ serverID, serverName, creator }) => {
  return (
    <div className="box-content flex-col w-72 h-44 bg-blue-gray-300 rounded-sm">
      <div className="box-content flex flex-col items-center justify-center w-full h-28 bg-grey-1 rounded-t-sm">
        預設圖片 / 自訂圖片
      </div>
      <div className="box-content w-full h-16 flex bg-grey-2">
        <div className="box-content h-16 flex flex-col bg-grey-2">
          <div className="text-md font-bold p-2">{serverName}</div>
          <div className="text-sm px-2">{creator}</div>
        </div>
        <Button
          className="bg-indigo-500 w-12 m-3 py-1 px-3 rounded hover:bg-indigo-400 ml-auto"
          onClick={() => {
            console.log(`加入 ${serverName}, ID: ${serverID}`);
          }}
        >
          加入
        </Button>
      </div>
    </div>
  );
};

export default SearchResult;
