import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button, Input } from "@material-tailwind/react";
import { API_URL } from "../../../constants";

const Login = () => {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [resHint, setResHint] = useState("");

  const submitHandler = async () => {
    try {
      const res = await fetch(`${API_URL}/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
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
        setResHint(data["error"]);
      }
    } catch (error) {
      console.error("fetch error:", error);
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
            val={email}
            setVal={setEmail}
          />
          <CustomInput
            title={"密碼*"}
            type={"password"}
            val={password}
            setVal={setPassword}
          />
          <Button
            disabled={!email || !password}
            onClick={() => {
              submitHandler();
            }}
            className="w-full rounded-sm bg-indigo-600 text-md mt-7"
          >
            登入 / 註冊
          </Button>
          <div className="text-sm text-red-400">{resHint}</div>
        </div>
      </div>
    </>
  );
};

const CustomInput = ({ title, type, val, setVal }) => {
  return (
    <>
      <div className="text-xs mr-auto mt-5">{title}</div>
      <Input
        type={type}
        labelProps={{
          className: "hidden",
        }}
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

export default Login;
