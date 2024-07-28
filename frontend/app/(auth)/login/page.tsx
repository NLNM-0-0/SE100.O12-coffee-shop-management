import Image from "next/image";
import LoginForm from "@/components/login/login-form";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Đăng nhập",
};
const Login = async () => {
  return (
    <div className="bg-auth-background bg-right bg-cover bg-no-repeat flex flex-1 flex-col p-8">
      <div className="flex flex-row align-middle">
        <Image
          src="/android-chrome-192x192.png"
          alt="logo"
          width={50}
          height={50}
        ></Image>
        <p className={`text-lg ml-2 font-semibold self-center`}>Coffee shop</p>
      </div>
      <div className="flex flex-1 flex-col justify-center">
        <div className="flex flex-row self-center gap-36">
          <Image
            src="/login-background.png"
            alt="login-background"
            width={500}
            height={500}
            className="lg:block hidden"
          />
          <LoginForm />
        </div>
      </div>
    </div>
  );
};

export default Login;
