import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export type Props = {
  id: string;
  address?: string;
  image?: string;
  name?: string;
  phone?: string;
};
export default async function updateStaff({
  id,
  address,
  image,
  name,
  phone,
}: Props) {
  const url = `${endPoint}/users/${id}/info`;

  const data = {
    ...(address && { address: address }),
    ...(image && { image: image }),
    ...(name && { name: name }),
    ...(phone && { phone: phone }),
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
