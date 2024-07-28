"use client";

import ImportSheet from "@/components/product-manage/import-sheet";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Controller,
  SubmitErrorHandler,
  SubmitHandler,
  useForm,
} from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { endPoint, required } from "@/constants";
import { Button } from "@/components/ui/button";
import { LuCheck } from "react-icons/lu";
import { FiTrash2 } from "react-icons/fi";
import { useState } from "react";
import { toast } from "@/components/ui/use-toast";
import { useSWRConfig } from "swr";
import getAllIngredient from "@/lib/getAllIngredient";
import Loading from "@/components/loading";
import ConfirmDialog from "@/components/confirm-dialog";
import IngredientExport from "@/components/stock-manage/ingredient-export";
import ExportNoteList from "@/components/export-note-list";
import createExportNote from "@/lib/export/createExportNote";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import NoRole from "@/components/no-role";

export const FormSchema = z.object({
  idNote: z.string().max(12, "Tối đa 12 ký tự"),
  reason: required,
  details: z
    .array(
      z.object({
        ingredientId: z.string(),
        amountExport: z.coerce.number().gte(1, "Số lượng phải lớn hơn 0"),
        unitTypeId: z.string(),
      })
    )
    .refine((details) => details.length > 0, {
      message: "Vui lòng chọn ít nhất một nguyên liệu xuất",
    }),
});

const AddNote = () => {
  const [openDialog, setOpenDialog] = useState(false);
  const { mutate } = useSWRConfig();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      idNote: "",
      reason: "",
      details: [],
    },
  });
  const {
    register,
    handleSubmit,
    control,
    setValue,
    trigger,
    watch,
    reset,
    formState: { errors, isDirty },
  } = form;
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    console.log(data);
    const response: Promise<any> = createExportNote({
      details: data.details.map((item) => {
        return {
          ingredientId: item?.ingredientId,
          amountExport: item?.amountExport,
          unitTypeId: item.unitTypeId,
        };
      }),
      id: data.idNote,
      reason: data.reason,
    });
    const responseData = await response;
    if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "success",
        title: "Thành công",
        description: "Xuất nguyên liệu thành công",
      });
      reset();
    }
  };
  const onError: SubmitErrorHandler<z.infer<typeof FormSchema>> = (data) => {
    console.log(data);
    if (data.hasOwnProperty("details")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: data.details?.message,
      });
    }
  };
  const handleFile = (reader: FileReader) => {
    let exportNote = {
      id: "",
      reason: "",
      details: [{}],
    };
    const ExcelJS = require("exceljs");
    const wb = new ExcelJS.Workbook();
    reader.onload = () => {
      const buffer = reader.result;
      wb.xlsx.load(buffer).then((workbook: any) => {
        workbook.eachSheet((sheet: any, id: any) => {
          sheet.eachRow((row: any, rowIndex: number) => {
            if (rowIndex === 1) {
              exportNote.id =
                row.getCell(2).value === "Nhập mã phiếu dưới 12 ký tự"
                  ? ""
                  : row.getCell(2).value;
            }
            if (rowIndex > 2) {
              const idIngre = row.getCell(1).value.toString();
              console.log(idIngre);
              const ing = data?.data.find((ingre) => ingre.id === idIngre);
              if (ing) {
                const detail = {
                  ingredientId: idIngre,
                  amountExport: row.getCell(3).value,
                  unitTypeId: data?.data.find((ingre) => ingre.id === idIngre)
                    ?.unitType.id,
                };

                exportNote.details.push(detail);
              } else {
                toast({
                  variant: "destructive",
                  title: "Có lỗi",
                  description: "Vui lòng kiểm tra file đúng định dạng",
                });
                return;
              }
            }
          });
          console.log(exportNote);
          reset({
            idNote: exportNote.id,
            reason: "",
            details: exportNote.details.filter(
              (value) => JSON.stringify(value) !== "{}"
            ),
          });
        });
      });
    };
  };

  const { data, isLoading, isError } = getAllIngredient();
  const { currentUser } = useCurrentUser();
  if (isError) return <div>Failed to load</div>;
  else if (!currentUser || isLoading) {
    return <Loading />;
  } else if (
    currentUser &&
    !includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["EXP_CREATE"],
    })
  ) {
    return <NoRole></NoRole>;
  } else {
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full px-0">
          <div className="flex justify-between gap-2">
            <h1 className="font-medium text-xxl self-start">Thêm phiếu xuất</h1>
            <ImportSheet
              sampleFileLink="/export-sample.xlsx"
              handleFile={handleFile}
            />
          </div>

          <form>
            <div className="flex flex-col gap-4">
              <div className="flex lg:flex-row flex-col gap-4">
                <Card className="flex-1">
                  <CardContent className="lg:p-6 p-4 flex lg:flex-row flex-col gap-4">
                    <div className="flex-1">
                      <Label>Mã phiếu</Label>
                      <Input
                        placeholder="Mã sinh tự động nếu để trống"
                        {...register("idNote")}
                      ></Input>
                      {errors.idNote && (
                        <span className="error___message">
                          {errors.idNote.message}
                        </span>
                      )}
                    </div>
                    <div className="flex-1">
                      <Label>Lý do</Label>
                      <Controller
                        control={control}
                        name="reason"
                        render={({ field }) => (
                          <ExportNoteList
                            status={field.value}
                            setStatus={(value) => field.onChange(value)}
                          />
                        )}
                      />
                      {errors.reason && (
                        <span className="error___message">
                          {errors.reason.message}
                        </span>
                      )}
                    </div>
                  </CardContent>
                </Card>
              </div>
              <Card>
                <CardContent className="lg:p-6 p-4">
                  <IngredientExport form={form} />
                </CardContent>
              </Card>
              <div className="flex md:justify-end justify-stretch gap-2">
                <Button
                  className="px-4 bg-white md:flex-none flex-1"
                  disabled={!isDirty}
                  variant={"outline"}
                  type="button"
                  onClick={() => {
                    reset({
                      idNote: "",
                      reason: "",
                      details: [],
                    });
                  }}
                >
                  <div className="flex flex-wrap gap-2 items-center">
                    <FiTrash2 className="text-muted-foreground" />
                    Hủy
                  </div>
                </Button>
                <ConfirmDialog
                  title="Xác nhận tạo phiếu xuất"
                  description="Bạn xác nhận muốn tạo phiếu xuất ?"
                  handleYes={() => handleSubmit(onSubmit)()}
                >
                  <Button className="px-4 pl-2 md:flex-none  flex-1">
                    <div className="flex flex-wrap gap-2 items-center">
                      <LuCheck />
                      Thêm
                    </div>
                  </Button>
                </ConfirmDialog>
              </div>
            </div>
          </form>
        </div>
      </div>
    );
  }
};

export default AddNote;
