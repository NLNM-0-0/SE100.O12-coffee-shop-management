import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function createExportNote({
  details,
  id,
  reason,
}: {
  details: {
    ingredientId: string;
    amountExport: number;
    unitTypeId: string;
  }[];
  id?: string;
  reason: string;
}) {
  const url = `${endPoint}/exportNotes`;

  const data = {
    details,
    id: id,
    reason: reason,
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
