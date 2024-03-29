import { useState } from "react";
import { Input } from "../ui/input";

import { Button } from "../ui/button";
import {
  Control,
  UseFormReturn,
  useFieldArray,
  useWatch,
} from "react-hook-form";

import { AiOutlineClose } from "react-icons/ai";
import { CiBoxes } from "react-icons/ci";
import { IoMdInformationCircleOutline } from "react-icons/io";
import { z } from "zod";
import { toVND } from "@/lib/utils";
import { toast } from "../ui/use-toast";
import Loading from "../loading";
import { getApiKey } from "@/lib/auth/action";
import { FormSchema } from "@/app/stock-manage/check/add/page";
import { AutoComplete } from "../autocomplete";
import getAllFoodForSale from "@/lib/food/getAllFoodForSale";
import getAllIngredient from "@/lib/getAllIngredient";
import { Ingredient } from "@/types";

const Final = ({
  control,
  index,
}: {
  control: Control<z.infer<typeof FormSchema>>;
  index: number;
}) => {
  const formValues = useWatch({
    name: `details.${index}`,
    control,
  });
  const addUp = formValues.initial + +formValues.difference;
  return <p>{addUp}</p>;
};

const CheckInsert = ({
  form,
}: {
  form: UseFormReturn<z.infer<typeof FormSchema>, any, undefined>;
}) => {
  const {
    register,
    handleSubmit,
    control,
    watch,
    getValues,
    reset,
    formState: { errors },
  } = form;
  const {
    fields: fieldsIngre,
    append: appendIngre,
    remove: removeIngre,
    replace,
  } = useFieldArray({
    control: control,
    name: "details",
  });
  const { data, isLoading, isError, mutate } = getAllIngredient();
  const [value, setValue] = useState<Ingredient>();
  const handleOnValueChange = (item: Ingredient) => {
    if (!fieldsIngre.find((ingre) => ingre.ingredientId === item.id)) {
      appendIngre({
        ingredientId: item.id,
        difference: 0,
        initial: item.amount,
      });
    }
  };
  if (isError) {
    return "Failed to fetch";
  } else if (isLoading || !data) {
    return <Loading />;
  } else {
    return (
      <div className="flex flex-col">
        <AutoComplete
          options={data.data}
          emptyMessage="Không có nguyên liệu khớp với từ khóa"
          placeholder="Tìm nguyên liệu"
          onValueChange={handleOnValueChange}
          value={value}
        />
        <div>
          <div className="grid grid-cols-5 lg:gap-3 gap-2 font-medium py-2 px-2 mt-2 rounded-t-md bg-[#ffe9db]">
            <h2 className="col-span-2">Tên nguyên liệu</h2>
            <h2 className=" text-right col-span-1">Ban đầu</h2>
            <h2 className=" text-right col-span-1"> Chênh lệch</h2>
            <h2 className=" text-right col-span-1 pr-12 ">Kiểm kê</h2>
          </div>
          <div className="border border-t-0 py-2 rounded-b-md">
            {fieldsIngre.length < 1 ? (
              <div className="flex flex-col items-center gap-4 py-8 text-muted-foreground font-medium">
                <CiBoxes className="h-24 w-24 text-muted-foreground/40" />
                {errors.details?.root ? (
                  <span className="error___message">
                    {errors.details.root?.message}
                  </span>
                ) : (
                  "Chọn sản phẩm kiểm kho"
                )}
              </div>
            ) : null}
            {fieldsIngre.map((ingre, index) => {
              const value = data.data.find(
                (item) => item.id === ingre.ingredientId
              );

              if (value) {
                return (
                  <div
                    key={ingre.id}
                    className="grid grid-cols-5  p-2 lg:gap-3 gap-2 items-start"
                  >
                    <div className="col-span-2 flex justify-between">
                      <div className="flex flex-col">
                        <h2 className="font-medium">{value?.name}</h2>
                        <span className="text-sm text-light">({value.id})</span>
                      </div>
                      <h2 className="font-medium">{value?.unitType.name}</h2>
                    </div>
                    <div className="relative p-1 col-span-1 text-right">
                      <p>{ingre.initial}</p>
                    </div>

                    <div className="relative p-1 col-span-1 flex-col flex items-end">
                      <Input
                        className="text-right p-2 max-w-[8rem]"
                        type="number"
                        defaultValue={ingre.difference}
                        {...register(`details.${index}.difference` as const)}
                      ></Input>
                      {Array.isArray(errors.details) &&
                        errors.details.length > 0 &&
                        errors.details.map((detailError, idx) => {
                          if (idx === index) {
                            if (
                              detailError &&
                              detailError.root &&
                              detailError.root.message
                            )
                              return (
                                <span key={index} className="error___message">
                                  {detailError.root.message}
                                </span>
                              );
                          }
                        })}
                    </div>
                    <div className="text-right flex justify-end gap-2 items-center col-span-1">
                      <div className="text-right">
                        <Final control={control} index={index} />
                      </div>
                      <Button
                        type="button"
                        variant={"ghost"}
                        className={`px-3`}
                        onClick={() => {
                          removeIngre(index);
                        }}
                      >
                        <AiOutlineClose />
                      </Button>
                    </div>
                  </div>
                );
              }
            })}
          </div>
        </div>
      </div>
    );
  }
};

export default CheckInsert;
