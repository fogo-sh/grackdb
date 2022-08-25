import { QueryClient } from "@tanstack/react-query";

export const dataSource = {
  endpoint: "/query",
  fetchParams: { headers: { "Content-Type": "application/json" } },
};

export const queryClient = new QueryClient();
