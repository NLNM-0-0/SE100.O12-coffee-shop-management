"use client";

import CreateIngredientDialog from "@/components/product-manage/create-ingredient";
import {
  StockTable,
  getFilterString,
} from "@/components/stock-manage/stock-table";
import { endPoint } from "@/constants";
import { useSearchParams } from "next/navigation";
import { useSWRConfig } from "swr";

const StockManage = () => {
  const { mutate } = useSWRConfig();
  const searchParams = useSearchParams();

  const { filters, stringToFilter } = getFilterString();
  const page = searchParams.get("page") ?? "1";

  const handleIngreAdded = async (ingreId: string) => {
    mutate(
      `${endPoint}/ingredients?page=${page ?? 1}&limit=10${
        stringToFilter ?? ""
      }`
    );
  };
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách tồn kho</h1>
        <div>
          <CreateIngredientDialog handleIngredientAdded={handleIngreAdded} />
        </div>
      </div>
      <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <StockTable />
      </div>
    </div>
  );
};

export default StockManage;
