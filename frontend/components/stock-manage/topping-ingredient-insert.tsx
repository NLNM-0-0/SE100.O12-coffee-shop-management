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
import { FormSchema } from "@/app/(normal)/product-manage/topping/add/page";
const ToppingInsert = ({
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
    setValue: setFieldValue,
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
        amountNeed: 0,
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
          <div className="grid grid-cols-2 lg:gap-3 gap-2 font-medium py-2 px-2 mt-2 rounded-t-md bg-stone-100">
            <h2 className="col-span-1">Tên nguyên liệu</h2>
            <h2 className="col-span-1 text-left">Số lượng</h2>
          </div>
          <div className="border border-t-0 py-2 rounded-b-md min-h-[10rem]">
            {fieldsIngre.map((ingre, nestedIndex) => {
              const value = data.data.find(
                (item) => item.id === ingre.ingredientId
              );
              if (value) {
                return (
                  <div
                    key={ingre.ingredientId}
                    className="grid grid-cols-2 p-2 lg:gap-3 gap-2 items-start"
                  >
                    <div className="flex col-span-1 items-center">
                      <h2 className="font-medium min-w-[3rem]">
                        {value?.name}
                      </h2>
                    </div>
                    <div className="flex flex-col col-span-1 gap-1">
                      <div className="flex items-center">
                        <Input
                          defaultValue={ingre.amountNeed}
                          {...register(
                            `details.${nestedIndex}.amountNeed` as const
                          )}
                        ></Input>

                        <span className="w-20 px-2 font-light">
                          {value.unitType.name}
                        </span>
                        <div className="text-right flex justify-end gap-2 items-center">
                          <Button
                            type="button"
                            variant={"ghost"}
                            className={`px-3`}
                            onClick={() => {
                              removeIngre(nestedIndex);
                            }}
                          >
                            <AiOutlineClose />
                          </Button>
                        </div>
                      </div>
                      {errors &&
                      errors.details &&
                      errors.details[nestedIndex] &&
                      errors.details[nestedIndex]?.amountNeed &&
                      (errors.details[nestedIndex]!.amountNeed as
                        | { message: string }
                        | undefined) ? (
                        <span className="error___message">
                          {errors.details[nestedIndex]!.amountNeed?.message}
                        </span>
                      ) : null}
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

export default ToppingInsert;
