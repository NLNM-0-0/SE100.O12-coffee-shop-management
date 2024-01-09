import { Card } from "../ui/card";

const DashboardCard = (props: any) => {
  const { title, icon, value } = props;
  return (
    <Card className="flex flex-col lg:gap-4 xl:p-8 gap-2 p-4 flex-1">
      <p className="text-gray-500">{title}</p>
      <p className="text-xl font-semibold">{value}</p>
    </Card>
  );
};

export default DashboardCard;
