import Loading from "@/components/loading";
import CreateDialog from "@/components/supplier-manage/create";
import TableLayout from "@/components/supplier-manage/table-layout";

import { Suspense } from "react";
import { Metadata } from "next";
import { withAuth } from "@/lib/role/withAuth";
export const metadata: Metadata = {
  title: "Quản lý nhà cung cấp",
};
function SupplierManage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  return (
    <div className="col">
      <div className="flex flex-row justify-between ">
        <h1>Danh sách nhà cung cấp</h1>
        <div className="flex gap-4">
          <CreateDialog />
        </div>
      </div>

      <div className="my-3 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <Suspense fallback={<Loading />}>
          <TableLayout searchParams={searchParams} />
        </Suspense>
      </div>
    </div>
  );
}

export default withAuth(SupplierManage, ["SUP_VIEW"]);
