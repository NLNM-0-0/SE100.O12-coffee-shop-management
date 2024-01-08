import Dashboard from "@/components/dashboard/dashboard";
import { getUser } from "@/lib/auth/action";
import { withAuth } from "@/lib/role/withAuth";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Trang chủ",
};
const DashboardPage = () => {
  return <Dashboard />;
};

export default withAuth(DashboardPage, ["RPT_SALE"]);