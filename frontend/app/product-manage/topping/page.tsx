import { Metadata } from "next";
import { withAuth } from "@/lib/role/withAuth";
import TableLayout from "./table-layout";
export const metadata: Metadata = {
  title: "Topping",
};
function ToppingManagement({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  return <TableLayout />;
}

export default withAuth(ToppingManagement, ["TOP_VIEW"]);
