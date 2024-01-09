import { useState } from "react";
import { Input } from "../ui/input";

import { Button } from "../ui/button";
import { UseFormReturn, useFieldArray } from "react-hook-form";

import { AiOutlineClose } from "react-icons/ai";
import { z } from "zod";
import { AutoComplete } from "../autocomplete";
import { Ingredient } from "@/types";
import getAllIngredient from "@/lib/getAllIngredient";
import Loading from "../loading";
import { FormSchema } from "@/app/product-manage/[foodId]/page";
const SizeUpdate = ({
  form,
  sizeIndex,
}: {
  form: UseFormReturn<z.infer<typeof FormSchema>, any, undefined>;
  sizeIndex: number;
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
    name: "sizes",
  });

  const { data, isLoading, isError, mutate } = getAllIngredient();
  const [value, setValue] = useState<Ingredient>();
  const handleOnValueChange = (item: Ingredient) => {
    if (
      !fieldsIngre[sizeIndex].recipe.details.find(
        (ingre) => ingre.ingredientId === item.id
      )
    ) {
      const updatedData = addDetailsToSize(getValues(), sizeIndex, [
        {
          ingredientId: item.id,
          amountNeed: 0,
        },
      ]);
    }
  };
  const addDetailsToSize = (
    formData: z.infer<typeof FormSchema>,
    sizeIndex: number,
    newDetails: { ingredientId: string; amountNeed: number }[]
  ): z.infer<typeof FormSchema> => {
    const updatedSizes = formData.sizes.map((size, index) => {
      if (index === sizeIndex) {
        const existingDetails = size.recipe.details;

        // Update or append new details
        newDetails.forEach((newDetail) => {
          const existingDetailIndex = existingDetails.findIndex(
            (detail) => detail.ingredientId === newDetail.ingredientId
          );

          if (existingDetailIndex !== -1) {
            // If the ingredient already exists, update it
            existingDetails[existingDetailIndex] = {
              ...existingDetails[existingDetailIndex],
              ...newDetail,
            };
          } else {
            // If the ingredient does not exist, append it
            existingDetails.push(newDetail);
          }
        });

        setFieldValue(
          `sizes.${sizeIndex}.recipe.details`,
          [...existingDetails],
          {
            shouldValidate: true,
          }
        );
        return {
          ...size,
          recipe: {
            ...size.recipe,
            details: [...existingDetails],
          },
        };
      }
      return size;
    });

    // Create a new structure with updated sizes
    return {
      ...formData,
      sizes: updatedSizes,
    };
  };
  const handleRemoveDetail = ({
    sizeIndex,
    detailIndex,
  }: {
    sizeIndex: number;
    detailIndex: number;
  }) => {
    const currentSizes = getValues("sizes");

    // Remove the specific detail item
    currentSizes[sizeIndex].recipe.details.splice(detailIndex, 1);

    // Update the sizes field with the modified value
    setFieldValue("sizes", currentSizes);
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
            {getValues("sizes") &&
              getValues("sizes").at(sizeIndex) &&
              getValues("sizes")
                .at(sizeIndex)!
                .recipe.details.map((ingre, nestedIndex) => {
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
                                `sizes.${sizeIndex}.recipe.details.${nestedIndex}.amountNeed` as const
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
                                  handleRemoveDetail({
                                    sizeIndex: sizeIndex,
                                    detailIndex: nestedIndex,
                                  });
                                }}
                              >
                                <AiOutlineClose />
                              </Button>
                            </div>
                          </div>
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

export default SizeUpdate;
