import { endPoint } from "@/constants";
import { Role, Unit } from "@/types";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

const fetcher = async (url: string) => {
  const token = await getApiKey();
  return fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
    },
  })
    .then((res) => {
      return res.json();
    })
    .then((json) => json.data);
};

export default function getAllUnit() {
  const { data, error, isLoading } = useSWR(`${endPoint}/unitTypes`, fetcher);

  return {
    units: data as Unit[],
    isLoading,
    isError: error,
  };
}
