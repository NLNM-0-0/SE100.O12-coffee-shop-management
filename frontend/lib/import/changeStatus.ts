import { endPoint } from "@/constants";
import { StatusNote } from "@/types";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function updateStatus({
  idNote,
  status,
}: {
  idNote: string;
  status: StatusNote;
}) {
  const url = `${endPoint}/importNotes/${idNote}`;
  const data = {
    status: status,
  };
  console.log(data);
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
