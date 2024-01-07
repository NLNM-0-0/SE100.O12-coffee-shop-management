import { IconType } from "react-icons";

export type Product = {
  id: string;
  name: string;
  price: number;
  status: "active" | "inactive";
  image?: string;
  idCate?: string;
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
};
export type Unit = {
  id: string;
  name: string;
  measureType: string;
  value: number;
};

export type IngredientForChoose = {
  id: string;
  name: string;
  unitId: string;
};

export type Ingredient = {
  id: string;
  name: string;
  amount: number;
  price: number;
  unitType: {
    id: string;
    name: string;
    measureType: string;
    value: number;
  };
};
export type IngredientDetail = {
  idIngre: string;
  expirationDate: Date;
  quantity: number;
};

export type Staff = {
  address?: string;
  email: string;
  id: string;
  isActive: boolean;
  name: string;
  image: string;
  phone?: string;
  role: {
    id: string;
    name: string;
  };
};
export type Role = {
  id: string;
  name: string;
  function?: string[];
};
export type RoleFunction = {
  id: string;
  name: string;
};

export type ImportNote = {
  id: string;
  supplierId: string;
  totalPrice: number;
  status: StatusNote;
  closedAt?: Date;
  closedBy?: {
    id: string;
    name: string;
  };
  createdAt: Date;
  createdBy: {
    id: string;
    name: string;
  };
  supplier: {
    id: string;
    name: string;
    phone: string;
  };
};
export type ImportNoteDetail = {
  ingredient: {
    id: string;
    name: string;
  };
  price: number;
  amountImport: number;
  unitTypeName: string;
};
export type ExportNote = {
  id: string;
  reason: string;
  createdAt: Date;
  createdBy: {
    id: string;
    name: string;
  };
};

export type CustomerInvoice = {
  id: string;
  totalPrice: number;
  amountReceived: number;
  amountPriceUsePoint: number;
  pointUse: number;
  pointReceive: number;
  createdBy: {
    id: string;
    name: string;
  };
  createdAt: Date;
};
export enum StatusNote {
  Inprogress = "InProgress",
  Done = "Done",
  Cancel = "Cancel",
}
export enum ExportReason {
  Damaged = "Damaged",
  OutOfDate = "OutOfDate",
}
export enum MeasureType {
  Volume = "Volume",
  Weight = "Weight",
  Unit = "Unit",
}
export enum StatusActive {
  Active = "Đang giao dịch",
  InActive = "Ngừng giao dịch",
}
export type Customer = {
  id: string;
  name: string;
  email?: string;
  phone: string;
  point: number;
};

export type Supplier = {
  id: string;
  name: string;
  email?: string;
  phone: string;
  debt: number;
};
export type SupplierDebt = {
  createdAt: Date;
  createdBy: {
    id: string;
    name: string;
  };
  id: string;
  amount: number;
  amountLeft: number;
  supplierId: string;
  type: string;
};
export interface UnitListProps {
  unit: string;
  setUnit: (unit: string) => void;
}
export interface RoleListProps {
  role: string;
  setRole: (role: string) => void;
}
export type FilterValue = {
  filters: {
    type: string;
    value: string;
  }[];
};
export interface StaffListProps {
  staff: string;
  setStaff: (role: string) => void;
}
export interface CategoryListProps {
  checkedCategory: Array<string>;
  onCheckChanged: (idCate: string) => void;
  canAdd?: boolean;
  readonly?: boolean;
  isEdit?: boolean;
  onRemove?: (index: number) => void;
}
export type PagingProps = {
  page: number;
  limit: number;
  total: number;
};
export interface StatusListProps {
  status?: boolean;
  setStatus: (role: boolean) => void;
  display: { trueText: string; falseText: string };
}
