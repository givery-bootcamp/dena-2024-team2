import clsx from "clsx";
import styles from "./button.module.scss";

type Props = {
	children: React.ReactNode;
	size: "sm" | "md" | "lg";
	variant: "fill" | "outline" | "ghost";
	color: "primary" | "secondary";
	type?: "button" | "submit" | "reset";
	onClick: () => void;
};

export const Button = ({
	children,
	size,
	variant,
	color,
	type = "button",
	onClick,
}: Props) => {
	return (
		<button
			className={clsx(
				styles.root,
				styles.button,
				styles[size],
				styles[variant],
				styles[color],
			)}
			type={type}
			onClick={onClick}
		>
			{children}
		</button>
	);
};
