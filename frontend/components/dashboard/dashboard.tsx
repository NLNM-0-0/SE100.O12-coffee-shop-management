"use client";
import getDashboard from "@/lib/dashboard/getDashboard";
import { CardDashboardInfo, Dashboard } from "@/types";
import { useState } from "react";
import { toast } from "../ui/use-toast";
import DashboardHeader from "./dashboard-header";
import DashboardCardHolder from "./dashboard-card-container";
import DashboardTopFoodContainer from "./dashboard-top-food-container";
import Loading from "../loading";
import DashboardChartContainer from "./dashboard-chart-container";

const DashboardComponent = () => {
  const [data, setData] = useState<Dashboard>({
    timeFrom: new Date(),
    timeTo: new Date(),
    totalSale: 0,
    totalCustomer: 0,
    totalSold: 0,
    totalPoint: 0,
    topSoldFoods: [] as unknown as [
      { id: string; name: string; amount: number }
    ],
    chartPriceComponents: [] as unknown as [{ time: Date; value: number }],
    chartCostComponents: [] as unknown as [{ time: Date; value: number }],
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
      title: "Số khách đã mua",
      value: data.totalCustomer.toString(),
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
        <DashboardHeader onClick={onGetDashboard} />
      </div>
      {isLoading ? (
        <Loading></Loading>
      ) : (
        <div className="flex flex-col w-full lg:gap-[6] gap-4">
          <DashboardCardHolder cardInfos={cardInfos} />
          <div className="flex lg:flex-row flex-col w-full lg:gap-[6] gap-4 h-auto">
            <DashboardChartContainer
              price={data?.chartPriceComponents}
              cost={data?.chartCostComponents}
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

export default DashboardComponent;
