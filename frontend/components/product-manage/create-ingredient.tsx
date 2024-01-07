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
import { useRouter } from "next/navigation";
import { required } from "@/constants";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import UnitList from "../unit-list";
import createIngredient from "@/lib/ingredient/createIngredient";

const FormSchema = z.object({
  id: z.string().max(12, "Tối đa 12 ký tự"),
  name: required,
  unitTypeId: z.string().min(1, "Không để trống trường này"),
  price: z.coerce.number().nonnegative("Đơn giá phải lớn hơn hoặc bằng 0"),
});

const CreateIngredientDialog = ({
  handleIngredientAdded,
}: {
  handleIngredientAdded: (ingreId: string) => void;
}) => {
  const { currentUser } = useCurrentUser();
  const {
    register,
    handleSubmit,
    reset,
    control,
    formState: { errors },
  } = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const router = useRouter();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    setOpen(false);

    const response: Promise<any> = createIngredient(data);
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
        description: "Thêm nguyên liệu thành công",
      });
      handleIngredientAdded(responseData.data);
      // router.refresh();
    }
  };

  const [open, setOpen] = useState(false);
  if (
    !currentUser ||
    (currentUser &&
      !includesRoles({
        currentUser: currentUser,
        allowedFeatures: ["SUP_CREATE"],
      }))
  ) {
    return null;
  } else
    return (
      <Dialog
        open={open}
        onOpenChange={(open) => {
          if (open) {
            reset();
          }
          setOpen(open);
        }}
      >
        <DialogTrigger asChild>
          <Button className="lg:px-4 px-2 whitespace-nowrap">
            Tạo nguyên liệu
          </Button>
        </DialogTrigger>
        <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0 bg-white">
          <DialogHeader>
            <DialogTitle className="p-6 pb-0"> Tạo nguyên liệu</DialogTitle>
          </DialogHeader>
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="p-6 flex flex-col gap-4 border-y-[1px]">
              <div>
                <Label htmlFor="idNcc">Mã nguyên liệu</Label>
                <Input
                  id="idNcc"
                  placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                  {...register("id")}
                ></Input>
                {errors.id && (
                  <span className="error___message">{errors.id.message}</span>
                )}
              </div>
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

                <Button type="submit">Thêm</Button>
              </div>
            </div>
          </form>
        </DialogContent>
      </Dialog>
    );
};

export default CreateIngredientDialog;
