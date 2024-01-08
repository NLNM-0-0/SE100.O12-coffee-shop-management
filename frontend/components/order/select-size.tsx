import { useState } from "react";
import { Button } from "../ui/button";
import {
  Dialog,
  DialogContent,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import { Label } from "../ui/label";
import { toVND } from "@/lib/utils";
const SelectSize = ({
  children,
  onSizeConfirm,
  sizes,
  selectedId,
}: {
  children: React.ReactNode;
  onSizeConfirm: (sizeId: string) => void;
  sizes: { sizeId: string; name: string; price: number }[] | undefined;
  selectedId: string;
}) => {
  const [open, setOpen] = useState(false);
  const [sizeId, setSizeId] = useState(selectedId);
  if (!sizes) {
    return children;
  } else
    return (
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogTrigger asChild>{children}</DialogTrigger>
        <DialogContent className="p-0">
          <DialogTitle className="p-6 pb-0">Chọn kích thước</DialogTitle>
          <div className="flex flex-col border-y-[1px] p-6 gap-4">
            <RadioGroup
              defaultValue={selectedId}
              onValueChange={(e: string) => setSizeId(e)}
            >
              {sizes.map((item) => (
                <div
                  key={item.sizeId}
                  className="flex py-1 justify-between font-medium"
                >
                  <div className="flex items-center space-x-2">
                    <RadioGroupItem value={item.sizeId} id={item.sizeId} />
                    <Label htmlFor={item.sizeId}>{item.name}</Label>
                  </div>
                  <Label htmlFor={item.sizeId}>{toVND(item.price)}</Label>
                </div>
              ))}
            </RadioGroup>
          </div>

          <div className="ml-auto p-6 pt-0">
            <div className="flex gap-4">
              <Button
                type="button"
                variant={"outline"}
                onClick={() => setOpen(false)}
              >
                Thoát
              </Button>

              <Button
                type="button"
                onClick={() => {
                  setOpen(false);
                  onSizeConfirm(sizeId);
                }}
              >
                Lưu
              </Button>
            </div>
          </div>
        </DialogContent>
      </Dialog>
    );
};

export default SelectSize;
