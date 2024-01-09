"use client";
import BillTab, { Total } from "@/components/order/bill-tab";
import ProductTab from "@/components/order/product-tab";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { FaChevronUp } from "react-icons/fa6";
import { useFieldArray, useForm } from "react-hook-form";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import getAllFoodForSale from "@/lib/food/getAllFoodForSale";
import { useState } from "react";
import { toast } from "@/components/ui/use-toast";
import createInvoice from "@/lib/invoice/createInvoice";
import Loading from "@/components/loading";
import { ToastAction } from "@/components/ui/toast";
import PrintInvoice from "@/components/invoice/print-invoice";
export type FormValues = {
  customer: {
    customerId: string;
    customerPoint: number;
  };
  isUsePoint: boolean;
  details: {
    foodId: string;
    foodName: string;
    size: {
      sizeId: string;
      sizeName: string;
      price: number;
    };
    amount: number;
    description: string;
    toppings: {
      id: string;
      name: string;
      price: string;
    }[];
  }[];
};
interface Size {
  sizeId: string;
  sizeName: string;
  price: number;
}

interface Topping {
  id: string;
  name: string;
  price: string;
}

interface FoodDetails {
  foodId: string;
  foodName: string;
  size: Size;
  amount: number;
  description: string;
  toppings: Topping[];
}
function sortToppings(toppings: Topping[]): Topping[] {
  return toppings.slice().sort((a, b) => a.id.localeCompare(b.id));
}
const OrderScreen = () => {
  const form = useForm<FormValues>({
    defaultValues: {
      customer: {},
      isUsePoint: false,
      details: [],
    },
  });
  const {
    foods,
    isLoading: isLoadingFood,
    isError: isErrorFood,
    mutate: mutateFood,
  } = getAllFoodForSale();
  const { register, control, setValue, watch, handleSubmit, reset } = form;

  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "details",
  });
  const onSubmit = async (data: FormValues) => {
    console.log(data);
    if (fields.length < 1) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: "Vui lòng chọn ít nhất một sản phẩm",
      });
      return;
    }
    let unduplicateDetails: FoodDetails[] = [];

    data.details.forEach((item: FoodDetails) => {
      const key = `${item.foodId}_${item.size.sizeId}_${JSON.stringify(
        sortToppings(item.toppings)
      )}`;

      const existingItemIndex = unduplicateDetails.findIndex(
        (existingItem) =>
          `${existingItem.foodId}_${existingItem.size.sizeId}_${JSON.stringify(
            sortToppings(existingItem.toppings)
          )}` === key
      );

      if (existingItemIndex !== -1) {
        // If the key already exists, add the amount
        unduplicateDetails[existingItemIndex].amount =
          +unduplicateDetails[existingItemIndex].amount + +item.amount;
        if (
          unduplicateDetails[existingItemIndex].description &&
          unduplicateDetails[existingItemIndex].description !== ""
        ) {
          if (item.description && item.description !== "") {
            unduplicateDetails[existingItemIndex].description +=
              ", " + item.description;
          }
        } else {
          unduplicateDetails[existingItemIndex].description = item.description;
        }
      } else {
        // If the key does not exist, create a new entry
        unduplicateDetails.push({ ...item });
      }
    });
    const createData = {
      customerId: data.customer.customerId,
      isUsePoint: data.isUsePoint,
      details: unduplicateDetails.map((item) => {
        return {
          foodId: item.foodId,
          sizeId: item.size.sizeId,
          amount: +item.amount,
          description: item.description,
          toppings: item.toppings.map((topping) => {
            return { id: topping.id };
          }),
        };
      }),
    };
    console.log(createData);
    const response: Promise<any> = createInvoice(createData);
    const responseData = await response;
    if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      const id = responseData.data.id;
      toast({
        variant: "success",
        title: "Thành công",
        description: "Thêm mới hóa đơn thành công",
        action: (
          <ToastAction altText="print" className="p-0">
            <PrintInvoice
              responseData={responseData.data.invoice}
              onPrint={() => {}}
            />
          </ToastAction>
        ),
      });
      reset({
        customer: {},
        isUsePoint: false,
        details: [],
      });
    }
  };
  const [open, setOpen] = useState(false);

  if (isLoadingFood) {
    return <Loading />;
  } else {
    return (
      <div className="flex gap-4 md:pb-0 pb-16 ">
        <div className="2xl:basis-3/5 xl:basis-1/2 md:basis-2/5  flex-1 ">
          <ProductTab
            append={append}
            update={update}
            fields={fields}
            foods={foods?.data}
          />
        </div>
        <div className="2xl:basis-2/5 xl:basis-1/2 md:basis-3/5  md:block hidden">
          <BillTab
            reset={reset}
            onPayClick={handleSubmit(onSubmit)}
            fields={fields}
            setValue={setValue}
            register={register}
            watch={watch}
            control={control}
            remove={remove}
            foods={foods?.data}
            update={update}
          />
        </div>
        <div className="fixed bottom-0 left-0 right-0">
          <Card className="md:hidden flex flex-col  h-16 bg-white rounded-none overflow-hidden">
            <div className="flex flex-1 justify-between items-center align-middle px-4">
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
                        variant={"outline"}
                        onClick={() => setOpen(false)}
                      >
                        Hủy
                      </Button>
                      <Button
                        onClick={() => {
                          handleSubmit(onSubmit)();
                          setOpen(false);
                        }}
                      >
                        Xác nhận
                      </Button>
                    </div>
                  </DialogFooter>
                </DialogContent>
              </Dialog>
              <div className="ml-auto">
                <Total control={control} />
              </div>
            </div>

            <Sheet>
              <SheetTrigger asChild>
                <Button className="w-8 h-8 absolute p-0 rounded-full top-[-14px] left-[50%]">
                  <FaChevronUp className="w-5 h-5" />
                </Button>
              </SheetTrigger>
              <SheetContent
                side={"bottom"}
                className="w-full p-0 bg-white pt-10"
              >
                <BillTab
                  onPayClick={handleSubmit(onSubmit)}
                  update={update}
                  fields={fields}
                  setValue={setValue}
                  register={register}
                  watch={watch}
                  control={control}
                  remove={remove}
                  isSheet
                  foods={foods?.data}
                  reset={reset}
                />
              </SheetContent>
            </Sheet>
          </Card>
        </div>
      </div>
    );
  }
};

export default OrderScreen;
