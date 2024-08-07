import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function createImportNote({
  details,
  id,
  supplierId,
}: {
  details: {
    ingredientId: string;
    amountImport: number;
    price: number;
    isReplacePrice?: boolean;
    unitTypeId: string;
  }[];
  id?: string;
  supplierId: string;
}) {
  const url = `${endPoint}/importNotes`;

  const data = {
    details,
    id: id,
    supplierId: supplierId,
  };
  console.log(data);
  const token = await getApiKey();
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,

    // Add other headers as needed
  };

  // Make a POST request with headers
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
