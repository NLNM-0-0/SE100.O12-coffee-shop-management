"use client";

import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { SubmitErrorHandler, SubmitHandler, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { endPoint } from "@/constants";
import { Button } from "@/components/ui/button";
import { LuCheck } from "react-icons/lu";
import { FiTrash2 } from "react-icons/fi";
import { useState } from "react";
import { toast } from "@/components/ui/use-toast";
import { useSWRConfig } from "swr";
import CheckInsert from "@/components/stock-manage/check-insert";
import createCheckNote from "@/lib/check/createCheckNote";
import { useCurrentUser } from "@/hooks/use-user";
import Loading from "@/components/loading";
import { includesRoles } from "@/lib/utils";
import NoRole from "@/components/no-role";

export const FormSchema = z.object({
  id: z.string().max(12, "Tối đa 12 ký tự"),
  details: z
    .array(
      z
        .object({
          ingredientId: z.string(),
          difference: z.coerce.number(),
          initial: z.coerce.number(),
        })
        .refine((data) => data.difference !== 0, {
          message: "Chênh lệch phải khác 0",
        })
    )
    .nonempty("Vui lòng chọn ít nhất một mặt hàng kiểm kho"),
});

const AddNote = () => {
  const [openDialog, setOpenDialog] = useState(false);
  const { mutate } = useSWRConfig();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      id: "",
      details: [],
    },
  });
  const {
    register,
    handleSubmit,
    setValue,
    trigger,
    watch,
    reset,
    formState: { errors, isDirty },
  } = form;
  const { mutate: mutateAllIngre } = useSWRConfig();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    const response: Promise<any> = createCheckNote({
      id: data.id,
      details: data.details.map((item) => {
        return {
          ingredientId: item?.ingredientId,
          difference: item?.difference,
        };
      }),
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
        description: "Thêm phiếu kiểm kho thành công",
      });
      reset({
        id: "",
        details: [],
      });
      mutateAllIngre(`${endPoint}/ingredients/all`);
    }
  };
  const { currentUser } = useCurrentUser();
  if (!currentUser) {
    return <Loading />;
  } else if (
    currentUser &&
    !includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["ICN_CREATE"],
    })
  ) {
    return <NoRole></NoRole>;
  } else
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full px-0">
          <div className="flex justify-between gap-2">
            <h1 className="font-medium text-xxl self-start">
              Thêm phiếu kiểm kho
            </h1>
          </div>

          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col gap-4">
              <div className="flex lg:flex-row flex-col gap-4">
                <Card className="flex-1">
                  <CardContent className="lg:p-6 p-4 flex lg:flex-row flex-col gap-4">
                    <div className="flex-1">
                      <Label>Mã phiếu</Label>
                      <Input
                        placeholder="Mã sinh tự động nếu để trống"
                        {...register("id")}
                      ></Input>
                      {errors.id && (
                        <span className="error___message">
                          {errors.id.message}
                        </span>
                      )}
                    </div>
                  </CardContent>
                </Card>
              </div>
              <Card>
                <CardContent className="lg:p-6 p-4">
                  <CheckInsert form={form} />
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
                      id: "",
                      details: [],
                    });
                  }}
                >
                  <div className="flex flex-wrap gap-2 items-center">
                    <FiTrash2 className="text-muted-foreground" />
                    Hủy
                  </div>
                </Button>
                <Button className="px-4 pl-2 md:flex-none  flex-1">
                  <div className="flex flex-wrap gap-2 items-center">
                    <LuCheck />
                    Thêm
                  </div>
                </Button>
              </div>
            </div>
          </form>
        </div>
      </div>
    );
};

export default AddNote;
