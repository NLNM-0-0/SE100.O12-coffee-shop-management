import { endPoint } from "@/constants";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

export default function getExportNoteDetail({ idNote }: { idNote: string }) {
  const fetcher = async (url: string) => {
    const token = await getApiKey();
    return fetch(url, {
      headers: {
        accept: "application/json",
        Authorization: `Bearer ${token}`,
      },
      cache: "no-store",
    })
      .then((res) => {
        return res.json();
      })
      .then((json) => json.data);
  };

  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/exportNotes/${idNote}`,
    fetcher
  );

  return {
    data: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
