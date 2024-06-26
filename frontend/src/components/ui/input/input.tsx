import styles from "./input.module.scss";

type Props = React.InputHTMLAttributes<HTMLInputElement> & {
	type: "text" | "password" | "email";
};

export const Input = ({ type, name, value, onChange, ...rest }: Props) => {
	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		onChange?.(e);
	};

	return (
		<input
			type={type}
			className={styles.input}
			name={name}
			id={name}
			value={value}
			onChange={handleChange}
			{...rest}
		/>
	);
};
