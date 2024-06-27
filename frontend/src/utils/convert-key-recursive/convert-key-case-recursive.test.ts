import { describe, expect, it } from "vitest";
import { convertKeyCaseRecursive } from "./convert-key-case-recursive";

const targetObj = {
	snake_case: {
		childCamelCase: null,
		child_snake_case: null,
	},
	camelCase: {
		childCamelCase: null,
		child_snake_case: null,
	},
};

describe("convertKeyCaseRecursive", () => {
	describe("to snake_case", () => {
		it("snake_caseに変換される", () => {
			const obj = targetObj;
			const result = convertKeyCaseRecursive(obj, "snake");
			expect(result).toEqual({
				snake_case: {
					child_camel_case: null,
					child_snake_case: null,
				},
				camel_case: {
					child_camel_case: null,
					child_snake_case: null,
				},
			});
		});
	});

	describe("to camelCase", () => {
		it("camelCaseに変換される", () => {
			const obj = targetObj;
			const result = convertKeyCaseRecursive(obj, "camel");
			expect(result).toEqual({
				snakeCase: {
					childCamelCase: null,
					childSnakeCase: null,
				},
				camelCase: {
					childCamelCase: null,
					childSnakeCase: null,
				},
			});
		});
	});
});
