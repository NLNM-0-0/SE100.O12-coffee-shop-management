import { IconType } from "react-icons";

export type Product = {
  id: string;
  name: string;
  price: number;
  status: "active" | "inactive";
  image?: string;
};

export type SidebarItem = {
  title: string;
  href: string;
  icon?: IconType;
  submenu?: boolean;
  subMenuItems?: SidebarItem[];
};

export type Category = {
  id: string;
  name: string;
  quantity: number;
};

export type MeasureUnit = {
  id: string;
  name: "g" | "kg" | "l" | "ml" | "đơn vị";
  covertDetails?: {
    measureUnit: MeasureUnit;
    times: number;
  };
};

export type IngredientForChoose = {
  id: string;
  name: string;
  unitId: string;
};

export type Ingredient = {
  id: string;
  name: string;
  total: number;
  unit: MeasureUnit;
  price: number;
};
export type IngredientDetail = {
  idIngre: string;
  expirationDate: Date;
  quantity: number;
};
export interface UnitListProps {
  unit: string;
  setUnit: (unit: string) => void;
}

export interface CategoryListProps {
  category: string;
  setCategory: (category: string) => void;
  canAdd?: boolean;
}

export enum StatusString {
  Inprogress = "Đang xử lý",
  Done = "Đã nhập",
  Cancel = "Đã huỷ",
}
export interface ImportNote {
  id: string;
  supplierId: string;
  totalPrice: number;
  status: StatusString;
  createBy: string;
  closeBy?: string;
  createAt: Date;
  closeAt?: Date;
}
