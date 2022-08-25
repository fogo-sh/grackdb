import { Outlet } from "react-router-dom";
import { Header } from "~/components/Header";

export function RootPage() {
  return (
    <div className="w-11/12 sm:w-2/3 lg:w-7/12 xl:w-1/2 2xl:w-5/12 mx-auto my-3">
      <Header />
      <hr className="my-3" />
      <Outlet />
    </div>
  );
}
