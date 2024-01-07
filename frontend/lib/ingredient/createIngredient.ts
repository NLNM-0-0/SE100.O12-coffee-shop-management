import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function createIngredient({
  id,
  name,
  unitTypeId,
  price,
}: {
  id: string;
  name: string;
  unitTypeId: string;
  price: number;
}) {
  const url = `${endPoint}/ingredients`;
  const data = {
    id: id,
    name: name,
    unitTypeId: unitTypeId,
    price: price,
  };

  const token = await getApiKey();
  const headers = {
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    accept: "application/json",
  };

  const res = axios
    .post(url, data, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
