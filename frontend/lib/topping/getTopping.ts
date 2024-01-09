import { endPoint } from "@/constants";
import { Staff } from "@/types";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

export type ToppingProps = {
  id: string;
  image: string;
  name: string;
  description: string;
  cookingGuide: string;
  isActive: boolean;
  price: number;
  cost: number;
  recipe: {
    details: {
      ingredient: {
        id: string;
        name: string;
        unitType: {
          id: string;
          name: string;
          measureType: string;
          value: number;
        };
      };
      amountNeed: number;
    }[];
  };
};
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

export default function getFood(idFood: string) {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/toppings/${idFood}`,
    fetcher
  );

  return {
    data: data as ToppingProps,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
