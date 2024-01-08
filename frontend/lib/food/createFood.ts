import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function createFood({
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
  const url = `${endPoint}/foods`;
  const token = await getApiKey();
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .post(url, food, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
