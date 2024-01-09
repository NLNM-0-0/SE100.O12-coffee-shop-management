import { useState } from "react";
import { Button } from "../ui/button";
import {
  Dialog,
  DialogContent,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "../ui/command";
import getAllToppingForSale from "@/lib/topping/getAllToppingForSale";
import { Checkbox } from "../ui/checkbox";
import { toVND } from "@/lib/utils";

export interface ToppingProps {
  checkedTopping: Array<string>;
  onCheckChanged: (
    idTopping: string,
    nameTopping: string,
    priceTopping: string
  ) => void;
  children: React.ReactNode;
  selectedId: string;
}
const SelectTopping = ({
  children,
  selectedId,
  checkedTopping,
  onCheckChanged,
}: ToppingProps) => {
  const [open, setOpen] = useState(false);
  const [toppingId, setToppingId] = useState(selectedId);
  const {
    toppings,
    isLoading: isLoadingFood,
    isError: isErrorFood,
    mutate: mutateFood,
  } = getAllToppingForSale();
  if (!toppings) {
    return children;
  } else
    return (
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogTrigger asChild>{children}</DialogTrigger>
        <DialogContent className="p-0 bg-white">
          <DialogTitle className="p-6 pb-0">Chọn topping</DialogTitle>
          <div className="flex flex-col border-y-[1px] p-6 pt-0 px-0 gap-4">
            <Command>
              <CommandInput placeholder="Tìm tên topping" />
              <CommandEmpty className="p-6 text-center">
                <div className="text-sm">Không tìm thấy topping</div>
              </CommandEmpty>
              <CommandGroup className="overflow-y-auto max-h-48">
                {toppings?.data.map((item: any) => (
                  <CommandItem
                    className="flex justify-between font-medium"
                    value={item.name}
                    key={item.id}
                    onSelect={() => {
                      onCheckChanged(item.id, item.name, item.price);
                    }}
                  >
                    <div>
                      <Checkbox
                        className="mr-2"
                        id={item.name}
                        checked={checkedTopping.includes(item.id)}
                      ></Checkbox>
                      {item.name}
                    </div>
                    <span>{toVND(item.price)}</span>
                  </CommandItem>
                ))}
              </CommandGroup>
            </Command>
          </div>
        </DialogContent>
      </Dialog>
    );
};

export default SelectTopping;
