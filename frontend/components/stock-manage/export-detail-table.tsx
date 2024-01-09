"use client";

import {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { ExportNoteDetail } from "@/types";

import { useState } from "react";
import { Input } from "../ui/input";
import { toVND } from "@/lib/utils";
import { useRouter } from "next/navigation";
import Loading from "../loading";

export const columns: ColumnDef<ExportNoteDetail>[] = [
  {
    id: "stt",
    header: ({ table }) => (
      <div className="flex justify-center font-semibold">STT</div>
    ),
    cell: ({ row }) => (
      <div className="flex justify-center">{row.index + 1}</div>
    ),
    enableSorting: false,
    enableHiding: false,
    size: 4,
  },
  {
    accessorKey: "id",
    accessorFn: (row) => row.ingredient.id,
    header: () => {
      return <span className="font-semibold">Mã nguyên liệu</span>;
    },
    cell: ({ row }) => (
      <div className="leading-6">{row.original.ingredient.id}</div>
    ),
    size: 4,
  },
  {
    accessorKey: "name",
    header: () => {
      return <span className="font-semibold">Tên nguyên liệu</span>;
    },
    cell: ({ row }) => (
      <div className="leading-6">{row.original.ingredient.name}</div>
    ),
    size: 4,
  },
  {
    accessorKey: "amountImport",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Xuất</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {row.original.amountExport.toLocaleString("vi-VN")}
        </div>
      );
    },
    size: 4,
  },
  {
    accessorKey: "unitTypeName",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Đơn vị</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {row.original.unitTypeName}
        </div>
      );
    },
    size: 4,
  },
];
export function ExportDetailTable({
  details,
}: {
  details: ExportNoteDetail[];
}) {
  // const data = Object.values(details);
  console.log(details);
  const router = useRouter();
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});
  const table = useReactTable({
    data: details,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
    initialState: {
      pagination: {
        pageSize: 1000,
      },
    },
  });
  if (!details) {
    return <Loading />;
  } else
    return (
      <div className="rounded-md border overflow-x-auto min-w-full max-w-[10vw]">
        <Table className="w-max min-w-full">
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow
                key={headerGroup.id}
                className="bg-orange-50 hover:bg-orange-50"
              >
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext()
                          )}
                    </TableHead>
                  );
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row, index) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  Không tìm thấy bản ghi
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    );
}
