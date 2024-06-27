const camelToSnake = (str: string): string =>
	str.replace(/[A-Z]/g, (s) => `_${s.toLowerCase()}`);
const snakeToCamel = (str: string): string =>
	str.replace(/_./g, (s) => s.slice(1).toUpperCase());

export const convertKeyCaseRecursive = (
	obj: unknown,
	to: "camel" | "snake",
): Record<string, unknown> | Record<string, unknown>[] | unknown => {
	if (obj === null || typeof obj !== "object") {
		return obj;
	}

	if (Array.isArray(obj)) {
		return obj.map((v) => convertKeyCaseRecursive(v, to));
	}

	return Object.fromEntries(
		Object.entries(obj).map(([key, value]) => {
			const newKey = to === "camel" ? snakeToCamel(key) : camelToSnake(key);
			return [newKey, convertKeyCaseRecursive(value, to)];
		}),
	);
};
