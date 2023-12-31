"use client";
import { Card } from "@/components/ui/card";
import ReportHeader from "@/components/report/ReportHeader";
import { DebtTable } from "@/components/supplier-manage/debt-table";
import { DebtReportTable } from "@/components/report/DebtReportTable";
import { DebtReport, DebtReportDetail } from "@/types";
import getReport from "@/lib/report/getReport";
import { useState } from "react";
import MonthReportHeader from "@/components/report/MonthReportHeader";
import { toast } from "../ui/use-toast";
import Loading from "../loading";
import { ExportDebtReport } from "./excel-export-debt-report";

const DebtReport = () => {
  const [data, setData] = useState<DebtReport>();
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const onGetDebt = async ({
    timeFrom,
    timeTo,
  }: {
    timeFrom: number;
    timeTo: number;
  }) => {
    setIsLoading(true);
    const responseData = await getReport({
      timeFrom: timeFrom,
      timeTo: timeTo,
      type: "debt",
    });
    if (responseData.hasOwnProperty("data")) {
      if (responseData.data) {
        setData(responseData.data);
      }
    } else if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: "Vui lòng thử lại sau",
      });
    }
    setIsLoading(false);
  };

  const onExport = () => {
    if (data == undefined || data.details.length < 1) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: "Không có báo cáo nợ nào",
      });
    } else {
      ExportDebtReport(data, "DebtReport.xlsx");
    }
  };

  return (
    <div>
      <div>
        <MonthReportHeader
          title="Báo cáo nợ"
          firstAction="Xem báo cáo"
          secondAction="Tải excel"
          onClick={onGetDebt}
          onExport={onExport}
        />
      </div>

      <div>
      <Card className="p-[10px] my-[22px]">
          {isLoading ? (
            <Loading />
          ) : (
            <DebtReportTable
              report={data}
              data={
                data == undefined || data == null
                  ? []
                  : (data!.details as DebtReportDetail[])
              }
            />
          )}
        </Card>
      </div>
    </div>
  );
};

export default DebtReport;
