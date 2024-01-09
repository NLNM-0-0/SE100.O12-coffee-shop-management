import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function updateFood({
  food,
}: {
  food: {
    id: string;
    name: string;
    description: string;
    cookingGuide: string;
    image: string;
    categories: string[];
    sizes: {
      sizeId: string;
      name: string;
      cost: number;
      price: number;
      recipe: {
        details: {
          ingredientId: string;
          amountNeed: number;
        }[];
      };
    }[];
  };
}) {
  const data = {
    name: food.name,
    description: food.description,
    cookingGuide: food.cookingGuide,
    image: food.image,
    categories: food.categories,
    sizes: food.sizes.map((item) => {
      return {
        ...(item.sizeId && item.sizeId !== "" && { sizeId: item.sizeId }),
        name: item.name,
        cost: item.cost,
        price: item.price,
        recipe: item.recipe,
      };
    }),
  };
  console.log(data);
  const url = `${endPoint}/foods/${food.id}`;
  const token = await getApiKey();
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .patch(url, data, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
