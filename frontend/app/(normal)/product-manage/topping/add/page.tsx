"use client";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Card, CardContent } from "@/components/ui/card";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import {
  SubmitErrorHandler,
  SubmitHandler,
  useFieldArray,
  useForm,
} from "react-hook-form";
import { AiOutlineClose } from "react-icons/ai";
import { required } from "@/constants";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { toast } from "@/components/ui/use-toast";
import ConfirmDialog from "@/components/confirm-dialog";
import { LuCheck } from "react-icons/lu";
import { FiTrash2 } from "react-icons/fi";
import { Textarea } from "@/components/ui/textarea";
import { useRouter } from "next/navigation";
import ToppingInsert from "@/components/stock-manage/topping-ingredient-insert";
import createTopping from "@/lib/topping/createTopping";
import { useCurrentUser } from "@/hooks/use-user";
import Loading from "@/components/loading";
import { includesRoles } from "@/lib/utils";
import NoRole from "@/components/no-role";
export const FormSchema = z.object({
  id: z.string().max(12, "Tối đa 12 ký tự"),
  name: required,
  description: z.string(),
  cookingGuide: z.string(),
  cost: z.coerce
    .number()
    .int("Giá vốn phải là số nguyên")
    .gt(0, "Giá vốn phải lớn hơn 0"),
  price: z.coerce
    .number()
    .int("Giá bán phải là số nguyên")
    .gt(0, "Giá bán phải lớn hơn 0"),
  details: z
    .array(
      z.object({
        ingredientId: required,
        amountNeed: z.coerce.number().gt(0, "Số lượng phải dương"),
      })
    )
    .refine((value) => value.length > 0, {
      message: "Vui lòng có ít nhất một nguyên liệu",
    }),
});
const InsertProductPage = () => {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      id: "",
      name: "",
      details: [],
      cost: 0,
      price: 0,
      description: "",
      cookingGuide: "",
    },
  });
  const {
    register,
    handleSubmit,
    control,
    watch,
    reset,
    formState: { errors, isDirty },
  } = form;

  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "details",
  });
  const router = useRouter();
  const onFormSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (
    data
  ) => {
    console.log(data);
    const topping = {
      id: data.id,
      name: data.name,
      price: data.price,
      cost: data.cost,
      description: data.description,
      cookingGuide: data.cookingGuide,
      recipe: {
        details: data.details,
      },
    };
    const response: Promise<any> = createTopping({ topping: topping });
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
        description: "Thêm topping thành công",
      });
      reset({
        id: "",
        name: "",
        details: [],
        cost: 0,
        price: 0,
        description: "",
        cookingGuide: "",
      });
      router.refresh();
    }
  };
  const onErrors: SubmitErrorHandler<z.infer<typeof FormSchema>> = (data) => {
    console.log(data);
  };
  const { currentUser } = useCurrentUser();

  if (!currentUser) {
    return <Loading />;
  } else if (
    currentUser &&
    !includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["TOP_CREATE"],
    })
  ) {
    return <NoRole></NoRole>;
  } else
    return (
      <div className="col items-center">
        <div className="col w-full xl:px-16">
          <div className="flex justify-between">
            <h1 className="font-medium text-xxl self-start">Thêm topping</h1>
            <div className="flex md:justify-end justify-stretch gap-2">
              <Button
                className="px-4 bg-white"
                disabled={!isDirty}
                variant={"outline"}
                type="button"
                onClick={() => {
                  reset({
                    id: "",
                    name: "",
                    details: [],
                    cost: 0,
                    price: 0,
                    description: "",
                    cookingGuide: "",
                  });
                }}
              >
                <div className="flex flex-wrap gap-2 items-center">
                  <FiTrash2 className="text-muted-foreground" />
                  Hủy
                </div>
              </Button>
              <ConfirmDialog
                title="Xác nhận"
                description="Bạn xác nhận muốn tạo topping ?"
                handleYes={() => handleSubmit(onFormSubmit, onErrors)()}
              >
                <Button className="px-4 pl-2">
                  <div className="flex flex-wrap gap-2 items-center">
                    <LuCheck />
                    Thêm
                  </div>
                </Button>
              </ConfirmDialog>
            </div>
          </div>
          <form>
            <div className="flex-1 flex flex-col gap-4">
              <Card>
                <CardContent className="p-6 flex flex-col gap-4">
                  <div className="flex flex-col gap-4 2xl:flex-row">
                    <div className="flex-1">
                      <Label htmlFor="masp">Mã topping</Label>
                      <Input
                        id="masp"
                        placeholder="Mã sinh tự động nếu để trống"
                        {...register(`id` as const)}
                      ></Input>
                    </div>
                    <div className="flex-1">
                      <Label htmlFor="prodName">Tên topping</Label>
                      <Input
                        id="prodName"
                        {...register(`name` as const)}
                      ></Input>
                      {errors.name && (
                        <span className="error___message">
                          {errors.name.message}
                        </span>
                      )}
                    </div>
                  </div>
                  <div className="flex flex-col gap-4 2xl:flex-row ">
                    <div className={`flex-1  `}>
                      <Label>Giá bán (VND)</Label>
                      <Input {...register(`price` as const)}></Input>
                      {errors.price && (
                        <span className="error___message">
                          {errors.price.message}
                        </span>
                      )}
                    </div>
                    <div className="flex-1">
                      <Label>Giá vốn (VND)</Label>
                      <Input {...register(`cost` as const)}></Input>
                      {errors.cost && (
                        <span className="error___message">
                          {errors.cost.message}
                        </span>
                      )}
                    </div>
                  </div>
                </CardContent>
              </Card>
              {/* ingredient list */}
              <Card>
                <CardContent className="p-6 flex flex-col gap-4">
                  {/* ingredients list  */}

                  <ToppingInsert form={form} />
                </CardContent>
              </Card>
              <Card>
                <CardContent className="p-6 flex flex-col gap-4">
                  <div>
                    <Label>Mô tả</Label>
                    <Textarea {...register("description")} />
                  </div>
                  <div>
                    <Label>Công thức nấu</Label>
                    <Textarea {...register("cookingGuide")} />
                  </div>
                </CardContent>
              </Card>
            </div>
          </form>
        </div>
      </div>
    );
};

export default InsertProductPage;
