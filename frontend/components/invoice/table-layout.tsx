import { Invoice, PagingProps } from "@/types";
import InvoiceTable from "./table";
import getAllInvoice from "@/lib/invoice/getAllInvoice";
import { useSearchParams } from "next/navigation";

const TableLayout = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const getFilterString = () => {
    const maxPrice = searchParams["maxPrice"] ?? undefined;
    const minPrice = searchParams["minPrice"] ?? undefined;
    const createdBy = searchParams["createdBy"] ?? undefined;
    const search = searchParams["search"] ?? undefined;
    const createdAtFrom = searchParams["createdAtFrom"] ?? undefined;
    const createdAtTo = searchParams["createdAtTo"] ?? undefined;
    "createdAtFrom" ?? undefined;
    let filters = [{ type: "", value: "" }];
    filters.pop();
    if (maxPrice) {
      filters = filters.concat({
        type: "maxPrice",
        value: maxPrice.toString(),
      });
    }
    if (minPrice) {
      filters = filters.concat({
        type: "minPrice",
        value: minPrice.toString(),
      });
    }
    if (search) {
      filters = filters.concat({ type: "search", value: search.toString() });
    }
    if (createdBy) {
      filters = filters.concat({
        type: "createdBy",
        value: createdBy.toString(),
      });
    }
    if (createdAtFrom) {
      filters = filters.concat({
        type: "createdAtFrom",
        value: createdAtFrom.toString(),
      });
    }
    if (createdAtTo) {
      filters = filters.concat({
        type: "createdAtTo",
        value: createdAtTo.toString(),
      });
    }

    let stringToFilter = "";
    filters.forEach((item) => {
      stringToFilter = stringToFilter.concat(`&${item.type}=${item.value}`);
    });
    return { stringToFilter: stringToFilter, filters: filters };
  };
  const page = searchParams["page"] ?? "1";
  const staffsData: Promise<{ paging: PagingProps; data: Invoice[] }> =
    getAllInvoice({
      page: +page,
      filterString: getFilterString().stringToFilter,
    });
  const staffs = await staffsData;

  const totalPage = Math.ceil(staffs.paging.total / staffs.paging.limit);
  return <InvoiceTable data={staffs.data} totalPage={totalPage} />;
};

export default TableLayout;
