import React from "react";

const Header = ({ serverName, channelName }) => {
  return (
    <div className="flex items-center h-10 bg-grey-2 rounded-tl-3xl shadow-md z-10 relative">
      <div className="mx-4 min-w-52 font-bold">{serverName}</div>
      <div className="grow flex h-10 items-center bg-grey-3 font-semibold">
        {`\xa0\xa0\xa0#\xa0${channelName}`}
      </div>
    </div>
  );
};

export default Header;
