import { useRouteError } from "react-router-dom";
import ErrorNotice from "./components/Error";

export default function ErrorPage() {
  const error = useRouteError() as any; // TODO any usage, probably fine

  return <ErrorNotice error={error} />;
}
