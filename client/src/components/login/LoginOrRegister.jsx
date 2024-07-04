import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button, Input } from "@material-tailwind/react";
import { API_URL } from "../../../constants";

const LoginOrRegister = () => {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [resHint, setResHint] = useState("");
  // 如果偵測到使用者輸入未註冊的電子郵件，則顯示此欄位
  const [showUsername, setShowUsername] = useState(false);
  const [username, setUsername] = useState("");

  const submitHandler = async () => {
    try {
      if (showUsername) {
        await register();
      }

      const res = await fetch(`${API_URL}/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email: email, password: password }),
      });

      const data = await res.json();
      if (res.status === 200) {
        const user = {
          username: data.user_name,
          id: data.id,
        };

        localStorage.setItem("user_info", JSON.stringify(user));
        navigate("/server");
      } else {
        // This "error" column is defined in backend
        switch (data["error"]) {
          case "此電子郵件尚未註冊，現在進行註冊":
            setShowUsername(true);
            setResHint(`${data["error"]}。請輸入您的使用者名稱`);
            break;
          case "密碼錯誤":
            setResHint(`此電子郵件已註冊，但${data["error"]}。請重新輸入`);
            setShowUsername(false);
            setPassword("");
            break;
          default:
            console.log(data["error"]);
            break;
        }
      }
    } catch (error) {
      console.error("fetch error:", error);
    }
  };

  const register = async () => {
    try {
      const res = await fetch(`${API_URL}/createUser`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: email,
          user_name: username,
          Password: password,
        }),
      });
      if (res.status === 200) {
        console.log("register success");
      } else {
        console.log(res);
      }
    } catch (error) {
      console.error("fetch error while registering", error);
    }
  };

  return (
    <>
      <div className="fixed flex items-center p-5">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="currentColor"
          className="size-10 mr-2"
        >
          <path
            fillRule="evenodd"
            d="M4.848 2.771A49.144 49.144 0 0 1 12 2.25c2.43 0 4.817.178 7.152.52 1.978.292 3.348 2.024 3.348 3.97v6.02c0 1.946-1.37 3.678-3.348 3.97a48.901 48.901 0 0 1-3.476.383.39.39 0 0 0-.297.17l-2.755 4.133a.75.75 0 0 1-1.248 0l-2.755-4.133a.39.39 0 0 0-.297-.17 48.9 48.9 0 0 1-3.476-.384c-1.978-.29-3.348-2.024-3.348-3.97V6.741c0-1.946 1.37-3.68 3.348-3.97ZM6.75 8.25a.75.75 0 0 1 .75-.75h9a.75.75 0 0 1 0 1.5h-9a.75.75 0 0 1-.75-.75Zm.75 2.25a.75.75 0 0 0 0 1.5H12a.75.75 0 0 0 0-1.5H7.5Z"
            clipRule="evenodd"
          />
        </svg>
        <div className="text-xl font-extrabold">EasyDC</div>
      </div>

      <div className="flex justify-center items-center min-h-screen">
        <div className="box-content min-w-[480px] h-88 bg-grey-2 flex flex-col items-center rounded-xl p-8 gap-2">
          <div className="text-2xl font-bold">歡迎回來 !</div>

          <CustomInput
            title={"電子郵件*"}
            type={"email"}
            disabled={showUsername}
            val={email}
            setVal={setEmail}
          />
          <CustomInput
            title={"密碼*"}
            type={"password"}
            disabled={showUsername}
            val={password}
            setVal={setPassword}
          />
          {showUsername && (
            <CustomInput
              title={"使用者名稱（16位英文數字混合）*"}
              type={"text"}
              maxLength={16}
              val={username}
              setVal={setUsername}
            />
          )}

          <Button
            // disable按鈕情況:
            // 1. 無論showUsername為true或false，電子郵件或密碼為空
            // 2. 當showUsername為true時，增加一條規則:使用者名稱為空
            disabled={!email || !password || (showUsername && !username)}
            onClick={() => {
              submitHandler();
            }}
            className="w-full rounded-sm bg-indigo-600 text-md mt-7"
          >
            登入 / 註冊
          </Button>
          <div className="text-sm text-red-400">{resHint}</div>
          {showUsername && (
            <div
              className="text-sm underline cursor-pointer"
              onClick={() => {
                setResHint("");
                setUsername("");
                setShowUsername(false);
              }}
            >
              返回登入
            </div>
          )}
        </div>
      </div>
    </>
  );
};

const CustomInput = ({ title, type, disabled, maxLength, val, setVal }) => {
  return (
    <>
      <div className="text-xs mr-auto mt-5">{title}</div>
      <Input
        type={type}
        labelProps={{
          className: "hidden",
        }}
        disabled={disabled}
        maxLength={maxLength}
        value={val}
        onChange={(e) => {
          setVal(e.target.value);
        }}
        containerProps={{
          className: "min-w-[100px]",
        }}
        className="!border !border-gray-900 !bg-grey-1 text-gray-300 placeholder:text-gray-600 placeholder:opacity-100 rounded-sm"
      />
    </>
  );
};

export default LoginOrRegister;
