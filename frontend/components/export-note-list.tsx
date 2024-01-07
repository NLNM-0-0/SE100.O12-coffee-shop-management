import { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { Button } from "./ui/button";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import { Command, CommandEmpty, CommandGroup, CommandItem } from "./ui/command";
import { cn, reasonToString } from "@/lib/utils";
import { ExportReason } from "@/types";

export interface StatusNoteListProps {
  status: string;
  setStatus: (value: string) => void;
}
const statuses = [ExportReason.Damaged, ExportReason.OutOfDate];
const ExportNoteList = ({ status, setStatus }: StatusNoteListProps) => {
  const [open, setOpen] = useState(false);
  return (
    <div className="flex gap-1">
      <DropdownMenu open={open} onOpenChange={setOpen}>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            className="justify-between w-full pl-2"
          >
            {status
              ? reasonToString(status as ExportReason)
              : "Chọn trạng thái"}
            <LuChevronsUpDown className="ml-1 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="DropdownMenuContent">
          <Command>
            <CommandEmpty className="py-2 px-6">
              <div className="text-sm">Không tìm thấy</div>
            </CommandEmpty>
            <CommandGroup className="max-h-48 overflow-y-auto">
              {statuses.map((item: ExportReason) => (
                <CommandItem
                  value={item}
                  key={item}
                  onSelect={() => {
                    setStatus(item);
                    setOpen(false);
                  }}
                >
                  <LuCheck
                    className={cn(
                      "mr-2 h-4 w-4",
                      item === status ? "opacity-100" : "opacity-0"
                    )}
                  />
                  {reasonToString(item)}
                </CommandItem>
              ))}
            </CommandGroup>
          </Command>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
};

export default ExportNoteList;
