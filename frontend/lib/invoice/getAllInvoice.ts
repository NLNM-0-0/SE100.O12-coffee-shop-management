import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";

export default async function getAllInvoice({
  page,
  filterString,
}: {
  filterString: string;
  page: number;
}) {
  const url = `${endPoint}/invoices?page=${page}${filterString}`;

  const token = await getApiKey();
  const res = await fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
    },
    cache: "no-store",
  });

  if (!res.ok) {
    console.error(res.json);
    return res.json();
  }

  return res.json().then((json) => {
    return {
      paging: json.paging,
      data: json.data,
    };
  });
}
