"use client";
import { CategoryTable } from "@/components/product-manage/category-table";
import CreateCategory from "@/components/product-manage/create-category";
import ImportSheet from "@/components/product-manage/import-sheet";
import Loading from "@/components/loading";
import { Button } from "@/components/ui/button";
import { toast } from "@/components/ui/use-toast";
import { endPoint } from "@/constants";
import { useCurrentUser } from "@/hooks/use-user";
import { getUser } from "@/lib/auth/action";
import createListCategory from "@/lib/category/createListCategory";
import { includesRoles } from "@/lib/utils";
import { useRouter } from "next/navigation";
import { Suspense, useEffect, useState } from "react";
import { useSWRConfig } from "swr";

const TableLayout = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const { mutate } = useSWRConfig();

  const router = useRouter();
  const page = searchParams["page"] ?? "1";

  const handleCategoryAdded = (name: string) => {
    mutate(`${endPoint}/categories?page=${page ?? 1}&limit=10`);
  };
  const handleFile = async (reader: FileReader) => {
    let categories: string[] = [];
    const ExcelJS = require("exceljs");
    const wb = new ExcelJS.Workbook();
    reader.onload = () => {
      const buffer = reader.result;
      wb.xlsx.load(buffer).then((workbook: any) => {
        workbook.eachSheet((sheet: any, id: any) => {
          sheet.eachRow((row: any, rowIndex: number) => {
            if (rowIndex > 1) {
              const name = row.getCell(1).value.toString();
              if (name && name != "") {
                categories.push(name);
              }
            }
          });
        });
      });
    };

    const response: Promise<any> = createListCategory({ names: categories });
    const responseData = await response;
    if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "success",
        title: "Thành công",
        description: "Thêm danh mục thành công",
      });
      handleCategoryAdded("");
    }
  };
  const { currentUser } = useCurrentUser();
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh mục</h1>
        {currentUser &&
        includesRoles({
          currentUser: currentUser,
          allowedFeatures: ["CAT_CREATE"],
        }) ? (
          <div className="flex gap-4">
            <CreateCategory handleCategoryAdded={handleCategoryAdded}>
              <Button>Thêm danh mục</Button>
            </CreateCategory>
          </div>
        ) : null}
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <Suspense fallback={<Loading />}>
          <CategoryTable
            searchParams={searchParams}
            currentUser={currentUser}
          />
        </Suspense>
      </div>
    </div>
  );
};

export default TableLayout;
