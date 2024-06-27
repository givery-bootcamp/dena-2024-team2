import { describe, expect, it } from "vitest";

import { formatDateTime } from "./date-utils";

describe("formatDateTime", () => {
	describe("正しいフォーマットが指定された場合", () => {
		describe("時間まで指定された場合", () => {
			it("Dateが yyyy-MM-dd HH:mm:ss 形式でフォーマットされる", () => {
				const date = new Date("2022-01-01T12:34:56");
				const format = "yyyy-MM-dd HH:mm:ss";
				const expected = "2022-01-01 12:34:56";
				const result = formatDateTime(date, format);
				expect(result).toEqual(expected);
			});
		});
		describe("日付までの指定だった場合", () => {
			it("Dateが yyyy/MM/dd 形式でフォーマットされる", () => {
				const date = new Date("2022-01-01T12:34:56");
				const format = "yyyy/MM/dd";
				const expected = "2022/01/01";
				const result = formatDateTime(date, format);
				expect(result).toEqual(expected);
			});
		});
	});

	describe("不正なフォーマットが指定された場合", () =>
		it("エラーがスローされる", () => {
			const date = new Date("2022-01-01T12:34:56");
			const format = "invalid";
			expect(() => formatDateTime(date, format)).toThrow();
		}));
});
