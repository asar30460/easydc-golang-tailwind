import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Button, Dialog, Input } from "@material-tailwind/react";
import { API_URL } from "../../constants";

const Sidebar = ({
  serverList,
  server,
  setServer,
  setCruding,
  setSwitchServer,
  ws,
}) => {
  const navigate = useNavigate();

  // 建立伺服器對話框參數設置
  const [openCreate, setOpenCreate] = useState(false);
  const [serverName, setServerName] = useState("");

  return (
    <div className="flex-none overflow-auto items-center h-screen w-16 m-0 pt-3 flex flex-col z-20 bg-grey-1 text-white gap-3">
      {Object.keys(serverList).map((key) => (
        <Link key={key} to={`${key}`}>
          <SideBarIcon
            key={key}
            alt={serverList[key]}
            text={serverList[key]}
            isSelected={key === server}
            onClick={() => {
              setSwitchServer(key);
              setServer(key);
            }}
          />
        </Link>
      ))}

      <div className="box-content w-1/2 bg-gray-700 h-0.5 rounded-full mt-1"></div>

      <SearchServer />

      <CreateServerDialog
        open={openCreate}
        setOpen={setOpenCreate}
        serverName={serverName}
        setServerName={setServerName}
        setCruding={setCruding}
      />
      <div
        className="h-10 w-12 relative flex items-center justify-center bg-grey-3 rounded-md group mt-auto mb-3
        text-sm font-bold cursor-pointer hover:bg-red-800 hover:text-white"
        onClick={async () => {
          localStorage.clear();
          await fetch(`${API_URL}/logout`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
          });
          ws.close();
          navigate("/");
        }}
      >
        登出
      </div>
    </div>
  );
};

const SideBarIcon = ({ alt, text, isSelected, onClick }) => (
  <div
    className={`${
      isSelected ? "rounded-2xl bg-blue-900" : "rounded-3xl"
    } sidebar-icon group`}
    onClick={onClick}
  >
    {alt}

    <span className="sidebar-tooltip group-hover:scale-100">{text}</span>
  </div>
);

// 搜尋伺服器
const SearchServer = () => {
  return (
    <Link to="/server/search">
      <div className="rounded-3xl sidebar-icon group text-green-500 hover:text-green-50 hover:bg-green-700">
        搜尋
      </div>
    </Link>
  );
};

// 建立伺服器對話窗
const CreateServerDialog = ({
  open,
  setOpen,
  serverName,
  setServerName,
  setCruding,
}) => {
  const navigate = useNavigate();
  return (
    <>
      <div
        className="rounded-3xl sidebar-icon group text-green-500 hover:text-green-50 hover:bg-green-700"
        onClick={() => {
          setOpen(true);
        }}
      >
        建立
      </div>
      <Dialog
        open={open}
        handler={() => {
          setOpen(!true);
        }}
        size="sm"
        className="flex flex-col items-center bg-grey-3 text-gray-200 gap-1 rounded-sm"
      >
        <div className="font-bold text-2xl pt-5">建立伺服器</div>
        <div className="box-content flex flex-col w-11/12 mt-5 gap-1">
          <div className="text-xs mr-auto ">伺服器名稱*</div>
          <Input
            labelProps={{
              className: "hidden",
            }}
            value={serverName}
            onChange={(e) => {
              setServerName(e.target.value);
            }}
            className="!border !border-gray-900 !bg-grey-1 text-gray-300 placeholder:text-gray-600 placeholder:opacity-100 rounded-sm"
          />
        </div>

        <div className="box-content w-full flex items-center justify-center bg-grey-2 mt-5 h-16">
          <div className="box-content flex w-11/12">
            <Button
              onClick={() => {
                setOpen(false);
                // 清空輸入框文字
                setServerName("");
              }}
              className="rounded-sm bg-transparent shadow-none mr-auto px-1 hover:shadow-none"
            >
              取消
            </Button>
            <Button
              disabled={!serverName}
              onClick={() => {
                fetch(`${API_URL}/server/createServer`, {
                  method: "POST",
                  credentials: "include", // 確保cookie包含在內
                  headers: {
                    "Content-Type": "application/json",
                  },
                  body: JSON.stringify({
                    server_name: serverName,
                  }),
                })
                  .then((res) => {
                    if (res.status === 200) {
                      setOpen(false);
                      // 建立或搜尋伺服器後清空輸入框
                      setServerName("");
                      return res.json();
                    }
                  })
                  .then((data) => {
                    setCruding(data["server_id"]);
                    navigate(`${data["server_id"]}`);
                  });
              }}
              className="rounded-sm bg-indigo-600 ml-auto"
            >
              <span>建立</span>
            </Button>
          </div>
        </div>
      </Dialog>
    </>
  );
};

export default Sidebar;
