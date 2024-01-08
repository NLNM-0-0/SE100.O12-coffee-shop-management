import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function createTopping({
  topping,
}: {
  topping: {
    id: string;
    name: string;
    cost: number;
    price: number;
    description: string;
    cookingGuide: string;
    recipe: {
      details: {
        ingredientId: string;
        amountNeed: number;
      }[];
    };
  };
}) {
  const url = `${endPoint}/toppings`;
  const token = await getApiKey();
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .post(url, topping, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
