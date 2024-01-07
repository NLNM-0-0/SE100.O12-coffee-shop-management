import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function updateIngredient({
  idIngre,
  name,
  price,
  unitTypeId,
}: {
  name: string;
  idIngre: string;
  price: number;
  unitTypeId: string;
}) {
  const url = `${endPoint}/ingredients/${idIngre}`;

  const data = {
    name: name,
    unitTypeId: unitTypeId,
    price: price,
  };

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
