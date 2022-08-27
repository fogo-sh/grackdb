import { useRouteError } from "react-router-dom";

export default function ErrorPage() {
  const error = useRouteError() as any; // TODO any usage, probably fine
  console.error(error);

  return (
    <div className="px-6 py-2 font-semibold">
      <h1 className="text-2xl italic text-red-600">Oops!</h1>
      <p className="text-lg text-red-700">
        Sorry, an unexpected error has occurred.
      </p>
      <div className="my-3 p-3 bg-black text-white rounded-lg">
        <p>{error.statusText || error.message}</p>
      </div>
    </div>
  );
}
