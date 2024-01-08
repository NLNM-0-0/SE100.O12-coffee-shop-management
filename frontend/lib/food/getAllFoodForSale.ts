import { endPoint } from "@/constants";
import useSWR from "swr";
import { getApiKey } from "../auth/action";
import { ProductForSale } from "@/types";

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
    .then((json) => {
      return {
        paging: json.paging,
        data: json.data as ProductForSale[],
      };
    });
};

export default function getAllFoodForSale() {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/foods/all`,
    fetcher
  );

  return {
    foods: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
