"use client";
import { useEffect, useRef, useState } from "react";
import Image from "next/image";
import { Input } from "../ui/input";
import { ProductForSale } from "@/types";
import {
  FieldArrayWithId,
  UseFieldArrayAppend,
  UseFieldArrayUpdate,
  useFieldArray,
} from "react-hook-form";
import { removeAccents, toVND } from "@/lib/utils";
import { AspectRatio } from "@radix-ui/react-aspect-ratio";
import { useCallback } from "react";
import { FormValues } from "@/app/(normal)/order/page-layout";
import getAllCategoryList from "@/lib/category/getAllCategoryList";
const ProductTab = ({
  append,
  fields,
  update,
  foods,
}: {
  append: UseFieldArrayAppend<FormValues, "details">;
  fields: FieldArrayWithId<FormValues, "details", "id">[];
  update: UseFieldArrayUpdate<FormValues, "details">;
  foods: ProductForSale[] | undefined;
}) => {
  const { categories, isLoading, isError, mutate } = getAllCategoryList();
  const [cateList, setCateList] = useState<
    {
      id: string;
      name: string;
      isSelected: boolean;
    }[]
  >();
  const [all, setAll] = useState(true);
  const [prodList, setProdList] = useState<Array<ProductForSale>>();
  const handleAllSelected = () => {
    if (!all) {
      setAll((prev) => !prev);
      setCateList(
        categories?.data.map((item: any) => {
          return { id: item.id, name: item.name, isSelected: false };
        })
      );
      setProdList(foods);
    }
  };

  const handleCateSelected = (id: string) => {
    if (all) {
      setAll(false);
      setProdList(new Array());
    }
    const newCateList = cateList?.map((item: any) => {
      return {
        id: item.id,
        name: item.name,
        isSelected: item.id === id ? !item.isSelected : item.isSelected,
      };
    });
    setCateList(newCateList);

    if (newCateList?.every((item) => !item.isSelected)) {
      handleAllSelected();
    } else {
      const categorySet = new Set(
        newCateList
          ?.filter((item: any) => item.isSelected === true)
          .map((value) => value.id)
      );
      const newProdList = new Array<ProductForSale>();
      foods?.forEach((prod: ProductForSale) => {
        for (let element of prod.categories) {
          if (categorySet.has(element.category.id)) {
            newProdList.push(prod);
            break;
          }
        }
      });
      setProdList(newProdList);
    }
  };

  const [inputValue, setInputValue] = useState<string>("");
  const [filteredList, setFilteredList] = useState<Array<ProductForSale>>();

  // Search Handler
  const searchHandler = useCallback(() => {
    const filteredData = prodList?.filter((prod) => {
      return prod.name.toLowerCase().includes(inputValue.toLowerCase());
    });
    setFilteredList(filteredData);
  }, [prodList, inputValue]);

  // EFFECT: Search Handler
  useEffect(() => {
    // Debounce search handler
    const timer = setTimeout(() => {
      searchHandler();
    }, 500);

    // Cleanup
    return () => {
      clearTimeout(timer);
    };
  }, [searchHandler]);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    document.addEventListener("keydown", detectKeyDown, true);
  }, []);
  const detectKeyDown = (e: any) => {
    if (e.key === "F2") {
      inputRef.current?.focus();
    }
    return () => {
      document.removeEventListener("keydown", detectKeyDown);
    };
  };
  useEffect(() => {
    if (categories) {
      setCateList(
        categories?.data.map((item: any) => {
          return { id: item.id, name: item.name, isSelected: false };
        })
      );
    }
  }, [categories]);

  useEffect(() => {
    if (foods) {
      setProdList(foods);
    }
  }, [foods]);

  return (
    <div className="flex flex-col gap-6">
      <div className="flex items-end">
        <Input
          ref={inputRef}
          className=" bg-white rounded-xl"
          placeholder="Tìm kiếm sản phẩm (F2)"
          value={inputValue}
          onChange={(e) => {
            setInputValue(e.target.value);
          }}
        ></Input>
      </div>

      {/* Category list */}
      <div className="flex flex-wrap gap-2">
        <div
          className={` rounded-xl flex self-start px-3 py-1 border-gray-200 border text-sm  cursor-pointer ${
            all
              ? "bg-orange-50 border-primary text-brown font-medium"
              : "bg-white text-muted-foreground"
          }`}
          onClick={handleAllSelected}
        >
          Tất cả
        </div>
        {cateList?.map((item) => (
          <div
            key={item.id}
            className={`rounded-xl flex self-start px-3 py-1 border-gray-200 border text-sm font-medium cursor-pointer
            ${
              item.isSelected
                ? "bg-orange-50 border-primary text-brown "
                : "bg-white text-muted-foreground"
            }`}
            onClick={() => handleCateSelected(item.id)}
          >
            {item.name}
          </div>
        ))}
      </div>
      <h1 className="text-lg">Sản phẩm</h1>

      {/* Product list */}
      <div className="grid 2xl:grid-cols-5 xl:grid-cols-4 lgr:grid-cols-3 md:grid-cols-2 sm:grid-cols-4 grid-cols-3 gap-4">
        {filteredList?.map((prod) => {
          const smallestSize = prod.sizes.reduce((minSize, currentSize) => {
            return currentSize.price < minSize.price ? currentSize : minSize;
          }, prod.sizes[0]);
          return (
            <div
              key={prod.id}
              className="bg-white shadow-sm rounded-xl  overflow-hidden cursor-pointer hover:shadow-md"
              onClick={() => {
                const index = fields.findIndex(
                  (value) => value.foodId === prod.id
                );

                append({
                  foodId: prod.id,
                  amount: 1,
                  size: {
                    sizeId: smallestSize.sizeId,
                    sizeName: smallestSize.name,
                    price: smallestSize.price,
                  },
                  foodName: prod.name,
                  toppings: [],
                  description: "",
                });
              }}
            >
              <AspectRatio ratio={1 / 1}>
                <Image
                  className=" object-cover"
                  src={prod.image ?? "/no-image.jpg"}
                  alt="image"
                  fill
                  sizes="(max-width: 768px) 33vw, 20vw"
                ></Image>
              </AspectRatio>
              <div className="px-1">
                <h1 className="text-base font-medium text-center">
                  {prod.name}
                </h1>
                <h1 className="text-base font-semibold text-primary text-center pb-1">
                  {toVND(smallestSize.price)}
                </h1>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default ProductTab;
