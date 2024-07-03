import React from "react";
import { Button } from "@material-tailwind/react";

const Member = ({ member }) => {
  return (
    <div className="flex-none min-w-56 mx-2 mt-5">
      <div className="grid">
        {Array.from(member).map(([email, name]) => (
          <div key={email} className="flex gap-4 ">
            <Button
              variant="text"
              color="white"
              className="flex items-center rounded-sm w-full gap-2"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth="1.5"
                stroke="currentColor"
                className="size-6"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z"
                />
              </svg>

              {name}
            </Button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Member;
