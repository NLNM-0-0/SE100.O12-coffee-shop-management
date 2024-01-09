import { Category, SidebarItem } from "@/types";
import { LuClipboardList, LuHome } from "react-icons/lu";
import { MdOutlineWarehouse } from "react-icons/md";
import { GoPeople, GoPerson } from "react-icons/go";
import { z } from "zod";
import { FaRegHandshake } from "react-icons/fa";
import { AiOutlineLineChart } from "react-icons/ai";
import { BsShop } from "react-icons/bs";
import { IoSettingsOutline } from "react-icons/io5";
export const apiKey =
  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOiJnM1cyMUE3U1IiLCJyb2xlIjoiIn0sImV4cCI6MTcwMzU1NTY5OCwiaWF0IjoxNzAzNDY5Mjk4fQ.zm-7b5WY4b98_RUuwy-9HSyYNMAzqtOnkw-Z0aOwPSI";
export const endPoint = "http://localhost:8080/v1";
export const phoneRegex = new RegExp(/(0[3|5|7|8|9])+([0-9]{8})\b/g);

export const required = z.string().min(1, "Không để trống trường này");

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
    title: "Bán hàng",
    href: "/order",
    icon: BsShop,
    submenu: false,
  },
  {
    title: "Báo cáo",
    href: "/report/stock",
    icon: AiOutlineLineChart,
    submenu: true,
    subMenuItems: [
      { title: "Báo cáo tồn kho", href: "/report/stock" },
      { title: "Báo cáo nợ", href: "/report/debt" },
      { title: "Báo cáo mặt hàng", href: "/report/sale" },
    ],
  },
  {
    title: "Quản lý hóa đơn",
    href: "/invoice",
    icon: LuClipboardList,
    submenu: false,
  },
  {
    title: "Quản lý sản phẩm",
    href: "/product-manage",
    icon: LuHome,
    submenu: true,
    subMenuItems: [
      { title: "Mặt hàng", href: "/product-manage" },
      { title: "Topping", href: "/product-manage/topping" },
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
      { title: "Xuất kho", href: "/stock-manage/export" },
      { title: "Kiểm kho", href: "/stock-manage/check" },
    ],
  },
  {
    title: "Quản lý nhà cung cấp",
    href: "/supplier",
    icon: FaRegHandshake,
    submenu: false,
  },
  {
    title: "Quản lý khách hàng",
    href: "/customer",
    icon: GoPerson,
    submenu: false,
  },
  {
    title: "Thiết lập cửa hàng",
    href: "/setting",
    icon: IoSettingsOutline,
    submenu: false,
  },
  {
    title: "Quản lý nhân viên",
    href: "/staff",
    icon: GoPeople,
    submenu: false,
  },
];

export const adminSidebarItems: SidebarItem[] = [
  {
    title: "Bán hàng",
    href: "/order",
    icon: BsShop,
    submenu: false,
  },
  {
    title: "Báo cáo",
    href: "/report/stock",
    icon: AiOutlineLineChart,
    submenu: true,
    subMenuItems: [
      { title: "Báo cáo tồn kho", href: "/report/stock" },
      { title: "Báo cáo nợ", href: "/report/debt" },
      { title: "Báo cáo mặt hàng", href: "/report/sale" },
    ],
  },
  {
    title: "Quản lý hóa đơn",
    href: "/invoice",
    icon: LuClipboardList,
    submenu: false,
  },
  {
    title: "Quản lý sản phẩm",
    href: "/product-manage",
    icon: LuHome,
    submenu: true,
    subMenuItems: [
      { title: "Mặt hàng", href: "/product-manage" },
      { title: "Topping", href: "/product-manage/topping" },
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
      { title: "Xuất kho", href: "/stock-manage/export" },
      { title: "Kiểm kho", href: "/stock-manage/check" },
    ],
  },
  {
    title: "Quản lý nhà cung cấp",
    href: "/supplier",
    icon: FaRegHandshake,
    submenu: false,
  },
  {
    title: "Quản lý khách hàng",
    href: "/customer",
    icon: GoPerson,
    submenu: false,
  },
  {
    title: "Thiết lập cửa hàng",
    href: "/setting",
    icon: IoSettingsOutline,
    submenu: false,
  },
  {
    title: "Quản lý nhân viên",
    href: "/staff",
    icon: GoPeople,
    submenu: true,
    subMenuItems: [
      { title: "Danh sách nhân viên", href: "/staff" },
      { title: "Phân quyền nhân viên", href: "/staff/role" },
    ],
  },
];
