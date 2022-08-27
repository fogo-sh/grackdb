import { useState } from "react";
import { Select, SelectItem } from "./Select";

export function Input({
  id,
  display,
  type = "text",
}: {
  id: string;
  display: string;
  type?: HTMLInputElement["type"];
}) {
  return (
    <div className="flex flex-col">
      <label htmlFor={id} className="mb-1">
        {display}
      </label>
      <input className="inp" id={id} name={id} autoComplete="off" type={type} />
    </div>
  );
}

export function SelectInput({
  id,
  display,
  options,
}: {
  id: string;
  display: string;
  options: SelectItem[];
}) {
  const [selected, setSelected] = useState<SelectItem>();
  return (
    <div className="flex flex-col">
      <label htmlFor={id} className="mb-1">
        {display}
      </label>
      <Select options={options} selected={selected} onChange={setSelected} />
      <input type="hidden" name={id} value={selected?.id ?? ""} />
    </div>
  );
}

/*
  TODO error state
  <p className="text-sm italic text-red-600 mt-0.5">{message}</p>
*/
