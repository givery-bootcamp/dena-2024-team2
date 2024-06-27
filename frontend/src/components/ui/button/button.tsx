import clsx from "clsx";
import styles from "./button.module.scss";

type Props = {
	children: React.ReactNode;
	size: "sm" | "md" | "lg";
	variant: "fill" | "outline" | "ghost";
	color: "primary" | "secondary";
	type?: "button" | "submit" | "reset";
	disabled?: boolean;
	onClick?: () => void;
};

export const Button = ({
	children,
	size,
	variant,
	color,
	type = "button",
	disabled = false,
	onClick,
}: Props) => {
	const handleClick = () => {
		onClick?.();
	};

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
			onClick={handleClick}
			disabled={disabled}
		>
			{children}
		</button>
	);
};
