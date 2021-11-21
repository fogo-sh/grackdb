import Cookies from "js-cookie";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../providers/AuthProvider";

export function LogoutPage() {
	const { refreshCurrentUser } = useAuth();
	const navigate = useNavigate();

	useEffect(() => {
		Cookies.remove("jwt");
		refreshCurrentUser();
		navigate("/");
	}, []);

	return null;
}
