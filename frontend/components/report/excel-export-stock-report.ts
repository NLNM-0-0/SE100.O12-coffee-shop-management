import { toLocalTime } from "@/lib/utils";
import { StockReport} from "@/types";
import { saveAs } from "file-saver";

export const ExportStockReport = (excelData: StockReport, fileName: string) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet("Stocks");

  // set title cell
  sheet.mergeCells("A1", "I1");
  sheet.getCell("A1").value = "Báo cáo tồn kho";
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

  sheet.getCell("A2").value = "Từ";
  sheet.mergeCells("B2", "I2");
  sheet.getCell("B2").value = toLocalTime(excelData.timeFrom);

  sheet.getCell("A3").value = "Đến";
  sheet.mergeCells("B3", "I3");
  sheet.getCell("B3").value = toLocalTime(excelData.timeTo);

  // set columns id
  sheet.columns = [
    { key: "id", width: 10 },
    { key: "name", width: 36 },
    { key: "unitTypeName", width: 10 },
    { key: "initial", width: 20 },
    { key: "import", width: 20 },
    { key: "export", width: 20 },
    { key: "modify", width: 20 },
    { key: "sell", width: 20 },
    { key: "final", width: 20 },
  ];

  sheet.addRow(1)
  
  // set column headers
  sheet.addRow(5).values = [
    "ID",
    "Tên nguyên liệu",
    "Đơn vị tính",
    "Tồn đầu",
    "Nhập",
    "Xuất",
    "Kiểm kho",
    "Bán",
    "Tồn cuối",
  ];

  // add data
  excelData.details.forEach((detail) => {
    let columnDetail={
      id: detail.ingredient.id,
      name: detail.ingredient.name,
      unitTypeName: detail.ingredient.unitType.name,
      initial: detail.initial,
      import: detail.import,
      export: detail.export,
      modify: detail.modify,
      sell: detail.sell,
      final: detail.final,
    }
    sheet.addRow(columnDetail);
  });

  // style header row
  sheet.getRow(5).eachCell({ includeEmpty: true }, function (cell: any) {
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

  // sheet global font size
  sheet.eachRow((row: any) => {
    row.eachCell((cell: any) => {
      // default styles
      if (!cell.font?.size) {
        cell.font = Object.assign(cell.font || {}, { size: 13 });
      }
      sheet.getCell(cell.address).border = {
        top: { style: "thin" },
        left: { style: "thin" },
        bottom: { style: "thin" },
        right: { style: "thin" },
      };
    });
  });

  workbook.xlsx
    .writeBuffer()
    .then((buffer: any) => saveAs(new Blob([buffer]), fileName))
    .catch((err: any) => console.log("Error writing excel export", err));
};
