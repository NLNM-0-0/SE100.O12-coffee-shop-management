import { reasonToString } from "@/lib/utils";
import { ExportNote, ExportNoteDetail, ExportReason } from "@/types";
import { saveAs } from "file-saver";

export const ExportExportNoteDetail = (
  excelData: ExportNote,
  ExportDetails: ExportNoteDetail[],
  fileName: string
) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet(excelData.id);

  // set title cell
  sheet.mergeCells("A1", "D1");
  sheet.getCell("A1").value = `Phiếu xuất`;
  sheet.getCell("A1").alignment = {
    horizontal: "center",
    vertical: "middle",
  };
  sheet.getCell("A1").border = {
    top: { style: "thin" },
    left: { style: "thin" },
    bottom: { style: "thin" },
    right: { style: "thin" },
  };
  sheet.getCell("A1").font = {
    bold: true,
    size: 18,
  };
  sheet.getRow(1).height = 40;

  // set date row
  sheet.mergeCells("A2", "D2");
  sheet.getCell("A2").value = `Ngày tạo: ${new Date(
    excelData.createdAt
  ).toLocaleDateString("vi-VN")}`;
  sheet.getCell("A2").alignment = {
    vertical: "middle",
  };
  sheet.getCell("A2").border = {
    top: { style: "thin" },
    left: { style: "thin" },
    bottom: { style: "thin" },
    right: { style: "thin" },
  };

  sheet.mergeCells("A3", "D3");
  sheet.getCell("A3").value = `Lý do: ${reasonToString(excelData.reason as ExportReason)}`;
  sheet.getCell("A3").alignment = {
    vertical: "middle",
  };
  sheet.getCell("A3").border = {
    top: { style: "thin" },
    left: { style: "thin" },
    bottom: { style: "thin" },
    right: { style: "thin" },
  };

  // set columns id
  sheet.columns = [
    { key: "id", width: 20 },
    { key: "name", width: 32 },
    { key: "unitTypeName", width: 24 },
    { key: "amountExport", width: 24 },
  ];

  // set column headers
  sheet.addRow(3).values = [
    "Id",
    "Tên nguyên liệu",
    "Đơn vị xuất",
    "Số lượng xuất"
  ];
  const values = ExportDetails.map((note) => ({
    id: note.ingredient.id,
    name: note.ingredient.name,
    unitTypeName: note.unitTypeName,
    amountExport: note.amountExport,
  }));
  values.forEach((row) => {
    sheet.addRow(row);
  });

  // style header row
  sheet.getRow(4).eachCell({ includeEmpty: true }, function (cell: any) {
    sheet.getCell(cell.address).fill = {
      type: "pattern",
      pattern: "solid",
      fgColor: { argb: "cbdff2" },
      bgColor: { argb: "cbdff2" },
    };
    sheet.getCell(cell.address).border = {
      top: { style: "thin" },
      left: { style: "thin" },
      bottom: { style: "thin" },
      right: { style: "thin" },
    };
  });

  workbook.xlsx
    .writeBuffer()
    .then((buffer: any) => saveAs(new Blob([buffer]), fileName))
    .catch((err: any) => console.log("Error writing excel export", err));
};
