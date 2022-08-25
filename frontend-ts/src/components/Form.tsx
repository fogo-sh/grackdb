export function Input({ id, name, type = "text" }) {
  return (
    <div className="flex flex-col">
      <label htmlFor={id} className="mb-1">
        {name}
      </label>
      <input className="inp" id={id} autoComplete="off" type={type} />
    </div>
  );
}

/*
  <p className="text-sm italic text-red-600 mt-0.5">{message}</p>
*/
