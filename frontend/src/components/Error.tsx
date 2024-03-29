const errorToText = (error: any) => {
  const text = error.statusText || error.message;
  return text;
};

export default function ErrorNotice({ error }: { error: any }) {
  // TODO any usage, probably fine
  console.error(error);

  return (
    <div className="px-6 py-2 font-semibold">
      <h1 className="text-2xl italic text-red-600">Oops!</h1>
      <p className="text-lg text-red-700">
        Sorry, an unexpected error has occurred.
      </p>
      <div className="my-3 p-3 bg-black text-white rounded-lg overflow-scroll">
        <pre>
          <p>{errorToText(error)}</p>
        </pre>
      </div>
      <span>
        Try{" "}
        <button
          onClick={() => window.location.reload()}
          className="underline text-blue-700"
        >
          refreshing the page
        </button>
        , <i>might</i> fix things!
      </span>
    </div>
  );
}
