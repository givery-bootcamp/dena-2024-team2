import styles from "./text.module.scss";

type Props = {
	as: "p" | "span" | "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
	size: "sm" | "md" | "lg" | "xl";
	children: React.ReactNode;
};

export const Text = ({ as, size, children }: Props) => {
	const Component = as;
	return (
		<Component className={`${styles.root} ${styles[size]}`}>
			{children}
		</Component>
	);
};
