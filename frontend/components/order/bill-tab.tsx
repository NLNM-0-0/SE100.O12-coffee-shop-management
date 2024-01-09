import React, { useEffect, useState } from "react";
import {
  Control,
  Controller,
  FieldArrayWithId,
  UseFieldArrayRemove,
  UseFieldArrayUpdate,
  UseFormRegister,
  UseFormSetValue,
  UseFormWatch,
  useWatch,
} from "react-hook-form";
import { Card, CardContent } from "../ui/card";
import { FiTrash2 } from "react-icons/fi";
import { HiPlus, HiMinus } from "react-icons/hi";
import { FaPen } from "react-icons/fa";
import { PiClipboardTextLight } from "react-icons/pi";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { toVND } from "@/lib/utils";
import { Label } from "../ui/label";
import { FormValues } from "@/app/order/page-layout";
import SelectSize from "./select-size";
import { ProductForSale } from "@/types";
import { TbChartBubbleFilled } from "react-icons/tb";
import SelectTopping from "./select-topping";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import CustomerList from "../customer-list";
import { endPoint } from "@/constants";
import { useSWRConfig } from "swr";
import { Checkbox } from "../ui/checkbox";
import { ShoppingBag } from "lucide-react";
import { useShop } from "@/hooks/use-shop";
const AddUp = ({
  control,
  index,
}: {
  control: Control<FormValues>;
  index: number;
}) => {
  const formValues = useWatch({
    name: `details.${index}`,
    control,
  });
  const toppingPrice = formValues.toppings.reduce(
    (acc, item) => acc + +item.price,
    0
  );

  const addUp = (+formValues.size.price + toppingPrice) * formValues.amount;
  return <span className="text-sm font-bold">{toVND(addUp)}</span>;
};

export const Total = ({ control }: { control: Control<FormValues> }) => {
  const { shop } = useShop();
  const formValues = useWatch({
    name: "details",
    control,
  });
  const total = formValues.reduce(
    (acc, current) =>
      acc +
      (current.size.price +
        current.toppings.reduce((acc, item) => acc + +item.price, 0) || 0) *
        (current.amount || 0),
    0
  );
  const totalQuantity = formValues.reduce(
    (acc, current) => acc + 1 * (current.amount || 0),
    0
  );
  return (
    <div className="flex flex-col w-max">
      <div className="flex gap-2 items-center justify-between">
        <div className="flex gap-2 items-center">
          <span className="min-w-[5rem]">Tổng tiền</span>
          <div className="pr-2 py-1">({totalQuantity})</div>
        </div>

        <h1 className="text-sm">{toVND(total)}</h1>
      </div>
      <Controller
        control={control}
        name="customer"
        render={({ field }) => {
          return (
            <div className="flex flex-col gap-2">
              <Controller
                control={control}
                name="isUsePoint"
                render={({ field: checkedField }) => {
                  const discount =
                    field.value.customerId &&
                    field.value.customerId !== "" &&
                    checkedField.value &&
                    field.value.customerPoint &&
                    shop
                      ? field.value.customerPoint * shop?.usePointPercent
                      : 0;
                  shop?.usePointPercent;
                  const finalTotal = total - discount;
                  return (
                    <div className="w-full flex flex-col gap-2">
                      {field.value.customerId &&
                        field.value.customerId !== "" &&
                        checkedField.value &&
                        shop && (
                          <div className="flex flex-col gap-2">
                            <div className="flex justify-between gap-2 items-center">
                              <span className="min-w-[5rem] \">Giảm</span>
                              <h1 className="text-sm self-end">
                                {toVND(
                                  field.value.customerPoint *
                                    (shop?.usePointPercent ?? 0)
                                )}
                              </h1>
                            </div>
                          </div>
                        )}
                      <div className="flex gap-2 items-center justify-between">
                        <div className="flex gap-2 items-center">
                          <span className="min-w-[5rem]">Thành tiền</span>
                        </div>
                        <h1 className="text-sm">{toVND(finalTotal)}</h1>
                      </div>
                    </div>
                  );
                }}
              />
            </div>
          );
        }}
      />
    </div>
  );
};

const BillTab = ({
  fields,
  setValue,
  register,
  watch,
  control,
  remove,
  update,
  foods,
  onPayClick,
  isSheet,
}: {
  fields: FieldArrayWithId<FormValues, "details", "id">[];
  setValue: UseFormSetValue<FormValues>;
  register: UseFormRegister<FormValues>;
  watch: UseFormWatch<FormValues>;
  update: UseFieldArrayUpdate<FormValues, "details">;
  onPayClick: () => void;
  control: Control<FormValues, any>;
  remove: UseFieldArrayRemove;
  isSheet?: boolean;
  foods: ProductForSale[] | undefined;
}) => {
  const invoices = watch("details");
  const [open, setOpen] = useState(false);
  const { shop } = useShop();
  useEffect(() => {
    document.addEventListener("keydown", detectKeyDown, true);
  }, []);
  const detectKeyDown = (e: any) => {
    if ((e.metaKey || e.ctrlKey) && e.key === "F7") {
      setOpen(true);
    }
    return () => {
      document.removeEventListener("keydown", detectKeyDown);
    };
  };
  const { mutate } = useSWRConfig();
  return (
    <Card className="sticky right-0 top-0 h-[86vh] overflow-hidden">
      <CardContent
        className={`flex flex-col p-0 overflow-hidden h-[86vh] ${
          isSheet ? "rounded-none" : ""
        }`}
      >
        <div className="bg-white  shadow-[0_2px_2px_-2px_rgba(0,0,0,0.2)]">
          <div className="p-4 flex flex-col gap-4">
            <Controller
              control={control}
              name="customer"
              render={({ field }) => (
                <>
                  <CustomerList
                    onRemove={() => {
                      field.onChange({
                        customerId: "",
                        customerPoint: 0,
                      });
                    }}
                    canRemove
                    handleCustomerAdded={(customerId) => {
                      mutate(`${endPoint}/customers/all`);
                      field.onChange({
                        customerId: customerId,
                        customerPoint: 0,
                      });
                    }}
                    canAdd
                    customerId={field.value.customerId}
                    setCustomerId={(id, point) =>
                      field.onChange({ customerId: id, customerPoint: point })
                    }
                  />
                  <Controller
                    control={control}
                    name="isUsePoint"
                    render={({ field: checkedField }) => (
                      <div className="flex gap-2">
                        {field.value.customerId &&
                          field.value.customerId !== "" &&
                          shop && (
                            <Checkbox
                              className="mr-2"
                              id="cbPoint"
                              checked={checkedField.value}
                              onCheckedChange={(isCheck) => {
                                if (isCheck && field.value.customerPoint > 0) {
                                  checkedField.onChange(isCheck);
                                } else if (!isCheck) {
                                  checkedField.onChange(isCheck);
                                }
                              }}
                            ></Checkbox>
                          )}

                        {field.value.customerId &&
                          field.value.customerId !== "" &&
                          shop && (
                            <Label>
                              Dùng {field.value.customerPoint} điểm (giảm{" "}
                              {toVND(
                                field.value.customerPoint *
                                  (shop?.usePointPercent ?? 0)
                              )}
                              )
                            </Label>
                          )}
                      </div>
                    )}
                  />
                </>
              )}
            />
          </div>
        </div>
        <div className="flex flex-col gap-2  overflow-auto pt-4 flex-1">
          {fields.length < 1 ? (
            <div className="flex flex-col items-center gap-4 py-8 text-muted-foreground font-medium pt-[20%]">
              <PiClipboardTextLight className="h-24 w-24 text-muted-foreground/40" />
              <span>Chọn sản phẩm</span>
            </div>
          ) : null}
          {fields.map((item, index) => {
            return (
              <div
                key={item.id}
                className={`flex ${
                  index === fields.length - 1 ? "" : "border-b"
                }  xl:px-4 px-2 pb-2 group gap-2`}
              >
                <div className="flex flex-col flex-1">
                  {/* Name size price row */}
                  <div className="flex">
                    <div className="flex basis-[35%]">
                      <h1 className="text-base font-medium">
                        {index + 1}. {item.foodName}
                      </h1>
                    </div>

                    <div className="flex flex-wrap basis-[65%] items-center justify-between xl:gap-3 gap-2">
                      {/* Quantity */}
                      <div className="flex gap-2 items-center">
                        <Button
                          className="p-[2px] bg-primary hover:bg-primary/90 rounded-full cursor-pointer text-white invisible  group-hover:visible h-5 w-5"
                          onClick={() => {
                            const quantity = +invoices.at(index)?.amount!;
                            if (quantity === 1) {
                              //TODO: remove
                            } else {
                              setValue(`details.${index}.amount`, quantity - 1);
                            }
                          }}
                        >
                          <HiMinus />
                        </Button>
                        <Input
                          type="number"
                          className="px-1 w-10 text-center [&::-webkit-inner-spin-button]:appearance-none"
                          {...register(`details.${index}.amount` as const)}
                        ></Input>

                        <Button
                          className="p-[2px] bg-primary hover:bg-primary/90 rounded-full cursor-pointer text-white invisible group-hover:visible h-5 w-5"
                          onClick={() => {
                            setValue(
                              `details.${index}.amount`,
                              +invoices.at(index)?.amount! + 1
                            );
                          }}
                        >
                          <HiPlus />
                        </Button>
                      </div>
                      <SelectSize
                        sizes={
                          foods?.find(
                            (food) => food.id === invoices[index].foodId
                          )?.sizes
                        }
                        selectedId={invoices.at(index)?.size.sizeId!}
                        onSizeConfirm={(value) => {
                          const selectedSize = foods
                            ?.find((food) => food.id === invoices[index].foodId)
                            ?.sizes.find((size) => size.sizeId === value);
                          const updatedDetails = invoices;

                          update(index, {
                            ...updatedDetails[index],
                            size: {
                              sizeName: selectedSize?.name!,
                              sizeId: selectedSize?.sizeId!,
                              price: selectedSize?.price!,
                            },
                          });
                        }}
                      >
                        <Button className="font-semibold text-sm text-center p-0 px-2 h-6 bg-orange-200 hover:bg-orange-200/90 text-primary rounded-full m-auto">
                          {invoices.at(index)?.size.sizeName}
                        </Button>
                      </SelectSize>
                      <span className="text-sm ml-auto">
                        {toVND(invoices.at(index)?.size.price!)}
                      </span>
                    </div>
                  </div>
                  {/* topping list */}
                  {invoices.at(index)?.toppings &&
                  invoices.at(index)?.toppings.length &&
                  invoices.at(index)?.toppings.length! > 0 ? (
                    <div className="flex flex-col">
                      <span className="font-medium text-sm">Topping:</span>
                      {invoices[index].toppings.map((item) => (
                        <div className="flex justify-between" key={item.id}>
                          <span className="text-muted-foreground text-sm">
                            {item.name}
                          </span>
                          <span className="text-muted-foreground text-sm">
                            {toVND(+item.price)}
                          </span>
                        </div>
                      ))}
                    </div>
                  ) : null}

                  <div className="flex">
                    <span className="text-sm font-semibold">Tổng</span>
                    <div className="text-right ml-auto">
                      <AddUp control={control} index={index} />
                    </div>
                  </div>
                  <div className="flex justify-between">
                    <div className="flex gap-1">
                      <Controller
                        control={control}
                        name={`details.${index}.toppings`}
                        render={({ field }) => {
                          return (
                            <SelectTopping
                              selectedId=""
                              checkedTopping={field.value.map(
                                (item) => item.id
                              )}
                              onCheckChanged={(id, name, price) => {
                                const selectedIndex = invoices[
                                  index
                                ].toppings.findIndex((item) => item.id === id);
                                let updatedToppings = field.value;
                                if (selectedIndex > -1) {
                                  updatedToppings.splice(selectedIndex, 1);
                                  field.onChange(updatedToppings);
                                } else {
                                  // appendCate({ idCate: idCate });
                                  updatedToppings.push({
                                    id: id,
                                    name: name,
                                    price: price,
                                  });
                                  field.onChange(updatedToppings);
                                }
                              }}
                            >
                              <div title="Topping">
                                <TbChartBubbleFilled className="text-muted-foreground h-6 w-6 cursor-pointer" />
                              </div>
                            </SelectTopping>
                          );
                        }}
                      />

                      <div className="relative flex-1 items-center">
                        <span className="w-full">
                          <input
                            id={`note${index}`}
                            className="outline-none border-0 text-sm ml-5 w-4/5"
                            placeholder="Ghi chú..."
                            {...register(
                              `details.${index}.description` as const
                            )}
                          ></input>
                        </span>

                        <Label htmlFor={`note${index}`}>
                          <FaPen className="text-muted-foreground h-3 absolute top-2 cursor-pointer" />
                        </Label>
                      </div>
                    </div>
                    <Button
                      variant={"ghost"}
                      className="h-8 p-0 px-2 rounded-lg"
                      onClick={() => remove(index)}
                    >
                      <FiTrash2 className="opacity-50" />
                    </Button>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
        <div className="flex flex-col gap-4 p-4 items-end shadow-[0_-2px_2px_-2px_rgba(0,0,0,0.2)] bg-white px-6">
          {/* Total */}
          <div className="flex flex-col">
            <div className="ml-auto">
              <Total control={control} />
            </div>
          </div>

          <RadioGroup defaultValue="all">
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="all" id="r1" />
              <Label htmlFor="r1" className="font-normal">
                Thanh toán bằng tiền mặt
              </Label>
            </div>
          </RadioGroup>
          <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
              <Button>Thanh toán (Ctrl + F7)</Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Xác nhận thanh toán</DialogTitle>
                <DialogDescription>
                  Bạn có chắc chắn muốn thanh toán
                </DialogDescription>
              </DialogHeader>
              <DialogFooter>
                <div className="flex gap-5 justify-end">
                  <Button
                    type="button"
                    variant={"outline"}
                    onClick={() => setOpen(false)}
                  >
                    Hủy
                  </Button>
                  <Button
                    type="button"
                    onClick={() => {
                      onPayClick();
                      setOpen(false);
                    }}
                  >
                    Xác nhận
                  </Button>
                </div>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </CardContent>
    </Card>
  );
};

export default BillTab;
