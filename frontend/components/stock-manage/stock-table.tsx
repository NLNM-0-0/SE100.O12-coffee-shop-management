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
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
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
import { FilterValue, Ingredient } from "@/types";

import { useState } from "react";
import { Input } from "../ui/input";
import { useRouter, useSearchParams } from "next/navigation";
import getListIngredient from "@/lib/ingredient/getListIngredient";
import {
  Controller,
  SubmitHandler,
  useFieldArray,
  useForm,
} from "react-hook-form";
import Loading from "../loading";
import { useCurrentUser } from "@/hooks/use-user";
import Paging from "../paging";
import { includesRoles, toVND } from "@/lib/utils";
import { LuFilter } from "react-icons/lu";
import { Label } from "../ui/label";
import { AiOutlineClose } from "react-icons/ai";
import UnitList from "../unit-list";
import { FaPen } from "react-icons/fa";
import EditIngredient from "../product-manage/edit-ingredient";
import { useSWRConfig } from "swr";
import { endPoint } from "@/constants";

export const columns: ColumnDef<Ingredient>[] = [
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
  },
  {
    accessorKey: "id",
    header: () => {
      return <span className="font-semibold">ID</span>;
    },
    cell: ({ row }) => <div className="leading-6">{row.original.id}</div>,
    size: 4,
  },
  {
    accessorKey: "name",
    header: ({ column }) => {
      return (
        <Button
          className="p-2"
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          <span className="font-semibold">Tên nguyên liệu</span>

          <CaretSortIcon className="ml-2 h-4 w-4" />
        </Button>
      );
    },
    cell: ({ row }) => (
      <div className="capitalize pl-2 leading-6">{row.getValue("name")}</div>
    ),
  },
  {
    accessorKey: "unit",
    accessorFn: (row) => row.unitType.name,
    header: () => {
      return <span className="font-semibold">Đơn vị</span>;
    },
    cell: ({ row }) => (
      <div className="leading-6">{row.original.unitType.name}</div>
    ),
    size: 4,
  },
  {
    accessorKey: "amount",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          <CaretSortIcon className="mr-1 h-4 w-4" />

          <span className="font-semibold">Số lượng</span>
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium pr-8">
          {row.getValue("amount")}
        </div>
      );
    },
  },
  {
    accessorKey: "price",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          <CaretSortIcon className="mr-1 h-4 w-4" />

          <span className="font-semibold">Đơn giá</span>
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = row.original.price;

      return <div className="text-right font-medium">{toVND(amount)}</div>;
    },
    size: 4,
  },
  {
    accessorKey: "actions",
    header: () => {
      return <div className="font-semibold flex justify-end">Thao tác</div>;
    },
    cell: ({ row }) => {
      return <></>;
    },
  },
];
const getRowStyles = ({ row, index }: { row: any; index: number }) => {
  //TODO
  // if (index % 2 === 0) {
  //   return {
  //     background: "primary",
  //   };
  // } else {
  //   return {
  //     background: "rgb(255 247 237 / var(--tw-bg-opacity))",
  //   };
  // }
};
export const getFilterString = () => {
  const searchParams = useSearchParams();
  const minPrice = searchParams.get("minPrice") ?? undefined;
  const maxPrice = searchParams.get("maxPrice") ?? undefined;
  const minAmount = searchParams.get("minAmount") ?? undefined;
  const maxAmount = searchParams.get("maxAmount") ?? undefined;
  const unitType = searchParams.get("unitType") ?? undefined;
  const search = searchParams.get("search") ?? undefined;

  let filters = [{ type: "", value: "" }];
  filters.pop();
  if (maxPrice) {
    filters = filters.concat({ type: "maxPrice", value: maxPrice });
  }
  if (minPrice) {
    filters = filters.concat({ type: "minPrice", value: minPrice });
  }
  if (search) {
    filters = filters.concat({ type: "search", value: search });
  }
  if (minAmount) {
    filters = filters.concat({ type: "minAmount", value: minAmount });
  }
  if (maxAmount) {
    filters = filters.concat({ type: "maxAmount", value: maxAmount });
  }
  if (unitType) {
    filters = filters.concat({ type: "unitType", value: unitType });
  }

  let stringToFilter = "";
  filters.forEach((item) => {
    stringToFilter = stringToFilter.concat(`&${item.type}=${item.value}`);
  });
  return { stringToFilter: stringToFilter, filters: filters };
};
export function StockTable() {
  const searchParams = useSearchParams();
  const page = searchParams.get("page") ?? "1";

  const [latestFilter, setLatestFilter] = useState("");
  const { filters, stringToFilter } = getFilterString();
  const {
    mutate: mutate,
    ingredients: response,
    isLoading,
    isError,
  } = getListIngredient({
    page: page,
    filter: stringToFilter,
  });
  const data = response?.data;
  const totalPage = Math.ceil(response?.paging.total / response?.paging.limit);
  const router = useRouter();

  const filterValues = [
    { type: "search", name: "Từ khoá" },
    { type: "minPrice", name: "Đơn giá nhỏ nhất" },
    { type: "maxPrice", name: "Đơn giá lớn nhất" },
    { type: "minAmount", name: "Số lượng nhỏ nhất" },
    { type: "maxAmount", name: "Số lượng lớn nhất" },
    { type: "unitType", name: "Đơn vị" },
    // { type: "isActive", name: "Trạng thái" },
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

  const onSubmit: SubmitHandler<FilterValue> = async (data) => {
    let filterString = "";
    data.filters.forEach((item) => {
      filterString = filterString.concat(`&${item.type}=${item.value}`);
    });
    router.push(`/stock-manage?page=1${filterString}`);
  };
  const handleIngredientEdited = () => {
    mutate();
  };

  const { currentUser } = useCurrentUser();
  if (isError) return <div>Failed to load</div>;
  else if (isLoading) {
    return <Loading />;
  } else
    return (
      <div className="w-full">
        <div className="flex items-center py-4 gap-2">
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
                <PopoverContent className="w-96 max-h-[32rem] mx-6 overflow-y-auto">
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
                            ) : item.type.includes("Price") ||
                              item.type.includes("Amount") ? (
                              <Input
                                {...register(`filters.${index}.value`)}
                                className="flex-1"
                                type="number"
                                required
                              ></Input>
                            ) : item.type.includes("unitType") ? (
                              <div className="flex-1">
                                <Controller
                                  control={control}
                                  name={`filters.${index}.value`}
                                  render={({ field }) => (
                                    <UnitList
                                      unit={field.value}
                                      setUnit={(value) => field.onChange(value)}
                                    />
                                  )}
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
                placeholder="Tìm kiếm nguyên liệu"
                value={
                  (table.getColumn("name")?.getFilterValue() as string) ?? ""
                }
                onChange={(event) =>
                  table.getColumn("name")?.setFilterValue(event.target.value)
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
                      {item.type.includes("Price")
                        ? toVND(+item.value)
                        : item.value}
                    </span>
                  </div>
                );
              })}
            </div>
          </div>
        </div>
        <div className="scroll___table">
          <Table>
            <TableHeader>
              {table.getHeaderGroups().map((headerGroup) => (
                <TableRow key={headerGroup.id}>
                  {headerGroup.headers.map((header) => {
                    if (
                      header.id.includes("actions") &&
                      (!currentUser ||
                        (currentUser &&
                          !includesRoles({
                            currentUser: currentUser,
                            allowedFeatures: ["ING_UP"],
                          })))
                    ) {
                      return null;
                    } else
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
                    // style={
                    //   typeof getRowStyles === "function"
                    //     ? getRowStyles({ row: row, index: index })
                    //     : {}
                    // }
                  >
                    {row.getVisibleCells().map((cell) => (
                      <TableCell key={cell.id}>
                        {cell.id.includes("actions") ? (
                          currentUser &&
                          includesRoles({
                            currentUser: currentUser,
                            allowedFeatures: ["ING_UP"],
                          }) ? (
                            <div className=" flex justify-end ">
                              <EditIngredient
                                ingredient={row.original}
                                handleIngredientEdited={handleIngredientEdited}
                              >
                                <Button
                                  size={"icon"}
                                  variant={"ghost"}
                                  className="rounded-full bg-orange-200/60 hover:bg-orange-200/90 text-primary hover:text-primary"
                                >
                                  <FaPen />
                                </Button>
                              </EditIngredient>
                            </div>
                          ) : null
                        ) : (
                          flexRender(
                            cell.column.columnDef.cell,
                            cell.getContext()
                          )
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
              router.push(`/stock-manage?page=${+page + 1}${stringToFilter}`)
            }
            onNavigateBack={() =>
              router.push(`/stock-manage?page=${+page - 1}${stringToFilter}`)
            }
            totalPage={totalPage}
            onPageSelect={(selectedPage) => {
              router.push(
                `/stock-manage?page=${selectedPage}${stringToFilter}`
              );
            }}
            onNavigateFirst={() =>
              router.push(`/stock-manage?page=${1}${stringToFilter}`)
            }
            onNavigateLast={() =>
              router.push(`/stock-manage?page=${totalPage}${stringToFilter}`)
            }
          />
        </div>
      </div>
    );
}
