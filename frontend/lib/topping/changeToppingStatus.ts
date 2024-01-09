import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function changeToppingStatus({
  toppingIds,
  isActive,
}: {
  isActive: boolean;
  toppingIds: string[];
}) {
  const url = `${endPoint}/toppings/status`;

  const data = {
    isActive: isActive,
    ids: toppingIds,
  };
  const token = await getApiKey();
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    // Add other headers as needed
  };
  console.log(data);
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
