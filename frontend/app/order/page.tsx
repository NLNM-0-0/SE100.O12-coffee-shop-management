import { Metadata } from "next";
import { withAuth } from "@/lib/role/withAuth";
import OrderScreen from "./page-layout";
export const metadata: Metadata = {
  title: "Bán hàng",
};
export const Sale = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return <OrderScreen />;
};

export default withAuth(Sale, ["INV_CREATE"]);
