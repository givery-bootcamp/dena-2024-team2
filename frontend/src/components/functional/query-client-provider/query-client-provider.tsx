import { QueryClientProvider as OriginalProvider } from "@tanstack/react-query";
import { queryClient } from "~/utils/query-client";

type Props = {
	children: React.ReactNode;
};

export const QueryClientProvider = ({ children }: Props) => {
	return <OriginalProvider client={queryClient}>{children}</OriginalProvider>;
};
