import clsx from "clsx";
import styles from "./text.module.scss";

type Props = {
	as: "p" | "span" | "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
	size: "sm" | "md" | "lg" | "xl";
	bold: boolean;
	children: React.ReactNode;
};

export const Text = ({ as, size, bold, children }: Props) => {
	const Component = as;
	return (
		<Component
			className={clsx([styles.root, styles[size], { [styles.bold]: bold }])}
		>
			{children}
		</Component>
	);
};
