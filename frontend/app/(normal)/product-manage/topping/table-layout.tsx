"use client";

import { ToppingTable } from "@/components/product-manage/topping-table";
import { endPoint } from "@/constants";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { useSWRConfig } from "swr";

export const getFilterString = () => {
  const searchParams = useSearchParams();

  const search = searchParams.get("search") ?? undefined;

  let filters = [{ type: "", value: "" }];
  filters.pop();
  if (search) {
    filters = filters.concat({ type: "search", value: search });
  }
  let stringToFilter = "";
  filters.forEach((item) => {
    stringToFilter = stringToFilter.concat(`&${item.type}=${item.value}`);
  });
  return { stringToFilter: stringToFilter, filters: filters };
};

const TableLayout = () => {
  const router = useRouter();
  const searchParams = useSearchParams();
  const { mutate } = useSWRConfig();

  const { filters, stringToFilter } = getFilterString();
  const page = searchParams.get("page") ?? "1";

  const { currentUser } = useCurrentUser();
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách topping</h1>
        {currentUser &&
        includesRoles({
          currentUser: currentUser,
          allowedFeatures: ["TOP_CREATE"],
        }) ? (
          <Link
            href="/product-manage/topping/add"
            className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2"
          >
            Thêm topping
          </Link>
        ) : null}
      </div>

      <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <ToppingTable />
      </div>
    </div>
  );
};

export default TableLayout;
