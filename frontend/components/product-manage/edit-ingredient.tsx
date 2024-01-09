"use client";
import { Controller, SubmitHandler, useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { useState } from "react";
import { toast } from "@/components/ui/use-toast";
import { Category, Ingredient } from "@/types";
import UnitList from "../unit-list";
import updateIngredient from "@/lib/ingredient/updateIngredient";
const required = z.string().min(1, "Không để trống trường này");

const FormSchema = z.object({
  name: required,
  unitTypeId: z.string().min(1, "Không để trống trường này"),
  price: z.coerce.number().nonnegative("Đơn giá phải lớn hơn hoặc bằng 0"),
});

const EditIngredient = ({
  ingredient,
  handleIngredientEdited,
  children,
}: {
  ingredient: Ingredient;
  handleIngredientEdited: () => void;
  children: React.ReactNode;
}) => {
  const {
    register,
    handleSubmit,
    reset,
    control,
    formState: { errors, isDirty },
  } = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    setOpen(false);

    const response: Promise<any> = updateIngredient({
      idIngre: ingredient.id,
      name: data.name,
      price: data.price,
      unitTypeId: data.unitTypeId,
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
        description: "Chỉnh sửa nguyên liệu thành công",
      });
      handleIngredientEdited();
      setOpen(false);
    }
  };

  const [open, setOpen] = useState(false);
  return (
    <Dialog
      open={open}
      onOpenChange={(open) => {
        reset({
          name: ingredient.name,
          price: ingredient.price,
          unitTypeId: ingredient.unitType.id,
        });
        setOpen(open);
      }}
    >
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className="max-w-[472px] p-0 bg-white">
        <DialogHeader>
          <DialogTitle className="p-6 pb-0">Chỉnh sửa nguyên liệu</DialogTitle>
        </DialogHeader>
        <form>
          <div className="p-6 flex flex-col gap-4 border-y-[1px]">
            <div>
              <Label htmlFor="nameNcc">Tên nguyên liệu</Label>
              <Input id="nameNcc" {...register("name")}></Input>
              {errors.name && (
                <span className="error___message">{errors.name.message}</span>
              )}
            </div>
            <div className="flex gap-4">
              <div className="flex-1">
                <Label htmlFor="phone">Đơn vị</Label>
                <Controller
                  control={control}
                  name="unitTypeId"
                  render={({ field }) => (
                    <UnitList
                      unit={field.value}
                      setUnit={(value) => {
                        field.onChange(value);
                      }}
                    ></UnitList>
                  )}
                />
                {errors.unitTypeId && (
                  <span className="error___message">
                    Không để trống trường này
                  </span>
                )}
              </div>
              <div className="flex-1">
                <Label htmlFor="noBanDau">Đơn giá</Label>
                <Input
                  id="noBanDau"
                  type="number"
                  {...register("price")}
                ></Input>
                {errors.price && (
                  <span className="error___message">
                    {errors.price.message}
                  </span>
                )}
              </div>
            </div>
          </div>
          <div className="p-4 flex-1 flex justify-end">
            <div className="flex gap-4">
              <Button
                type="reset"
                variant={"outline"}
                onClick={() => setOpen(false)}
              >
                Huỷ
              </Button>

              <Button
                disabled={!isDirty}
                type="button"
                onClick={handleSubmit(onSubmit)}
              >
                Lưu
              </Button>
            </div>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export default EditIngredient;
