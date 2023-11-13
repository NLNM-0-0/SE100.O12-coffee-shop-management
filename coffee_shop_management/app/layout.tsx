import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Sidebar from "@/components/sidebar";
import Header from "@/components/header";
import HeaderMobile from "@/components/header-mobile";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={`${inter.className} flex overflow-y-hidden`}>
        <Sidebar />
        <main className="flex-1 h-screen">
          <div className="flex flex-col">
            <Header />
            <HeaderMobile />
            <div className="md:p-10 p-4 h-[100vh] overflow-auto pt-[80px] md:pt-[80px]">
              {children}
            </div>
          </div>
        </main>
      </body>
    </html>
  );
}
