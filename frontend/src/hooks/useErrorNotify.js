import { useEffect } from "react";
import toast from "react-hot-toast";

export const useErrorNotify = (error) => {
	useEffect(() => {
		if (error) {
			console.error(error);
			toast.error(error.message);
		}
	}, [error]);
};
