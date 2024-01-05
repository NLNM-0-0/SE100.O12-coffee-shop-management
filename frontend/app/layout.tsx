import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Sidebar from "@/components/sidebar";
import Header from "@/components/header";
import HeaderMobile from "@/components/header-mobile";
import { Toaster } from "@/components/ui/toaster";
import { auth } from "@/lib/auth/auth";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Coffee Shop",
  description: "Coffee shop management app",
};

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}
) {
  const isAuthented = await auth().then(res => res?.user)
  return (
    <html lang="en" className="h-full">
      <body className={`${inter.className} flex overflow-y-hidden h-full`}>
        {isAuthented ?
          <>
            <Sidebar />
            <main className="flex flex-1">
              <div className="flex w-full flex-col overflow-y-hidden">
                <Header />
                <HeaderMobile />
                <div className="md:p-10 p-4 overflow-auto">{children}</div>
                <Toaster />
              </div>
            </main>
          </>
          :
          <main className="flex flex-1">
            {children}
            <Toaster />
          </main>
        }
      </body>
    </html>
  );
}
