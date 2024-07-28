import { Metadata } from "next";
import TableLayout from "./table-layout";
import { withAuth } from "@/lib/role/withAuth";
export const metadata: Metadata = {
  title: "Danh mục",
};
const CategoryPage = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return <TableLayout searchParams={searchParams} />;
};

export default withAuth(CategoryPage, ["CAT_VIEW"]);
