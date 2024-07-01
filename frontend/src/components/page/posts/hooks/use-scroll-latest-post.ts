import { useEffect, useRef } from "react";
import type { Post } from "~/domains/posts";

export const useScrollLatestPost = (posts?: Post[]) => {
	const boardRef = useRef<HTMLDivElement>(null);

	useEffect(() => {
		if (posts && boardRef.current) {
			boardRef.current.scrollTop = boardRef.current.scrollHeight;
		}
	}, [posts]);

	return { boardRef };
};
