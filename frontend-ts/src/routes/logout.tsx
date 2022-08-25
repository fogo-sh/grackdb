import Cookies from "js-cookie";
import { useEffect } from "react";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../providers/AuthProvider";

export function LogoutPage() {
  const { refreshCurrentUser } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    Cookies.remove("jwt");
    refreshCurrentUser();
    navigate("/");
    toast.success("Logged out successfully");
  }, []);

  return null;
}
