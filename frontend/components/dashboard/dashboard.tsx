"use client";
import getDashboard from "@/lib/dashboard/getReport";
import { CardDashboardInfo, Dashboard, CharComponent } from "@/types";
import { useState } from "react";
import { toast } from "../ui/use-toast";
import DashboardHeader from "./dashboard-header";
import DashboardCardHolder from "./dashboard-card-container";
import DashboardChart from "./dashboard-chart";
import { da } from "date-fns/locale";
import DashboardTopFoodContainer from "./dashboard-top-food-container";
import Loading from "../loading";
import DashboardChartContainer from "./dashboard-chart-container";

const DebtReport = () => {
  const [data, setData] = useState<Dashboard>({
    timeFrom: new Date(),
    timeTo: new Date(),
    totalSale: 0,
    totalProduct: 0,
    totalSold: 0,
    totalPoint: 0,
    topSoldFoods: [] as unknown as [
      { id: string; name: string; amount: number }
    ],
    chartSaleComponents: [] as unknown as [{ time: Date; value: number }],
    chartAmountReceiveComponents: [] as unknown as [
      { time: Date; value: number }
    ],
  });
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const onGetDashboard = async ({
    timeFrom,
    timeTo,
  }: {
    timeFrom: number;
    timeTo: number;
  }) => {
    setIsLoading(true);
    const responseData = await getDashboard({
      timeFrom: timeFrom,
      timeTo: timeTo,
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
    //   if (data == undefined || data.details.length < 1) {
    //     toast({
    //       variant: "destructive",
    //       title: "Có lỗi",
    //       description: "Không có báo cáo nợ nào",
    //     });
    //   } else {
    //     ExportDebtReport(data, "DebtReport.xlsx");
    //   }
  };

  let cardInfos: CardDashboardInfo[] = [];
  const amount = data.totalSale;
  const formatted = new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(amount);
  if (data != undefined) {
    cardInfos.push({
      title: "Doanh thu",
      value: formatted,
      icon: "",
    });
    cardInfos.push({
      title: "Số điểm tích được",
      value: data.totalPoint.toString(),
      icon: "",
    });
    cardInfos.push({
      title: "Tổng số sản phẩm",
      value: data.totalProduct.toString(),
      icon: "",
    });
    cardInfos.push({
      title: "Số sản phẩm bán được",
      value: data.totalSold.toString(),
      icon: "",
    });
  }
  return (
    <div className="flex flex-col lg:gap-[6] gap-4">
      <div>
        <DashboardHeader onClick={onGetDashboard} onExport={onExport} />
      </div>
      {isLoading ? (
        <Loading></Loading>
      ) : (
        <div className="flex flex-col lg:gap-[6] gap-4">
          <DashboardCardHolder cardInfos={cardInfos} />
          <div className="flex flex-row lg:gap-[6] gap-4 h-auto">
            <DashboardChartContainer
              sale={data?.chartSaleComponents}
              receive={data?.chartAmountReceiveComponents}
              timeFrom={data?.timeFrom}
              timeTo={data?.timeTo}
            />
            <DashboardTopFoodContainer foods={data?.topSoldFoods} />
          </div>
        </div>
      )}
    </div>
  );
};

export default DebtReport;
