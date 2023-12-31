import {
  Category,
  ImportNote,
  Ingredient,
  IngredientDetail,
  IngredientForChoose,
  MeasureUnit,
  Product,
  SidebarItem,
  StatusString,
} from "@/types";
import { LuHome } from "react-icons/lu";
import { MdOutlineCoffee, MdOutlineWarehouse } from "react-icons/md";

export const products: Product[] = [
  {
    id: "ABC123",
    name: "Tra sua",
    price: 15000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AEF355",
    name: "Ca phe",
    price: 17000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "ABC243",
    name: "Sandwich",
    price: 32000,
    status: "inactive",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "GHJ123",
    name: "Sua tuoi",
    price: 12000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AFH123",
    name: "Sua chua",
    price: 18000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AEF355",
    name: "Ca phe",
    price: 17000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "ABC243",
    name: "Sandwich",
    price: 32000,
    status: "inactive",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "GHJ123",
    name: "Sua tuoi",
    price: 12000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AFH123",
    name: "Sua chua",
    price: 18000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AEF355",
    name: "Ca phe",
    price: 17000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "ABC243",
    name: "Sandwich",
    price: 32000,
    status: "inactive",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "GHJ123",
    name: "Sua tuoi",
    price: 12000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AFH123",
    name: "Sua chua",
    price: 18000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AEF355",
    name: "Ca phe",
    price: 17000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "ABC243",
    name: "Sandwich",
    price: 32000,
    status: "inactive",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "GHJ123",
    name: "Sua tuoi",
    price: 12000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AFH123",
    name: "Sua chua",
    price: 18000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AEF355",
    name: "Ca phe",
    price: 17000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "ABC243",
    name: "Sandwich",
    price: 32000,
    status: "inactive",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "GHJ123",
    name: "Sua tuoi",
    price: 12000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AFH123",
    name: "Sua chua",
    price: 18000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AEF355",
    name: "Ca phe",
    price: 17000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "ABC243",
    name: "Sandwich",
    price: 32000,
    status: "inactive",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "GHJ123",
    name: "Sua tuoi",
    price: 12000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AFH123",
    name: "Sua chua",
    price: 18000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
];

export const categories: Category[] = [
  {
    id: "12",
    name: "Ca phe",
    quantity: 5,
  },
  {
    id: "14",
    name: "Tra sua",
    quantity: 5,
  },
  {
    id: "18",
    name: "Sua chua",
    quantity: 5,
  },
  {
    id: "23",
    name: "Sinh to",
    quantity: 6,
  },
  {
    id: "25",
    name: "Nuoc ep",
    quantity: 1,
  },
  {
    id: "42",
    name: "Sandwich",
    quantity: 14,
  },
];

export const measureUnits: MeasureUnit[] = [
  {
    id: "1",
    name: "g",
  },
  {
    id: "2",
    name: "ml",
  },
  {
    id: "3",
    name: "kg",
  },
  {
    id: "4",
    name: "l",
  },
  {
    id: "5",
    name: "đơn vị",
  },
];

export const ingredientForChoose: IngredientForChoose[] = [
  {
    id: "1",
    name: "Sua",
    unitId: "2",
  },
  {
    id: "2",
    name: "Duong",
    unitId: "1",
  },
];

export const ingredients: Ingredient[] = [
  {
    id: "1",
    name: "Sua",
    price: 28000,
    unit: measureUnits[3],
    total: 20,
  },
  {
    id: "2",
    name: "Ca phe hat",
    price: 98000,
    unit: measureUnits[2],
    total: 50,
  },
  {
    id: "3",
    name: "Duong",
    price: 32000,
    unit: measureUnits[2],
    total: 5,
  },
  {
    id: "4",
    name: "Sua dac",
    price: 112000,
    unit: measureUnits[3],
    total: 15,
  },
  {
    id: "5",
    name: "Coca",
    price: 112000,
    unit: measureUnits[4],
    total: 50,
  },
];

export const ingredientDetails: IngredientDetail[] = [
  {
    idIngre: "1",
    quantity: 17,
    expirationDate: new Date(2024, 8, 29),
  },
  {
    idIngre: "1",
    quantity: 3,
    expirationDate: new Date(2023, 12, 29),
  },
];

export const importNotes: ImportNote[] = [
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusString.Inprogress,
    createAt: new Date(),
    createBy: "NV002",
  },
  {
    id: "NGAY2",
    supplierId: "DT01",
    totalPrice: 3720000,
    status: StatusString.Done,
    createAt: new Date(2023, 9, 8),
    createBy: "NV002",
  },
  {
    id: "NGAY3",
    supplierId: "DT01",
    totalPrice: 4660000,
    status: StatusString.Cancel,
    createAt: new Date(2023, 10, 1),
    createBy: "NV002",
  },
];

export const statuses = [
  {
    isActive: true,
    label: "Đang giao dịch",
  },
  {
    isActive: false,
    label: "Ngừng giao dịch",
  },
];
export const sidebarItems: SidebarItem[] = [
  {
    title: "Quản lý sản phẩm",
    href: "/product-manage",
    icon: LuHome,
    submenu: true,
    subMenuItems: [
      { title: "Danh sách mặt hàng", href: "/product-manage" },
      { title: "Danh mục", href: "/product-manage/category" },
    ],
  },
  {
    title: "Quản lý kho",
    href: "/stock-manage",
    icon: MdOutlineWarehouse,
    submenu: true,
    subMenuItems: [
      { title: "Danh sách tồn kho", href: "/stock-manage" },
      { title: "Nhập kho", href: "/stock-manage/import" },
    ],
  },
  {
    title: "Quản lý nhân viên",
    href: "/staff-manage",
    icon: MdOutlineCoffee,
  },
];
