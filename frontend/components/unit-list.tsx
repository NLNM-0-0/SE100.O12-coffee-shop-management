import { UnitListProps } from "@/types";
import { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "./ui/command";
import { Check, ChevronsUpDown } from "lucide-react";
import { Button } from "./ui/button";
import { cn } from "@/lib/utils";
import getAllUnit from "@/lib/ingredient/getAllUnit";
import Loading from "./loading";

const UnitList = ({ unit, setUnit }: UnitListProps) => {
  const [openUnit, setOpenUnit] = useState(false);
  const { units, isLoading, isError } = getAllUnit();

  if (isError) return <div>Failed to load</div>;
  if (!units) {
    <Loading />;
  } else
    return (
      <DropdownMenu open={openUnit} onOpenChange={setOpenUnit}>
        <DropdownMenuTrigger asChild>
          <Button
            id="cateList"
            variant="outline"
            role="combobox"
            aria-expanded={openUnit}
            className="justify-between w-full px-2"
          >
            {unit
              ? units.find((item) => item.id === unit)?.name
              : "Chọn đơn vị"}
            <ChevronsUpDown className="ml-1 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className=" DropdownMenuContent">
          <Command>
            <CommandInput placeholder="Tìm điều kiện lọc" />
            <CommandEmpty className="p-2 text-sm">
              Không tìm thấy điều kiện lọc.
            </CommandEmpty>
            <CommandGroup>
              {units.map((item) => (
                <CommandItem
                  value={item.name}
                  key={item.id}
                  onSelect={() => {
                    setUnit(item.id);
                    setOpenUnit(false);
                  }}
                >
                  <Check
                    className={cn(
                      "mr-2 h-4 w-4",
                      item.id === unit ? "opacity-100" : "opacity-0"
                    )}
                  />
                  {item.name}
                </CommandItem>
              ))}
            </CommandGroup>
          </Command>
        </DropdownMenuContent>
      </DropdownMenu>
    );
};

export default UnitList;
