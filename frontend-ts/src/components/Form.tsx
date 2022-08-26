export function Input({
  id,
  name,
  display,
  type = "text",
}: {
  id: string;
  name: string;
  display: string;
  type?: HTMLInputElement["type"];
}) {
  return (
    <div className="flex flex-col">
      <label htmlFor={id} className="mb-1">
        {display}
      </label>
      <input
        className="inp"
        id={id}
        autoComplete="off"
        type={type}
        name={name}
      />
    </div>
  );
}

/*
  <p className="text-sm italic text-red-600 mt-0.5">{message}</p>
*/
