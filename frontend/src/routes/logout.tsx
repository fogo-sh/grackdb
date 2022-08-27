import Cookies from "js-cookie";
import toast from "react-hot-toast";
import { LoaderFunction, redirect } from "react-router-dom";

export const loader: LoaderFunction = async () => {
  Cookies.remove("jwt");
  toast.success("Logged out successfully");
  return redirect("/");
};

export function LogoutPage() {
  return null;
}
