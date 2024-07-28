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

import { useState } from "react";
import { toVND } from "@/lib/utils";
import { useRouter } from "next/navigation";
import { InvoiceDetailProps } from "@/app/(normal)/invoice/[invoiceId]/page";

export const columns: ColumnDef<InvoiceDetailProps>[] = [
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
    accessorKey: "name",
    header: () => {
      return <span className="font-semibold">Tên món</span>;
    },
    cell: ({ row }) => (
      <div className="leading-6 flex flex-col">
        {row.original.food.name} ({row.original.sizeName})
        {row.original.toppings && (
          <span className="font-light">
            Topping: {row.original.toppings.map((item) => item.name).join(", ")}
          </span>
        )}
        {row.original.description && row.original.description !== "" && (
          <span className="font-light">({row.original.description})</span>
        )}
      </div>
    ),
    size: 4,
  },
  {
    accessorKey: "amount",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Số lượng</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {row.original.amount.toLocaleString("vi-VN")}
        </div>
      );
    },
    size: 4,
  },
  {
    accessorKey: "unitPrice",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Đơn giá</span>
      </div>
    ),
    cell: ({ row }) => {
      const amount = row.original.unitPrice;

      return <div className="text-right font-medium">{toVND(amount)}</div>;
    },
    size: 4,
  },
  {
    accessorKey: "totalUnit",
    accessorFn: (row) => row.amount * row.unitPrice,
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Thành tiền</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {toVND(row.original.amount * row.original.unitPrice)}
        </div>
      );
    },
    size: 4,
  },
];
export function InvoiceDetailTable({
  details,
}: {
  details: InvoiceDetailProps[];
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
  return (
    <div className="rounded-md border overflow-x-auto flex-1 min-w-full max-w-[20vw]">
      <Table>
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
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
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                Không tìm thấy kết quả.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
