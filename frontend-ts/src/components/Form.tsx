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

/*
  <p className="text-sm italic text-red-600 mt-0.5">{message}</p>
*/
