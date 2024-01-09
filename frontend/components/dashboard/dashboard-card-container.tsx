import DashboardCard from "./dashboard-card";

const DashboardCardHolder = (props: any) => {
  const { cardInfos } = props;
  return (
    <div className="flex lg:flex-row flex-col gap-4">
      <div className="flex flex-row gap-4 flex-1">
        <DashboardCard
          title={cardInfos[0].title}
          icon={cardInfos[0].icon}
          value={cardInfos[0].value}
        />
        <DashboardCard
          title={cardInfos[1].title}
          icon={cardInfos[1].value}
          value={cardInfos[1].value}
        />
      </div>
      <div className="flex flex-row lg:gap-[6] gap-4 flex-1">
        <DashboardCard
          title={cardInfos[2].title}
          icon={cardInfos[2].icon}
          value={cardInfos[2].value}
        />
        <DashboardCard
          title={cardInfos[3].title}
          icon={cardInfos[3].value}
          value={cardInfos[3].value}
        />
      </div>
    </div>
  );
};

export default DashboardCardHolder;
