import { Listbox } from "@headlessui/react";
import { FaCaretDown, FaCheck } from "react-icons/fa";

export type SelectItem = { id: string; value: string };

export function Select({
  options,
  selected,
  onChange,
}: {
  options: SelectItem[];
  selected?: SelectItem;
  onChange: (item: SelectItem) => void;
}) {
  return (
    <Listbox value={selected} onChange={onChange}>
      <div className="relative min-w-[15rem] h-full">
        <Listbox.Button className="relative w-full h-full border border-gray-500 rounded-sm pl-3 pr-2 py-1 flex items-center justify-between mx-auto">
          {selected?.value ?? "???"} <FaCaretDown className="ml-1 opacity-50" />
        </Listbox.Button>
        <Listbox.Options className="absolute w-full mt-1 overflow-auto text-base bg-white rounded-m max-h-60 border border-gray-500">
          {options.map((option) => (
            <Listbox.Option key={option.id} value={option}>
              {({ selected, active }) => (
                <div
                  className={`${
                    active ? "bg-gray-100" : ""
                  } flex items-center px-2.5 py-1.5`}
                >
                  <span
                    className={`${
                      selected ? "font-medium" : "font-normal"
                    } block truncate`}
                  >
                    {option.value}
                  </span>
                  {selected ? (
                    <FaCheck className="ml-2 opacity-50" aria-hidden="true" />
                  ) : null}
                </div>
              )}
            </Listbox.Option>
          ))}
        </Listbox.Options>
      </div>
    </Listbox>
  );
}
