import { Metadata } from "next";
import InsertProductPage from "./detail-layout";
import { withAuth } from "@/lib/role/withAuth";
export const metadata: Metadata = {
  title: "Thêm mặt hàng",
};
const AddFood = () => {
  return <InsertProductPage />;
};

export default withAuth(AddFood, ["FOD_CREATE"]);
