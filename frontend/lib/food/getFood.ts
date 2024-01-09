import { endPoint } from "@/constants";
import { Staff } from "@/types";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

export type FoodProps = {
  id: string;
  image: string;
  name: string;
  description: string;
  cookingGuide: string;
  isActive: true;
  categories: {
    category: {
      id: string;
      name: string;
    };
  }[];
  sizes: {
    foodId: string;
    sizeId: string;
    name: string;
    cost: 0;
    price: 0;
    recipe: {
      details: {
        ingredient: {
          id: string;
          name: string;
          unitType: {
            id: string;
            name: string;
            measureType: string;
            value: 1;
          };
        };
        amountNeed: number;
      }[];
    };
  }[];
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
    `${endPoint}/foods/${idFood}`,
    fetcher
  );

  return {
    data: data as FoodProps,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
