import { endPoint } from "@/constants";
import useSWR from "swr";
import { getApiKey } from "../auth/action";
import { ToppingForSale } from "@/types";

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
        data: json.data as ToppingForSale[],
      };
    });
};

export default function getAllToppingForSale() {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/toppings/all`,
    fetcher
  );

  return {
    toppings: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
