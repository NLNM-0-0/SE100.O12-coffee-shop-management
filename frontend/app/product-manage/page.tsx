import { Metadata } from "next";
import { withAuth } from "@/lib/role/withAuth";
import TableLayout from "./table-layout";
export const metadata: Metadata = {
  title: "Mặt hàng",
};
function ProductManagement({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  return <TableLayout />;
}

export default withAuth(ProductManagement, ["ING_VIEW"]);
