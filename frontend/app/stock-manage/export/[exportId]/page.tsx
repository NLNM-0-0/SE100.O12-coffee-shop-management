"use client";
import ConfirmDialog from "@/components/confirm-dialog";
import Loading from "@/components/loading";
import { ExportExportNoteDetail } from "@/components/stock-manage/excel-export-detail";
import { ExportImportNoteDetail } from "@/components/stock-manage/excel-import-detail";
import { ExportDetailTable } from "@/components/stock-manage/export-detail-table";
import { Button } from "@/components/ui/button";
import { toast } from "@/components/ui/use-toast";
import getExportNoteDetail from "@/lib/export/getExportDetail";

import updateStatus from "@/lib/import/changeStatus";
import getImportNoteDetail from "@/lib/import/getImportDetail";
import { reasonToString, toVND } from "@/lib/utils";
import { ExportReason, StatusNote } from "@/types";

import { BiBox } from "react-icons/bi";
import { FaRegFileExcel } from "react-icons/fa";
import { FiTrash2 } from "react-icons/fi";
import { LuPackageCheck, LuPhone } from "react-icons/lu";
const ExportDetail = ({ params }: { params: { exportId: string } }) => {
  const { data, isLoading, isError, mutate } = getExportNoteDetail({
    idNote: params.exportId,
  });
  if (isError) return <div>Failed to load</div>;
  if (isLoading) {
    return <Loading />;
  } else
    return (
      <div className="flex flex-col xl:mx-[20%] gap-6">
        <div className="shadow-sm bg-white flex flex-col gap-6 md:px-8 px-4 pb-6">
          <div className="flex justify-between gap-6 font-bold text-lg border-b flex-1 py-2 pt-6">
            <div className="flex gap-4">
              <span className="font-light">Mã phiếu xuất</span>
              <span>{data.id}</span>
            </div>
            <div className="flex gap-2">
              <Button
                variant={"outline"}
                className="p-2"
                onClick={() => {
                  ExportExportNoteDetail(data, data.details, "phieuxuat.xlsx");
                }}
              >
                <FaRegFileExcel className="mr-1 h-5 w-5 text-green-700" />
                <span>Xuất file</span>
              </Button>
            </div>
          </div>
          <div className="grid grid-cols-2 text-sm">
            <div className="flex flex-col gap-4 w-fit">
              <div className="flex font-medium">
                <span className="w-16">Tạo</span>
                <div className="flex flex-col">
                  <span>{new Date(data.createdAt).toLocaleDateString()}</span>
                  <span className="font-light">{data.createdBy.name}</span>
                </div>
              </div>
            </div>
            <div className="flex flex-col items-end gap-4">
              <div className="w-fit">
                <div className="flex flex-col gap-2 font-medium">
                  <div className="flex items-center gap-1">
                    <BiBox className="h-4 w-4" />
                    {reasonToString(data.reason as ExportReason)}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="shadow-sm bg-white flex flex-col gap-6 py-6 md:px-6 px-4">
          <ExportDetailTable {...data.details} />
        </div>
      </div>
    );
};

export default ExportDetail;
