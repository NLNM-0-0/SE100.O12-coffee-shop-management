"use client";

import Link from "next/link";
import Image from "next/image";
import Profile from "./profile";
const Header = () => {
  return (
    <div className=" flex z-10 bg-white w-[100%] border-b border-gray-200">
      <div className="flex flex-1 h-[47px] items-center justify-between px-4 pl-12">
        <div className="flex items-center space-x-4">
          <Link
            href="/"
            className="flex flex-row space-x-3 items-center justify-center md:hidden"
          >
            <div
              className={`flex align-middle justify-center items-center p-[4px] h-[40px] w-[40px]  rounded-xl`}
            >
              <Image
                src="/android-chrome-192x192.png"
                alt="logo"
                className="sidebar__logo"
                width={36}
                height={36}
              ></Image>
            </div>
            <span className="font-semibold text-lg flex ">Coffee Shop</span>
          </Link>
        </div>
        <div className="self-center flex">
          <Profile />
        </div>
      </div>
    </div>
  );
};

export default Header;
