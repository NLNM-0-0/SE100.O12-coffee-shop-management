import { useState } from "react";
import { Input } from "../ui/input";

import { Button } from "../ui/button";
import {
  Control,
  Controller,
  UseFormReturn,
  useFieldArray,
  useWatch,
} from "react-hook-form";

import { AiOutlineClose } from "react-icons/ai";
import { CiBoxes } from "react-icons/ci";
import { IoMdInformationCircleOutline } from "react-icons/io";
import { z } from "zod";
import { AutoComplete } from "../autocomplete";
import { Ingredient } from "@/types";
import { toVND } from "@/lib/utils";
import getAllIngredient from "@/lib/getAllIngredient";
import Loading from "../loading";
import UnitListType from "./unit-list-type";
import { FormSchema } from "@/app/(normal)/stock-manage/export/add-note/page";

const IngredientExport = ({
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
        amountExport: 0,
        unitTypeId: item.unitType.id,
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
        <div className="text-sm">
          <div className="grid grid-cols-12 lg:gap-3 gap-2 font-medium py-2 px-2 mt-2 rounded-t-md bg-[#ffe9db]">
            <h2 className="col-span-5">Tên nguyên liệu</h2>
            <h2 className="col-span-3 text-left">Đơn vị</h2>
            <h2 className="col-span-3 text-left">Số lượng</h2>
          </div>
          <div className="border border-t-0 py-2 rounded-b-md">
            {fieldsIngre.length < 1 ? (
              <div className="flex flex-col items-center gap-4 py-8 text-muted-foreground font-medium">
                <CiBoxes className="h-24 w-24 text-muted-foreground/40" />
                {errors.details ? (
                  <p className="error___message">{errors.details.message}</p>
                ) : (
                  "Chọn sản phẩm xuất kho"
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
                    className="grid grid-cols-12 p-2 lg:gap-3 gap-2 items-start"
                  >
                    <div className="flex col-span-5 items-center">
                      <h2 className="font-medium min-w-[4rem]">
                        {value?.name}
                      </h2>
                      <span className="font-normal text-muted-foreground">
                        (Tồn: {value.amount} {value.unitType.name})
                      </span>
                    </div>
                    <div className="col-span-3 min-w-[4rem]">
                      <Controller
                        control={control}
                        name={`details.${index}.unitTypeId`}
                        render={({ field }) => (
                          <UnitListType
                            measureType={value.unitType.measureType}
                            unit={field.value}
                            setUnit={(value) => field.onChange(value)}
                          />
                        )}
                      />
                    </div>
                    <div className="col-span-3 flex">
                      <div>
                        <Input
                          defaultValue={ingre.amountExport}
                          {...register(
                            `details.${index}.amountExport` as const
                          )}
                        ></Input>
                        {errors &&
                        errors.details &&
                        errors.details[index] &&
                        (errors.details[index]!.amountExport as
                          | { message: string }
                          | undefined) ? (
                          <span className="error___message">
                            {errors.details[index]!.amountExport!.message}
                          </span>
                        ) : null}
                      </div>
                    </div>
                    <div className="text-right flex justify-end gap-2 items-center col-span-1">
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
              } else {
                //TODO
              }
            })}
          </div>
        </div>
      </div>
    );
  }
};

export default IngredientExport;
