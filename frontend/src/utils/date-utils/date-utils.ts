import { formatDate } from "date-fns";

export const formatDateTime = (date: Date, format: string): string => {
	return formatDate(date, format);
};
