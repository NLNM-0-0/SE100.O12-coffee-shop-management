"use client";

import { CaretSortIcon } from "@radix-ui/react-icons";
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

import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { CheckNote, FilterValue, StatusNote } from "@/types";

import { useEffect, useState } from "react";
import { Input } from "../ui/input";
import { statusNoteToString } from "@/lib/utils";
import { toast } from "../ui/use-toast";
import { useRouter, useSearchParams } from "next/navigation";
import Loading from "../loading";
import Paging, { PagingProps } from "../paging";
import {
  Controller,
  SubmitHandler,
  useFieldArray,
  useForm,
} from "react-hook-form";
import { LuFilter } from "react-icons/lu";
import { Label } from "../ui/label";
import { AiOutlineClose } from "react-icons/ai";
import { FilterDatePicker } from "./date-picker";
import StaffList from "../staff-list";
import getAllImportNoteForExcel from "@/lib/import/getAllImportNoteForExcel";
import SupplierList from "../supplier-list";
import ExportDialog from "../supplier-manage/export-dialog";
import StatusNoteList from "../status-note-list";
import getAllCheckNote from "@/lib/check/getAllCheck";
import { ExportCheckNote } from "./excel-check-list";
import getAllCheckNoteForExcel from "@/lib/check/getAllCheckNoteForExcel";

export const columns: ColumnDef<CheckNote>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={table.getIsAllPageRowsSelected()}
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
      />
    ),
    enableSorting: false,
    enableHiding: false,
    size: 4,
  },
  {
    accessorKey: "id",
    header: () => {
      return <span className="font-semibold">Mã phiếu</span>;
    },
    cell: ({ row }) => <div className="leading-6">{row.getValue("id")}</div>,
    size: 4,
  },
  {
    accessorKey: "createdBy",
    header: () => {
      return <span className="font-semibold">Người tạo</span>;
    },
    cell: ({ row }) => (
      <div className="leading-6 flex flex-col text-left">
        <span>{row.original.createdBy.name}</span>
      </div>
    ),
    size: 4,
  },
  {
    accessorKey: "createdAt",
    header: ({ column }) => {
      return (
        <div className="flex justify-end">
          <Button
            className="p-2 justify-end whitespace-normal"
            variant="ghost"
            onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          >
            <CaretSortIcon className=" h-4 w-4" />
            <span className="font-semibold">Ngày tạo</span>
          </Button>
        </div>
      );
    },
    cell: ({ row }) => (
      <div className="leading-6 flex flex-col text-right">
        <span>
          {new Date(row.original.createdAt).toLocaleDateString("vi-VN")}
        </span>
        <span className="font-light">{row.original.createdBy.name}</span>
      </div>
    ),

    sortingFn: "datetime",
    size: 5,
  },
];
export function CheckTable() {
  const searchParams = useSearchParams();
  const page = searchParams.get("page") ?? "1";
  const createdAtFrom = searchParams.get("createdAtFrom") ?? undefined;
  const createdAtTo = searchParams.get("createdAtTo") ?? undefined;

  const search = searchParams.get("search") ?? undefined;

  const createdBy = searchParams.get("createdBy") ?? undefined;
  useEffect(() => {
    if (createdBy) {
      setStaffCreate(createdBy);
    }
  }, [createdBy]);
  const [latestFilter, setLatestFilter] = useState("");
  const [staffCreate, setStaffCreate] = useState("");
  const [staffClose, setStaffClose] = useState("");
  const [supplierId, setSupplierId] = useState("");
  const [status, setStatus] = useState("");
  let filters = [{ type: "", value: "" }];
  filters.pop();
  if (search) {
    filters = filters.concat({ type: "search", value: search });
  }
  if (createdAtFrom) {
    filters = filters.concat({ type: "createdAtFrom", value: createdAtFrom });
  }
  if (createdAtTo) {
    filters = filters.concat({ type: "createdAtTo", value: createdAtTo });
  }
  if (createdBy) {
    filters = filters.concat({ type: "createdBy", value: createdBy });
  }

  let stringToFilter = "";
  filters.forEach((item) => {
    stringToFilter = stringToFilter.concat(`&${item.type}=${item.value}`);
  });

  const {
    mutate: mutate,
    data: response,
    isLoading,
    isError,
  } = getAllCheckNote({
    page: page,
    filterString: stringToFilter,
  });
  const data = response?.data;
  const totalPage = Math.ceil(response?.paging.total / response?.paging.limit);
  const router = useRouter();

  const filterValues = [
    { type: "search", name: "Từ khoá" },

    { type: "createdAtFrom", name: "Tạo từ ngày" },

    { type: "createdAtTo", name: "Tạo đến ngày" },

    { type: "createdBy", name: "Mã người tạo" },
  ];

  const { register, handleSubmit, reset, control, getValues } =
    useForm<FilterValue>({
      defaultValues: {
        filters: filters,
      },
    });
  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "filters",
  });
  const [openFilter, setOpenFilter] = useState(false);

  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});
  const table = useReactTable({
    data,
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
  });
  const [exportOption, setExportOption] = useState("all");
  const handleExport = async () => {
    if (exportOption === "all") {
      const importNoteData: Promise<{
        data: CheckNote[];
        paging: PagingProps;
      }> = getAllCheckNoteForExcel({ page: "1", limit: 10000 });
      const notesToExport = await importNoteData;
      if (notesToExport.data.length < 1) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "Không có phiếu kiểm kho nào",
        });
      } else {
        ExportCheckNote(notesToExport.data, `Danh sách phiếu kiểm kho.xlsx`);
      }
    } else {
      if (table.getFilteredSelectedRowModel().rows.length < 1) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "Không có phiếu kiểm kho nào",
        });
      } else {
        ExportCheckNote(
          table.getFilteredSelectedRowModel().rows.map((row) => row.original),
          `Danh sách phiếu nhập.xlsx`
        );
      }
    }
  };
  const onSubmit: SubmitHandler<FilterValue> = async (data) => {
    let filterString = "";
    data.filters.forEach((item) => {
      filterString = filterString.concat(`&${item.type}=${item.value}`);
    });
    router.push(`/stock-manage/check?page=${Number(page)}${filterString}`);
  };
  if (isError) return <div>Failed to load</div>;
  else if (isLoading) {
    return <Loading />;
  } else
    return (
      <div>
        <div className="flex items-start py-4 gap-2">
          <ExportDialog
            style="m-0"
            handleExport={handleExport}
            setExportOption={setExportOption}
            isImport
          />
          <div className="flex-1">
            <div className="flex gap-2">
              <Popover
                open={openFilter}
                onOpenChange={(open) => {
                  setOpenFilter(open);
                  reset({ filters: filters });
                }}
              >
                <PopoverTrigger asChild>
                  <Button variant="outline" className="lg:px-3 px-2">
                    Lọc
                    <LuFilter className="ml-1 h-4 w-4" />
                  </Button>
                </PopoverTrigger>
                <PopoverContent className="w-96 mx-6 max-h-[32rem] overflow-y-auto">
                  <form
                    className="flex flex-col gap-4"
                    onSubmit={handleSubmit(onSubmit)}
                  >
                    <div className="space-y-2">
                      <p className="text-sm text-muted-foreground">
                        Hiển thị phiếu nhập theo
                      </p>
                    </div>
                    <div className="flex flex-col gap-4">
                      {fields.map((item, index) => {
                        const name = filterValues.find(
                          (v) => v.type === item.type
                        );
                        return (
                          <div
                            className="flex gap-2 items-center"
                            key={item.id}
                          >
                            <Label className="basis-1/3">{name?.name}</Label>
                            {item.type === "search" ? (
                              <Input
                                {...register(`filters.${index}.value`)}
                                className="flex-1"
                                type="text"
                                required
                              ></Input>
                            ) : item.type.includes("Price") ? (
                              <Input
                                {...register(`filters.${index}.value`)}
                                className="flex-1"
                                type="number"
                                required
                              ></Input>
                            ) : item.type.includes("At") ? (
                              <Controller
                                control={control}
                                name={`filters.${index}.value`}
                                render={({ field }) => (
                                  <FilterDatePicker
                                    handleDateSelected={(date) => {
                                      const unixTimeMilliseconds =
                                        date?.getTime();
                                      const unixTimeSeconds = Math.floor(
                                        unixTimeMilliseconds! / 1000
                                      );
                                      if (item.type.includes("To")) {
                                        field.onChange(unixTimeSeconds + 86399);
                                      } else {
                                        field.onChange(unixTimeSeconds);
                                      }
                                    }}
                                    date={new Date(+field.value * 1000)}
                                  />
                                )}
                              />
                            ) : item.type === "createdBy" ? (
                              <div className="flex-1">
                                <StaffList
                                  staff={staffCreate}
                                  setStaff={(value) => {
                                    setStaffCreate(value);
                                    update(index, {
                                      type: item.type,
                                      value: value,
                                    });
                                  }}
                                />
                              </div>
                            ) : item.type === "closedBy" ? (
                              <StaffList
                                staff={staffClose}
                                setStaff={(value) => {
                                  setStaffClose(value);
                                  update(index, {
                                    type: item.type,
                                    value: value,
                                  });
                                }}
                              />
                            ) : item.type === "supplier" ? (
                              <div className="flex-1">
                                <SupplierList
                                  supplierId={supplierId}
                                  setSupplierId={(value) => {
                                    setSupplierId(value);
                                    update(index, {
                                      type: item.type,
                                      value: value,
                                    });
                                  }}
                                />
                              </div>
                            ) : item.type === "status" ? (
                              <div className="flex-1">
                                <StatusNoteList
                                  status={status}
                                  setStatus={(value) => {
                                    setStatus(value);
                                    update(index, {
                                      type: item.type,
                                      value: value,
                                    });
                                  }}
                                />
                              </div>
                            ) : null}
                            <Button
                              variant={"ghost"}
                              className={`px-3 `}
                              onClick={() => {
                                remove(index);
                              }}
                            >
                              <AiOutlineClose />
                            </Button>
                          </div>
                        );
                      })}
                    </div>
                    {fields.length === filterValues.length ? null : (
                      <div className="flex justify-end pr-12">
                        <Select
                          value={latestFilter}
                          onValueChange={(value) => {
                            append({ type: value, value: "" });
                          }}
                        >
                          <SelectTrigger className="w-[180px] flex justify-center ml-8 px-3">
                            <SelectValue placeholder="Chọn điều kiện lọc" />
                          </SelectTrigger>
                          <SelectContent>
                            <SelectGroup>
                              {filterValues.map((item) => {
                                return fields.findIndex(
                                  (v) => v.type === item.type
                                ) === -1 ? (
                                  <SelectItem key={item.type} value={item.type}>
                                    {item.name}
                                  </SelectItem>
                                ) : null;
                              })}
                            </SelectGroup>
                          </SelectContent>
                        </Select>
                      </div>
                    )}
                    <Button type="submit" className="self-end">
                      Lọc
                    </Button>
                  </form>
                </PopoverContent>
              </Popover>
              <Input
                className="flex-1"
                placeholder="Tìm kiếm mã phiếu"
                value={
                  (table.getColumn("id")?.getFilterValue() as string) ?? ""
                }
                onChange={(event) =>
                  table.getColumn("id")?.setFilterValue(event.target.value)
                }
              />
            </div>
            <div className="flex gap-2 mt-2">
              {filters.map((item, index) => {
                const name = filterValues.find((v) => v.type === item.type);
                return (
                  <div
                    key={item.type}
                    className="rounded-xl flex self-start px-3 py-1 h-fit outline-none text-sm text-primary  bg-orange-100 items-center gap-1 group"
                  >
                    <span>
                      {name?.name}
                      {": "}
                      {item.type.includes("At")
                        ? new Date(+item.value * 1000).toLocaleDateString(
                            "vi-VN"
                          )
                        : item.type === "status"
                        ? statusNoteToString(status as StatusNote)
                        : item.value}
                    </span>
                  </div>
                );
              })}
            </div>
          </div>
        </div>

        <div className="rounded-md border w-full">
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
                      <TableCell
                        key={cell.id}
                        onClick={() => {
                          if (!cell.id.includes("select")) {
                            router.push(
                              `/stock-manage/check/${row.getValue("id")}`
                            );
                          }
                        }}
                      >
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
        <div className="flex items-center justify-end space-x-2 py-4">
          <div className="flex-1 text-sm text-muted-foreground">
            {table.getFilteredSelectedRowModel().rows.length} of{" "}
            {table.getFilteredRowModel().rows.length} row(s) selected.
          </div>
          <Paging
            page={page}
            onNavigateNext={() =>
              router.push(`/stock-manage/check?page=${+page + 1}`)
            }
            onNavigateBack={() =>
              router.push(`/stock-manage/check?page=${+page - 1}`)
            }
            totalPage={totalPage}
            onPageSelect={(selectedPage) => {
              router.push(`/stock-manage/check?page=${selectedPage}`);
            }}
            onNavigateFirst={() => router.push(`/stock-manage/check?page=${1}`)}
            onNavigateLast={() =>
              router.push(`/stock-manage/check?page=${totalPage}`)
            }
          />
        </div>
      </div>
    );
}
