"use client";
import React from "react";
import { InvoiceDetailTable } from "@/components/invoice/invoice-detail-table";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { toVND } from "@/lib/utils";
import Image from "next/image";
import { useRef } from "react";
import ReactToPrint, { useReactToPrint } from "react-to-print";
import { FaPrint } from "react-icons/fa";
import { InvoiceDetailProps } from "@/app/invoice/[invoiceId]/page";
const PrintInvoice = ({
  onPrint,
  responseData,
}: {
  responseData: any;
  onPrint: () => void;
}) => {
  const componentRef = useRef(null);
  const details = responseData.details as InvoiceDetailProps[];
  const handlePrint = useReactToPrint({
    content: () => componentRef.current,
  });
  return (
    <Card>
      <CardContent className="flex flex-col p-0 gap-4 relative">
        <button
          className="whitespace-nowrap text-primary-foreground shadow py-2 inline-flex h-8 shrink-0 items-center justify-center rounded-md border bg-green-500 px-3 text-sm font-medium   disabled:pointer-events-none disabled:opacity-50 hover:bg-green-500/90 boder-none"
          onClick={() => {
            handlePrint();
            onPrint();
          }}
        >
          In hóa đơn
        </button>
        <div ref={componentRef} className="printScreen">
          <div className="p-6">
            <div className="flex justify-center gap-2 mb-6">
              <Image
                className="object-contain w-auto h-auto"
                src="/android-chrome-192x192.png"
                priority
                alt="logo"
                width={50}
                height={50}
              ></Image>
              <h1 className=" py-6 text-3xl uppercase font-medium">
                Coffee Shop
              </h1>
            </div>

            <div className="flex flex-row justify-between p-4 mb-6 border rounded-md">
              <div className="flex flex-col items-end gap-2 text-sm">
                <div className="flex gap-2">
                  <span className="font-light">Mã hóa đơn:</span>
                  <span className="font-semibold">{responseData.id}</span>
                </div>
              </div>

              <div className="flex flex-col items-end gap-2 text-sm">
                <span className="font-semibold">
                  {new Date(responseData.createdAt).toLocaleTimeString(
                    "vi-VN",
                    {
                      hour: "2-digit",
                      minute: "2-digit",
                    }
                  )}
                  {", "}
                  {new Date(responseData.createdAt).toLocaleDateString("vi-VN")}
                </span>
                <span>Nhân viên: {responseData.createdBy.name}</span>
              </div>
            </div>
            <div className="flex flex-col gap-4">
              <InvoiceDetailTable {...details} />
              <div className="flex justify-end space-x-2 py-4 font-semibold">
                <span>Tổng tiền: </span>
                <span>{toVND(responseData.totalPrice)}</span>
              </div>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default PrintInvoice;
