import { endPoint } from "@/constants";
import { Ingredient } from "@/types";
import useSWR from "swr";
import { getApiKey } from "./auth/action";
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
        data: json.data as Ingredient[],
      };
    });
};
export default function getAllIngredient() {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/ingredients/all`,
    fetcher
  );

  return {
    data: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
